package main

import "C"
import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

var primTypes = map[string]string{
	"int":      "int",
	"uint":     "uint",
	"char":     "uint8",
	"uchar":    "uint8",
	"ulong":    "uint32",
	"uint8_t":  "uint8",
	"uint16_t": "uint16",
	"uint32_t": "uint32",
	"int64_t":  "int64",
	"uint64_t": "uint64",
	"size_t":   "uint64",
	"float":    "float32",
	"double":   "float64",
}

type Generator struct {
	input *Module
}

func Gen(i *Module) {
	g := &Generator{
		input: i,
	}

	g.generateConstants()
	g.generateEnums()
	g.generateStructs()
	g.generateFuncs()
}

func newFile() *jen.File {
	o := jen.NewFile("ffmpeg")

	for _, file := range files {
		o.CgoPreamble(fmt.Sprintf("#include <%v>", file))
	}

	return o
}

func (g *Generator) generateConstants() {
	i := g.input

	o := newFile()

	for _, constName := range i.constantOrder {
		constant := i.constants[constName]

		log.Println("Generating constant", constant.Name)

		goName := g.convCamel(constant.Name)

		if strings.HasPrefix(constName, "AVERROR_") {
			goName = fmt.Sprintf("%vConst", goName)
		}

		o.Commentf("%v wraps %v.", goName, constant.Name)

		o.Const().Id(goName).Op("=").Qual("C", constName)
	}

	err := o.Save("constants.gen.go")
	if err != nil {
		log.Panicln(err)
	}
}

func (g *Generator) generateEnums() {
	i := g.input

	o := newFile()

	for _, enumName := range i.enumOrder {
		enum := i.enums[enumName]

		log.Println("Generating enum", enum.Name)
		o.Commentf("--- Enum %v ---", enum.Name)
		o.Line()

		goName := enumName

		o.Commentf("%v wraps %v.", goName, enum.Name)

		if enum.Comment != "" {
			o.Comment(enum.Comment)
		}

		cName := enum.CName()

		o.Type().Id(goName).Qual("C", cName)

		o.Const().Id(fmt.Sprintf("SizeOf%v", goName)).Op("=").Qual("C", fmt.Sprintf("sizeof_%v", cName))

		o.Func().
			Id(fmt.Sprintf("To%vArray", goName)).
			Params(jen.Id("ptr").Qual("unsafe", "Pointer")).
			Op("*").Id("Array").Types(jen.Id(goName)).
			Block(
				jen.If(jen.Id("ptr").Op("==").Id("nil")).Block(
					jen.Return(jen.Id("nil")),
				),
				jen.Line(),
				jen.Return(
					jen.Op("&").Id("Array").Types(jen.Id(goName)).Values(jen.Dict{
						jen.Id("ptr"):      jen.Id("ptr"),
						jen.Id("elemSize"): jen.Id(fmt.Sprintf("SizeOf%v", goName)),

						jen.Id("loadPtr"): jen.Func().
							Params(jen.Id("pointer").Qual("unsafe", "Pointer")).
							Id(goName).
							Block(
								jen.Id("ptr").Op(":=").Parens(jen.Op("*").Id(goName)).Parens(jen.Id("pointer")),
								jen.Return(jen.Op("*").Id("ptr")),
							),

						jen.Id("storePtr"): jen.Func().
							Params(
								jen.Id("pointer").Qual("unsafe", "Pointer"),
								jen.Id("value").Id(goName),
							).
							Block(
								jen.Id("ptr").Op(":=").Parens(jen.Op("*").Id(goName)).Parens(jen.Id("pointer")),
								jen.Op("*").Id("ptr").Op("=").Id("value"),
							),
					}),
				),
			)

		o.Line()

		o.Func().
			Id(fmt.Sprintf("Alloc%vArray", goName)).
			Params(jen.Id("size").Id("uint64")).
			Op("*").Id("Array").Types(jen.Id(goName)).
			Block(
				jen.Return(jen.Id(fmt.Sprintf("To%vArray", goName)).Params(
					jen.Id("AVCalloc").Params(jen.Id("size"), jen.Id(fmt.Sprintf("SizeOf%v", goName))),
				)),
			)

		var valDefs []jen.Code

		for _, constant := range enum.Constants {
			constName := g.convCamel(constant.Name)

			valDefs = append(valDefs, jen.Commentf("%v wraps %v.", constName, constant.Name))

			if constant.Comment != "" {
				valDefs = append(valDefs, jen.Comment(constant.Comment))
			}

			valDefs = append(valDefs, jen.Id(constName).Id(goName).Op("=").Qual("C", constant.Name))
		}

		if len(valDefs) > 0 {
			o.Const().Defs(valDefs...)
		}
	}

	err := o.Save("enums.gen.go")
	if err != nil {
		log.Panicln(err)
	}
}

