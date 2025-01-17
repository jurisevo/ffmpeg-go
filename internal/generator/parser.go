package main

import (
	"fmt"
	"github.com/go-clang/bootstrap/clang"
	"log"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"unicode"
)

var (
	AVLibPath, _ = filepath.Abs("include")
)

var files = []string{
	"libavcodec/avcodec.h",
	"libavcodec/codec.h",
	"libavcodec/codec_desc.h",
	"libavcodec/codec_id.h",
	"libavcodec/codec_par.h",
	"libavcodec/defs.h",
	"libavcodec/packet.h",
	"libavcodec/version.h",
	"libavcodec/version_major.h",
	"libavdevice/version.h",
	"libavdevice/version_major.h",
	"libavfilter/avfilter.h",
	"libavfilter/buffersink.h",
	"libavfilter/buffersrc.h",
	"libavfilter/version.h",
	"libavfilter/version_major.h",
	"libavformat/avformat.h",
	"libavformat/avio.h",
	"libavformat/version.h",
	"libavformat/version_major.h",
	"libavutil/avutil.h",
	"libavutil/buffer.h",
	"libavutil/channel_layout.h",
	"libavutil/dict.h",
	"libavutil/error.h",
	"libavutil/frame.h",
	"libavutil/hwcontext.h",
	"libavutil/log.h",
	"libavutil/mathematics.h",
	"libavutil/mem.h",
	"libavutil/opt.h",
	"libavutil/pixfmt.h",
	"libavutil/rational.h",
	"libavutil/samplefmt.h",
	"libavutil/version.h",
	"libpostproc/version.h",
	"libpostproc/version_major.h",
	"libswresample/version.h",
	"libswresample/version_major.h",
	"libswscale/version.h",
	"libswscale/version_major.h",
}

func Parse() *Module {
	p := &Parser{
		mod: &Module{
			functions: make(map[string]*Function),
			structs:   make(map[string]*Struct),
			enums:     make(map[string]*Enum),
			callbacks: make(map[string]*Function),
			constants: make(map[string]*Constant),
		},
	}

	for _, file := range files {
		filePath := path.Join(AVLibPath, file)

		fmt.Println(filePath)

		p.parseFile(fmt.Sprintf("[%v]", file), filePath)
	}

	return p.mod
}

type Parser struct {
	path string
	mod  *Module
	tu   clang.TranslationUnit
}

func (p *Parser) parseFile(indent string, path string) {
	p.path = path

	idx := clang.NewIndex(0, 1)
	defer idx.Dispose()

	tu := idx.ParseTranslationUnit(
		path,
		[]string{
			"-fparse-all-comments",
			fmt.Sprintf("-I%v", AVLibPath),
			"-I/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/System/Library/Frameworks/Kernel.framework/Headers/",
		},
		nil,
		clang.TranslationUnit_IncludeBriefCommentsInCodeCompletion|clang.TranslationUnit_DetailedPreprocessingRecord,
	)
	defer tu.Dispose()

	diagnostics := tu.Diagnostics()
	for _, d := range diagnostics {
		fmt.Println("PROBLEM:", d.Spelling())
	}

	p.tu = tu

	tu.
		TranslationUnitCursor().
		Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
			p.parseTopLevel(indent, cursor)

			return clang.ChildVisit_Continue
		})
}

func (p *Parser) parseTopLevel(indent string, c clang.Cursor) {
	loc, _, _, _ := c.Location().FileLocation()
	if loc.Name() != p.path {
		return
	}

	switch c.Kind() {
	case clang.Cursor_TypedefDecl:
		p.parseTypedef(indent, c)

	case clang.Cursor_VarDecl:
		log.Println("vardecl", "name", c.Spelling())
		log.Println(" ", c.Type().Spelling())

	case clang.Cursor_FunctionDecl:
		p.parseFunction(indent, c)

	case clang.Cursor_MacroDefinition:
		if c.IsMacroFunctionLike() {
			return
		}

		name := c.Spelling()

		log.Println("macro def", "name", name)

		if strings.HasSuffix(name, "_H") {
			return
		}

		if strings.HasPrefix(name, "av_malloc_") {
			return
		}

		if strings.HasPrefix(name, "AV_CHANNEL_LAYOUT_") {
			return
		}

		if strings.HasSuffix(name, "_VERSION") || strings.HasSuffix(name, "_IDENT") {
			return
		}

		if name == "AV_TIME_BASE_Q" || name == "AV_CH_LAYOUT_NATIVE" {
			return
		}

		p.mod.constants[name] = &Constant{
			Name: name,
		}
		p.mod.constantOrder = append(p.mod.constantOrder, name)

	case clang.Cursor_EnumDecl:
		p.parseEnum(indent, c, false)

	case clang.Cursor_StructDecl:
		p.parseStruct(indent, c, false)

	case clang.Cursor_MacroExpansion, clang.Cursor_InclusionDirective, clang.Cursor_UnionDecl:
		// ignore

	default:
		log.Panicln("Unexpected top level", "kind", c.Kind().String())
	}
}