func (g *Generator) generateStructs() {
	i := g.input

	o := newFile()

	for _, stName := range i.structOrder {
		st := i.structs[stName]

		log.Println("Generating struct", st.Name)
		o.Commentf("--- Struct %v ---", st.Name)
		o.Line()

		goName := st.Name

		o.Commentf("%v wraps %v.", goName, st.Name)

		if st.Comment != "" {
			o.Comment(st.Comment)
		}

		cName := st.CName()

		if st.ByValue {
			o.Type().Id(goName).Struct(
				jen.Id("value").Qual("C", cName),
			)
		} else {
			o.Type().Id(goName).Struct(
				jen.Id("ptr").Op("*").Qual("C", cName),
			)

			o.Func().
				Params(jen.Id("s").Op("*").Id(goName)).
				Id("RawPtr").
				Params().
				Qual("unsafe", "Pointer").
				Block(jen.Return(jen.Qual("unsafe", "Pointer").Params(jen.Id("s").Dot("ptr"))))

			o.Line()

			o.Func().
				Id(fmt.Sprintf("To%vArray", goName)).
				Params(jen.Id("ptr").Qual("unsafe", "Pointer")).
				Op("*").Id("Array").Types(jen.Op("*").Id(goName)).
				Block(
					jen.If(jen.Id("ptr").Op("==").Id("nil")).Block(
						jen.Return(jen.Id("nil")),
					),
					jen.Line(),
					jen.Return(
						jen.Op("&").Id("Array").Types(jen.Op("*").Id(goName)).Values(jen.Dict{
							jen.Id("ptr"):      jen.Id("ptr"),
							jen.Id("elemSize"): jen.Id("ptrSize"),

							jen.Id("loadPtr"): jen.Func().
								Params(jen.Id("pointer").Qual("unsafe", "Pointer")).
								Op("*").Id(goName).
								Block(
									jen.Id("ptr").Op(":=").Parens(jen.Op("**").Qual("C", cName)).Parens(jen.Id("pointer")),
									jen.Id("value").Op(":=").Op("*").Id("ptr"),

									jen.Var().Id("valueMapped").Op("*").Id(goName),
									jen.If(jen.Id("value").Op("!=").Id("nil")).Block(
										jen.Id("valueMapped").Op("=").Op("&").Id(goName).Values(jen.Dict{
											jen.Id("ptr"): jen.Id("value"),
										}),
									),
									jen.Return(jen.Id("valueMapped")),
								),

							jen.Id("storePtr"): jen.Func().
								Params(
									jen.Id("pointer").Qual("unsafe", "Pointer"),
									jen.Id("value").Op("*").Id(goName),
								).
								Block(
									jen.Id("ptr").Op(":=").Parens(jen.Op("**").Qual("C", cName)).Parens(jen.Id("pointer")),

									jen.If(jen.Id("value").Op("!=").Id("nil")).Block(
										jen.Op("*").Id("ptr").Op("=").Id("value").Dot("ptr"),
									).Else().Block(
										jen.Op("*").Id("ptr").Op("=").Id("nil"),
									),
								),
						}),
					),
				)

			o.Line()

		}

	fieldLoop:
		for _, field := range st.Fields {
			fName := strcase.ToCamel(field.Name)

			cName := field.Name
			if cName == "type" || cName == "range" {
				cName = fmt.Sprintf("_%v", cName)
			}

			var (
				getBody    []jen.Code
				getRetType []jen.Code
				getParams  []jen.Code
				setBody    []jen.Code
				setParams  []jen.Code
				tgt        *jen.Statement

				refField bool
			)

			if field.BitWidth != -1 {
				o.Commentf("%v skipped due to bitfield", field.Name)
				o.Line()

				continue fieldLoop
			}

			if st.ByValue {
				tgt = jen.Id("s").Dot("value").Dot(cName)
			} else {
				tgt = jen.Id("s").Dot("ptr").Dot(cName)
			}

			switch v := field.Type.(type) {

			case *IdentType:

				if m, ok := primTypes[v.Name]; ok {
					getRetType = []jen.Code{jen.Id(m)}
					setParams = append(setParams, jen.Id("value").Id(m))

					getBody = append(getBody, jen.Return(jen.Id(m).Params(jen.Id("value"))))

					if v.IsAnonEnum {
						setBody = append(
							setBody,
							tgt.Op("=").Id("value"),
						)
					} else {
						setBody = append(
							setBody,
							tgt.Op("=").Params(jen.Qual("C", v.Name)).Params(jen.Id("value")),
						)
					}
				} else if s, ok := g.input.structs[v.Name]; ok {
					if s.ByValue {
						getRetType = []jen.Code{
							jen.Op("*").Id(s.Name),
						}
						setParams = append(setParams, jen.Id("value").Op("*").Id(s.Name))

						getBody = append(
							getBody,
							jen.Return(jen.Op("&").Id(s.Name).Values(jen.Dict{
								jen.Id("value"): jen.Id("value"),
							})),
						)

						setBody = append(
							setBody,
							tgt.Op("=").Id("value").Dot("value"),
						)
					} else {
						refField = true

						getRetType = []jen.Code{
							jen.Op("*").Id(s.Name),
						}

						getBody = append(
							getBody,
							jen.Return(jen.Op("&").Id(s.Name).Values(jen.Dict{
								jen.Id("ptr"): jen.Id("value"),
							})),
						)
					}

				} else if _, ok := g.input.callbacks[v.Name]; ok {
					o.Commentf("%v skipped due to ident callback", field.Name)
					o.Line()

					continue fieldLoop
				} else if e, ok := g.input.enums[v.Name]; ok {
					getRetType = []jen.Code{jen.Id(v.Name)}
					setParams = append(setParams, jen.Id("value").Id(v.Name))

					getBody = append(getBody, jen.Return(jen.Id(v.Name).Params(jen.Id("value"))))

					setBody = append(
						setBody,
						tgt.Op("=").Params(jen.Qual("C", e.CName())).Params(jen.Id("value")),
					)
				} else {
					log.Panicln("unexpected")
				}

			case *PointerType:
				switch iv := v.Inner.(type) {
				case nil:
					getRetType = []jen.Code{
						jen.Qual("unsafe", "Pointer"),
					}
					setParams = append(setParams, jen.Id("value").Qual("unsafe", "Pointer"))
					getBody = append(getBody, jen.Return(jen.Id("value")))
					setBody = append(setBody, tgt.Op("=").Id("value"))

				case *IdentType:

					if iv.Name == "URLContext" || iv.Name == "AVFilterCommand" || iv.Name == "AVCodecInternal" {
						o.Commentf("%v skipped due to ptr to ignored type", field.Name)
						o.Line()

						continue fieldLoop
					} else if iv.Name == "char" {
						getRetType = []jen.Code{
							jen.Op("*").Id("CStr"),
						}
						setParams = append(setParams, jen.Id("value").Op("*").Id("CStr"))
						getBody = append(getBody, jen.Return(jen.Id("wrapCStr").Params(jen.Id("value"))))
						setBody = append(setBody, tgt.Op("=").Id("value").Dot("ptr"))

					} else if iv.Name == "uint8_t" {
						getRetType = []jen.Code{
							jen.Qual("unsafe", "Pointer"),
						}
						getBody = append(getBody, jen.Return(jen.Qual("unsafe", "Pointer").Params(jen.Id("value"))))

						setParams = append(setParams, jen.Id("value").Qual("unsafe", "Pointer"))
						setBody = append(setBody, tgt.Op("=").Params(jen.Op("*").Qual("C", iv.Name)).Params(jen.Id("value")))
					} else if _, ok := primTypes[iv.Name]; ok {
						o.Commentf("%v skipped due to prim ptr", field.Name)
						o.Line()

						continue fieldLoop
					} else if enum, ok := g.input.enums[iv.Name]; ok {
						getRetType = []jen.Code{
							jen.Op("*").Id("Array").Types(jen.Id(iv.Name)),
						}

						getBody = append(
							getBody,
							jen.Return(jen.Id(fmt.Sprintf("To%vArray", iv.Name)).Params(
								jen.Qual("unsafe", "Pointer").Params(jen.Id("value")),
							)),
						)

						setParams = append(setParams, jen.Id("value").Op("*").Id("Array").Types(jen.Id(iv.Name)))

						cName := enum.CName()

						setBody = append(
							setBody,
							jen.If(jen.Id("value").Op("!=").Id("nil")).Block(
								tgt.Clone().Op("=").Params(jen.Op("*").Qual("C", cName)).Params(jen.Id("value").Dot("ptr")),
							).Else().Block(
								tgt.Clone().Op("=").Id("nil"),
							),
						)

					} else if _, ok := g.input.callbacks[iv.Name]; ok {
						o.Commentf("%v skipped due to callback ptr", field.Name)
						o.Line()

						continue fieldLoop
					} else if ist, ok := g.input.structs[iv.Name]; ok {
						if ist.ByValue {
							o.Commentf("%v skipped due to struct value ptr", field.Name)
							o.Line()

							continue fieldLoop
						}

						getRetType = []jen.Code{
							jen.Op("*").Id(iv.Name),
						}
						setParams = append(setParams, jen.Id("value").Op("*").Id(iv.Name))

						getBody = append(
							getBody,
							jen.Var().Id("valueMapped").Op("*").Id(iv.Name),
							jen.If(jen.Id("value").Op("!=").Id("nil")).Block(
								jen.Id("valueMapped").Op("=").Op("&").Id(iv.Name).Values(jen.Dict{
									jen.Id("ptr"): jen.Id("value"),
								}),
							),
							jen.Return(jen.Id("valueMapped")),
						)

						setBody = append(
							setBody,
							jen.If(jen.Id("value").Op("!=").Id("nil")).Block(
								tgt.Clone().Op("=").Id("value").Dot("ptr"),
							).Else().Block(
								tgt.Clone().Op("=").Id("nil"),
							),
						)
					} else {
						log.Panicln("unexpected ", iv.Name)
					}

				case *FuncType:
					o.Commentf("%v skipped due to func ptr", field.Name)
					o.Line()

					continue fieldLoop

				case *PointerType:

					switch iiv := iv.Inner.(type) {
					case *IdentType:
						if st, ok := g.input.structs[iiv.Name]; ok {

							getRetType = []jen.Code{
								jen.Op("*").Id("Array").Types(jen.Op("*").Id(iiv.Name)),
							}

							getBody = append(
								getBody,
								jen.Return(jen.Id(fmt.Sprintf("To%vArray", iiv.Name)).Params(
									jen.Qual("unsafe", "Pointer").Params(jen.Id("value")),
								)),
							)

							setParams = append(setParams, jen.Id("value").Op("*").Id("Array").Types(jen.Id(iiv.Name)))

							cName := st.CName()

							setBody = append(
								setBody,
								jen.If(jen.Id("value").Op("!=").Id("nil")).Block(
									tgt.Clone().Op("=").Params(jen.Op("**").Qual("C", cName)).Params(jen.Id("value").Dot("ptr")),
								).Else().Block(
									tgt.Clone().Op("=").Id("nil"),
								),
							)
						} else {
							o.Commentf("%v skipped due to unknown ptr ptr", field.Name)
							o.Line()

							continue fieldLoop
						}

					default:
						log.Panicln("unexpected type", reflect.TypeOf(iv.Inner))
					}

				default:
					log.Panicln("unexpected type", reflect.TypeOf(v.Inner))
				}

			case *ConstArray:
				refField = true

				switch iv := v.Inner.(type) {
				case *IdentType:
					if pt, ok := primTypes[iv.Name]; ok {
						getRetType = []jen.Code{
							jen.Op("*").Id("Array").Types(jen.Id(pt)),
						}

						goName := strcase.ToCamel(pt)

						getBody = append(
							getBody,
							jen.Return(jen.Id(fmt.Sprintf("To%vArray", goName)).Params(
								jen.Qual("unsafe", "Pointer").Params(jen.Id("value")),
							)),
						)
					} else {
						o.Commentf("%v skipped due to unknown const array", field.Name)
						o.Line()

						continue fieldLoop
					}

				case *PointerType:

					switch iiv := iv.Inner.(type) {
					case *IdentType:
						if iiv.Name == "uint8_t" {
							getRetType = []jen.Code{
								jen.Op("*").Id("Array").Types(jen.Qual("unsafe", "Pointer")),
							}

							getBody = append(
								getBody,
								jen.Return(jen.Id("ToUint8PtrArray").Params(
									jen.Qual("unsafe", "Pointer").Params(jen.Id("value")),
								)),
							)
						} else if st, ok := g.input.structs[iiv.Name]; ok {
							if st.ByValue {
								o.Commentf("%v skipped due to value ident ptr const array", field.Name)
								o.Line()

								continue fieldLoop
							}

							getRetType = []jen.Code{
								jen.Op("*").Id("Array").Types(jen.Op("*").Id(iiv.Name)),
							}

							getBody = append(
								getBody,
								jen.Return(jen.Id(fmt.Sprintf("To%vArray", iiv.Name)).Params(
									jen.Qual("unsafe", "Pointer").Params(jen.Id("value")),
								)),
							)
						} else {
							o.Commentf("%v skipped due to ident pointer const array", field.Name)
							o.Line()

							continue fieldLoop
						}

					default:
						log.Panicln("unexpected type", reflect.TypeOf(iv.Inner))
					}

				case *ConstArray:
					o.Commentf("%v skipped due to multi dim const array", field.Name)
					o.Line()

					continue fieldLoop

				default:
					log.Panicln("unexpected type", reflect.TypeOf(v.Inner))
				}

			case *UnionType:
				o.Commentf("%v skipped due to union type", field.Name)
				o.Line()

				continue fieldLoop

			default:
				log.Panicln("unexpected type", reflect.TypeOf(field.Type))
			}

			if refField {
				if st.ByValue {
					o.Commentf("%v skipped due to ref field of value struct", field.Name)
					o.Line()

					continue fieldLoop
				} else {
					getBody = slices.Insert(
						getBody, 0,
						jen.Code(jen.Id("value").Op(":=").Op("&").Id("s").Dot("ptr").Dot(cName)),
					)
				}

			} else if st.ByValue {
				getBody = slices.Insert(
					getBody, 0,
					jen.Code(jen.Id("value").Op(":=").Id("s").Dot("value").Dot(cName)),
				)
			} else {
				getBody = slices.Insert(
					getBody, 0,
					jen.Code(jen.Id("value").Op(":=").Id("s").Dot("ptr").Dot(cName)),
				)
			}

			o.Commentf("%v gets the %v field.", fName, field.Name)

			if field.Comment != "" {
				o.Comment(field.Comment)
			}

			o.Func().
				Params(jen.Id("s").Op("*").Id(goName)).
				Id(fName).
				Params(getParams...).
				Add(getRetType...).
				Block(getBody...)

			o.Line()

			if len(setBody) > 0 {
				o.Commentf("Set%v sets the %v field.", fName, field.Name)

				if field.Comment != "" {
					o.Comment(field.Comment)
				}

				o.Func().
					Params(jen.Id("s").Op("*").Id(goName)).
					Id(fmt.Sprintf("Set%v", fName)).
					Params(setParams...).
					Block(setBody...)

				o.Line()
			}

		}
	}

	err := o.Save("structs.gen.go")
	if err != nil {
		log.Panicln(err)
	}
}