func (p *Parser) parseTypedef(indent string, c clang.Cursor) {
	name := c.Spelling()
	indent = fmt.Sprintf("%v[%v]", indent, name)

	log.Println("typedef", "name", c.Spelling())

	var params []*Param

	c.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		log.Println("  --- ", "kind", cursor.Kind().String(), "name", cursor.Spelling())

		if cursor.Kind() == clang.Cursor_ParmDecl {
			name := cursor.Spelling()
			if name == "" {
				log.Fatal("no param name")
			}

			ty := p.parseType(fmt.Sprintf("%v[%v]", indent, name), cursor.Type())

			params = append(params, &Param{
				Name: name,
				Type: ty,
			})
		}

		return clang.ChildVisit_Continue
	})

	log.Println("dk ", c.Definition().Kind())
	log.Println("ds ", c.Definition().Spelling())
	log.Println("s ", c.TypedefDeclUnderlyingType().Spelling())
	log.Println("k ", c.TypedefDeclUnderlyingType().Kind())
	log.Println("cs ", c.TypedefDeclUnderlyingType().CanonicalType().Spelling())
	log.Println("ck ", c.TypedefDeclUnderlyingType().CanonicalType().Kind())

	if len(params) > 0 {
		ut := c.TypedefDeclUnderlyingType()
		pt := ut.PointeeType()

		ptr := true

		if pt.Kind() == clang.Type_Invalid {
			pt = ut
			ptr = false
		}

		if pt.NumArgTypes() != int32(len(params)) {
			log.Panicln("arg mismatch", pt.NumArgTypes(), int32(len(params)))
		}

		result := p.parseType(fmt.Sprintf("%v[%v]", indent, name), pt.ResultType())

		if _, ok := p.mod.callbacks[name]; ok {
			log.Panicln("callback already exists")

			return
		}

		p.mod.callbacks[name] = &Function{
			Name:   name,
			Args:   params,
			Result: result,
			Ptr:    ptr,
		}
		p.mod.callbackOrder = append(p.mod.callbackOrder, name)

		return
	}

	dec := c.TypedefDeclUnderlyingType().Declaration()

	dec.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		log.Println("  Inner ", "kind", cursor.Kind().String(), "name", cursor.Spelling())

		return clang.ChildVisit_Continue
	})

	switch dec.Kind() {
	case clang.Cursor_StructDecl:
		p.parseStruct(indent, dec, true)

	case clang.Cursor_EnumDecl:
		p.parseEnum(indent, dec, true)

	default:
		log.Panicln("Unknown typedef", "kind", dec.Kind())
	}

}

func processComment(in string) string {
	txt := in
	txt = strings.ReplaceAll(txt, "\r\n", "\n")
	txt = strings.TrimSpace(txt)
	txt = strings.TrimPrefix(txt, "/**\n")
	txt = strings.TrimSuffix(txt, "*/")

	txt = strings.TrimRightFunc(txt, unicode.IsSpace)

	var rebuilt []string

	for _, s := range strings.Split(txt, "\n") {
		s = strings.TrimSpace(s)
		s = strings.TrimPrefix(s, "* ")
		s = strings.TrimPrefix(s, "/// ")
		s = strings.TrimPrefix(s, "///")
		s = strings.TrimPrefix(s, "// ")
		s = strings.TrimPrefix(s, "//")

		if strings.HasPrefix(s, "/**") {
			rebuilt = nil
			s = strings.TrimPrefix(s, "/** ")
			s = strings.TrimPrefix(s, "/**")

			if strings.TrimSpace(s) == "" {
				continue
			}
		}

		if strings.HasPrefix(s, "/*") {
			rebuilt = nil
			s = strings.TrimPrefix(s, "/* ")
			s = strings.TrimPrefix(s, "/*")

			if strings.TrimSpace(s) == "" {
				continue
			}
		}

		if strings.HasPrefix(s, "@defgroup") || strings.HasPrefix(s, "@ingroup") ||
			strings.HasPrefix(s, "@addtogroup") || strings.HasPrefix(s, "@}") {
			continue
		}

		if strings.HasPrefix(s, "@{") {
			rebuilt = nil
			continue
		}

		if s == "*" {
			s = ""
		}

		// Double space to enter verbatim mode
		rebuilt = append(rebuilt, fmt.Sprintf("  %v", s))
	}

	txt = strings.Join(rebuilt, "\n")
	txt = strings.TrimRightFunc(txt, unicode.IsSpace)

	if strings.Count(txt, "\n") == 0 {
		txt = strings.TrimLeftFunc(txt, unicode.IsSpace)
		txt = strings.TrimPrefix(txt, "<")
		txt = fmt.Sprintf("  %v", txt)
	}

	if strings.TrimSpace(txt) == "" {
		return ""
	}

	return txt
}