var (
	fileType = &PointerType{
		Inner: &IdentType{Name: "FILE"},
	}
	vaListType = &IdentType{Name: "va_list"}
)

func (g *Generator) generateFuncs() {
	i := g.input

	o := newFile()

outer:
	for _, fnName := range i.functionOrder {
		fn := i.functions[fnName]

		log.Println("Generating func", fn.Name)
		o.Commentf("--- Function %v ---", fn.Name)
		o.Line()

		// Check if function contains unsupported features
		if fn.Variadic {
			o.Commentf("%v skipped due to variadic arg.", fn.Name)
			o.Line()

			continue
		}

		if typeEquals(fn.Result, fileType) {
			o.Commentf("%v skipped due to return.", fn.Name)
			o.Line()

			continue outer
		}

		for _, arg := range fn.Args {
			skip := false

			if typeEquals(arg.Type, fileType) || typeEquals(arg.Type, vaListType) {
				skip = true
			}

			switch v := arg.Type.(type) {
			case *PointerType:
				switch v.Inner.(type) {
				case *FuncType:
					skip = true
				}
			}

			if skip {
				o.Commentf("%v skipped due to %v.", fn.Name, arg.Name)
				o.Line()

				continue outer
			}
		}

		goName := g.convCamel(fn.Name)

		var (
			params   []jen.Code
			args     []jen.Code
			body     []jen.Code
			postCall []jen.Code
		)

		for _, arg := range fn.Args {
			pName := convParamName(arg.Name)

			switch v := arg.Type.(type) {
			case *IdentType:
				if m, ok := primTypes[v.Name]; ok {
					params = append(params, jen.Id(pName).Id(m))
					args = append(args, jen.Qual("C", v.Name).Params(jen.Id(pName)))
				} else if e, ok := g.input.enums[v.Name]; ok {
					params = append(params, jen.Id(pName).Id(v.Name))
					args = append(args, jen.Qual("C", e.CName()).Params(jen.Id(pName)))
				} else if s, ok := g.input.structs[v.Name]; ok {
					if s.ByValue {
						params = append(params, jen.Id(pName).Op("*").Id(s.Name))
						args = append(args, jen.Id(pName).Dot("value"))
					} else {
						o.Commentf("%v skipped due to %v", fn.Name, pName)
						o.Line()

						continue outer
					}
				} else {
					params = append(params, jen.Id(pName).Id(v.Name))
					args = append(args, jen.Qual("C", v.Name).Params(jen.Id(pName)))
				}

			case *PointerType:
				switch iv := v.Inner.(type) {
				case nil:
					params = append(params, jen.Id(pName).Qual("unsafe", "Pointer"))
					args = append(args, jen.Id(pName))

				case *IdentType:
					if iv.Name == "char" {
						params = append(params, jen.Id(pName).Op("*").Id("CStr"))
						convName := fmt.Sprintf("tmp%v", pName)

						body = append(
							body,
							jen.Var().Id(convName).Op("*").Qual("C", "char"),
							jen.If(jen.Id(pName).Op("!=").Id("nil")).Block(
								jen.Id(convName).Op("=").Id(pName).Dot("ptr"),
							),
						)

						args = append(args, jen.Id(convName))
					} else if iv.Name == "uint8_t" || iv.Name == "uchar" {
						params = append(params, jen.Id(pName).Qual("unsafe", "Pointer"))
						args = append(args, jen.Params(jen.Op("*").Qual("C", iv.Name)).Params(jen.Id(pName)))
					} else {

						if m, ok := primTypes[iv.Name]; ok {
							params = append(params, jen.Id(pName).Op("*").Id(m))

							o.Commentf("%v skipped due to %v", fn.Name, pName)
							o.Line()
							continue outer

						} else if s, ok := g.input.structs[iv.Name]; ok {
							if s.ByValue {
								o.Commentf("%v skipped due to %v", fn.Name, pName)
								o.Line()
								continue outer
							}

							params = append(params, jen.Id(pName).Op("*").Id(iv.Name))

							convName := fmt.Sprintf("tmp%v", pName)

							body = append(
								body,
								jen.Var().Id(convName).Op("*").Qual("C", s.CName()),
								jen.If(jen.Id(pName).Op("!=").Id("nil")).Block(
									jen.Id(convName).Op("=").Id(pName).Dot("ptr"),
								),
							)

							args = append(args, jen.Id(convName))
						} else {
							//goType = jen.Id(iv.Name)

							params = append(params, jen.Id(pName).Op("*").Id(iv.Name))

							o.Commentf("%v skipped due to %v", fn.Name, pName)
							o.Line()
							continue outer
						}

						//retType = []jen.Code{
						//	jen.Op("*").Id(iv.Name),
						//}
					}

				case *PointerType:

					switch iiv := iv.Inner.(type) {
					case *IdentType:

						if iiv.Name == "uint8_t" {
							params = append(params, jen.Id(pName).Id("TODO"))

							o.Commentf("%v skipped due to %v", fn.Name, pName)
							o.Line()
							continue outer
						} else if iiv.Name == "char" {
							params = append(params, jen.Id(pName).Id("TODO"))

							o.Commentf("%v skipped due to %v", fn.Name, pName)
							o.Line()
							continue outer
						} else {

							if m, ok := primTypes[iiv.Name]; ok {
								params = append(params, jen.Id(pName).Op("**").Id(m))

								o.Commentf("%v skipped due to %v", fn.Name, pName)
								o.Line()
								continue outer
							} else if s, ok := g.input.structs[iiv.Name]; ok {
								params = append(params, jen.Id(pName).Op("**").Id(iiv.Name))

								ptrName := fmt.Sprintf("ptr%v", pName)
								tmpName := fmt.Sprintf("tmp%v", pName)
								oldName := fmt.Sprintf("oldTmp%v", pName)
								innerName := fmt.Sprintf("inner%v", pName)

								body = append(
									body,
									jen.Var().Id(ptrName).Op("**").Qual("C", s.CName()),
									jen.Var().Id(tmpName).Op("*").Qual("C", s.CName()),
									jen.Var().Id(oldName).Op("*").Qual("C", s.CName()),
									jen.If(jen.Id(pName).Op("!=").Id("nil")).Block(
										jen.Id(innerName).Op(":=").Op("*").Id(pName),
										jen.If(jen.Id(innerName).Op("!=").Id("nil")).Block(
											jen.Id(tmpName).Op("=").Id(innerName).Dot("ptr"),
											jen.Id(oldName).Op("=").Id(tmpName),
										),
										jen.Id(ptrName).Op("=").Op("&").Id(tmpName),
									),
								)

								postCall = append(
									postCall,
									jen.If(jen.Id(tmpName).Op("!=").Id(oldName).Op("&&").Id(pName).Op("!=").Id("nil")).Block(

										jen.If(jen.Id(tmpName).Op("!=").Id("nil")).Block(
											jen.Op("*").Id(pName).Op("=").Op("&").Id(iiv.Name).Values(jen.Dict{
												jen.Id("ptr"): jen.Id(tmpName),
											}),
										).Else().Block(
											jen.Op("*").Id(pName).Op("=").Id("nil"),
										),
									),
								)

								args = append(args, jen.Id(ptrName))
							} else {
								params = append(params, jen.Id(pName).Op("**").Id(iiv.Name))

								o.Commentf("%v skipped due to %v", fn.Name, pName)
								o.Line()
								continue outer
							}
						}

						//params = append(params, jen.Id(pName).Op("**").Id(iiv.Name))

					default:
						params = append(params, jen.Id(pName).Id("TODO"))

						o.Commentf("%v skipped due to %v", fn.Name, pName)
						o.Line()
						continue outer

					}

				default:
					log.Panicln("unexpected type", reflect.TypeOf(v.Inner))
				}

			case *Array:
				params = append(params, jen.Id(pName).Id("TODO"))

				o.Commentf("%v skipped due to %v", fn.Name, pName)
				o.Line()
				continue outer

			default:
				log.Panicln("unexpected type", reflect.TypeOf(arg.Type))
			}
		}

		cc := jen.Qual("C", fn.Name).Params(args...)

		var retType []jen.Code

		switch v := fn.Result.(type) {
		case nil:
			// nothing
			body = append(body, cc)
			body = append(body, postCall...)

		case *IdentType:

			body = append(body, jen.Id("ret").Op(":=").Add(cc))
			body = append(body, postCall...)

			if v.Name == "int" {
				retType = []jen.Code{jen.Params(jen.Id("int"), jen.Id("error"))}
				body = append(
					body,
					jen.Return(
						jen.Id("int").Params(jen.Id("ret")).Op(",").
							Id("WrapErr").Params(jen.Id("int").Params(jen.Id("ret"))),
					),
				)
			} else if m, ok := primTypes[v.Name]; ok {
				retType = []jen.Code{jen.Id(m)}
				body = append(body, jen.Return(jen.Id(m).Params(jen.Id("ret"))))
			} else if s, ok := g.input.structs[v.Name]; ok {
				if s.ByValue {
					retType = []jen.Code{jen.Op("*").Id(v.Name)}
					body = append(
						body,
						jen.Return(jen.Op("&").Id(v.Name).Values(jen.Dict{
							jen.Id("value"): jen.Id("ret"),
						})),
					)
				} else {
					o.Commentf("%v skipped due to return", fn.Name)
					o.Line()
					continue outer
				}
			} else {
				retType = []jen.Code{jen.Id(v.Name)}
				body = append(body, jen.Return(jen.Id(v.Name).Params(jen.Id("ret"))))
			}

		case *PointerType:
			body = append(body,
				jen.Id("ret").Op(":=").Add(cc),
			)
			body = append(body, postCall...)

			switch iv := v.Inner.(type) {
			case nil:
				retType = []jen.Code{
					jen.Qual("unsafe", "Pointer"),
				}
				body = append(body, jen.Return(jen.Id("ret")))

			case *IdentType:

				if iv.Name == "char" {
					retType = []jen.Code{
						jen.Op("*").Id("CStr"),
					}
					body = append(body, jen.Return(jen.Id("wrapCStr").Params(jen.Id("ret"))))
				} else if iv.Name == "uint8_t" {
					retType = []jen.Code{
						jen.Qual("unsafe", "Pointer"),
					}
					body = append(
						body,
						jen.Return(jen.Qual("unsafe", "Pointer").Params(jen.Id("ret"))),
					)
				} else if _, ok := g.input.structs[iv.Name]; ok {
					retType = []jen.Code{
						jen.Op("*").Id(iv.Name),
					}

					body = append(
						body,
						jen.Var().Id("retMapped").Op("*").Id(iv.Name),
						jen.If(jen.Id("ret").Op("!=").Id("nil")).Block(
							jen.Id("retMapped").Op("=").Op("&").Id(iv.Name).Values(jen.Dict{
								jen.Id("ptr"): jen.Id("ret"),
							}),
						),
						jen.Return(jen.Id("retMapped")),
					)
				} else {
					retType = []jen.Code{
						jen.Op("*").Id(iv.Name),
					}
					body = append(body, jen.Return(jen.Id("ret")))
				}

			default:
				log.Panicln("unexpected type", reflect.TypeOf(v.Inner))
			}

		default:
			log.Panicln("unexpected type", reflect.TypeOf(fn.Result))
		}

		o.Commentf("%v wraps %v.", goName, fn.Name)

		if fn.Comment != "" {
			o.Comment(fn.Comment)
		}

		o.Func().
			//Params(jen.Id("s").Op("*").Id(goName)).
			Id(goName).
			Params(params...).
			Add(retType...).
			Block(body...)
	}

	err := o.Save("functions.gen.go")
	if err != nil {
		log.Panicln(err)
	}
}