func (p *Parser) parseFunction(indent string, c clang.Cursor) {
	name := c.Spelling()

	if _, ok := p.mod.functions[name]; ok {
		log.Panicln("Function", name, "already exists")
	}

	log.Println("Parsing function", name)
	indent = fmt.Sprintf("%v[%v]", indent, name)

	comment := processComment(c.RawCommentText())

	result := p.parseType(fmt.Sprintf("%v[return]", indent), c.ResultType())

	var args []*Param

	for i := 0; i < int(c.NumArguments()); i++ {
		arg := c.Argument(uint32(i))

		if arg.Kind() != clang.Cursor_ParmDecl {
			log.Panicln(indent, "Argument not of parmdecl type", arg.Kind())
		}

		name := arg.Spelling()
		if name == "" {
			log.Println(indent, "no param name")

			name = fmt.Sprintf("param%v", i)
		}

		aIndent := fmt.Sprintf("%v[%v]", indent, name)

		ty := p.parseType(aIndent, arg.Type())

		args = append(args, &Param{
			Name: name,
			Type: ty,
		})
	}

	p.mod.functions[name] = &Function{
		Name:     name,
		Args:     args,
		Result:   result,
		Variadic: c.IsVariadic(),
		Comment:  comment,
	}
	p.mod.functionOrder = append(p.mod.functionOrder, name)
}

func (p *Parser) parseEnum(indent string, c clang.Cursor, typedef bool) {
	log.Println("enum", "name", c.Spelling())

	name := c.Spelling()
	indent = fmt.Sprintf("%v[%v]", indent, name)

	if strings.HasPrefix(name, "enum (unnamed") {
		log.Println(indent, "Treating unnamed enum as constants")

		c.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
			if cursor.Kind() != clang.Cursor_EnumConstantDecl {
				log.Panicln("Unknown enum type", "kind", cursor.Kind().String())
			}

			name := cursor.Spelling()

			p.mod.constants[name] = &Constant{
				Name:     name,
				FromEnum: true,
			}
			p.mod.constantOrder = append(p.mod.constantOrder, name)

			return clang.ChildVisit_Continue
		})

		return
	}

	if strings.Contains(name, " ") {
		log.Panicln(indent, "Name contains spaces")
	}

	if val, ok := p.mod.enums[name]; ok && len(val.Constants) > 0 {
		log.Println(indent, "already exists")

		if typedef {
			val.Typedefd = true
		}

		return
	}

	comment := processComment(c.RawCommentText())

	enum := &Enum{
		Name:     name,
		Typedefd: typedef,
		Comment:  comment,
	}

	c.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() != clang.Cursor_EnumConstantDecl {
			log.Panicln("Unknown enum type", "kind", cursor.Kind().String())
		}

		comment := processComment(cursor.RawCommentText())

		enum.Constants = append(enum.Constants, &Constant{
			Name:    cursor.Spelling(),
			Comment: comment,
		})

		return clang.ChildVisit_Continue
	})

	p.mod.enumOrder = slices.DeleteFunc(p.mod.enumOrder, func(s string) bool {
		return s == name
	})

	p.mod.enums[name] = enum
	p.mod.enumOrder = append(p.mod.enumOrder, name)
}

func (p *Parser) parseStruct(indent string, c clang.Cursor, typedef bool) {
	log.Println("struct", "name", c.Spelling())

	name := c.Spelling()
	indent = fmt.Sprintf("%v[%v]", indent, name)

	if val, ok := p.mod.structs[name]; ok && len(val.Fields) > 0 {
		log.Println(indent, "already exists")

		if typedef {
			val.Typedefd = true
		}

		return
	}

	comment := processComment(c.RawCommentText())

	s := &Struct{
		Name:     name,
		Typedefd: typedef,
		Comment:  comment,
	}

	c.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() == clang.Cursor_FieldDecl {
			name := cursor.Spelling()
			if name == "" {
				log.Fatal("no field name")
			}

			cmt := processComment(cursor.RawCommentText())

			fIndent := fmt.Sprintf("%v[%v]", indent, name)

			ty := p.parseType(fIndent, cursor.Type())

			s.Fields = append(s.Fields, &Field{
				Name:     name,
				Type:     ty,
				BitWidth: cursor.FieldDeclBitWidth(),
				Comment:  cmt,
			})
		}

		return clang.ChildVisit_Continue
	})

	p.mod.structOrder = slices.DeleteFunc(p.mod.structOrder, func(s string) bool {
		return s == name
	})

	p.mod.structs[name] = s
	p.mod.structOrder = append(p.mod.structOrder, name)
}