func convParamName(val string) string {
	val = strcase.ToLowerCamel(val)

	if val == "type" {
		val = fmt.Sprintf("_%v", val)
	}

	return val
}

var acronyms = []string{
	"av", "hw", "lib", "ff", "io", "api",
}

func (g *Generator) convPart(val string) string {
	val = strings.ToLower(val)

	removed := true

	var prefixes []string

	for removed {
		removed = false

		for _, acronym := range acronyms {
			if strings.HasPrefix(val, acronym) {
				val = strings.TrimPrefix(val, acronym)
				prefixes = append(prefixes, acronym)

				removed = true
			}
		}
	}

	if len(val) > 0 {
		a := val[0:1]
		b := val[1:]
		val = strings.ToUpper(a) + b
	}

	for i := len(prefixes) - 1; i >= 0; i-- {
		val = strings.ToUpper(prefixes[i]) + val
	}

	return val
}

var digitRegex = regexp.MustCompile(`(\d)`)

func (g *Generator) convCamel(val string) string {
	divs := digitRegex.ReplaceAllString(val, "${1}_")
	parts := strings.Split(divs, "_")

	var newParts []string

	for _, part := range parts {
		newParts = append(newParts, g.convPart(part))
	}

	res := strings.Join(newParts, "")

	// Temporary hack
	if _, ok := g.input.structs[res]; ok {
		res = fmt.Sprintf("%v_", res)
	}

	return res
}
