# ffmpeg-go
Go bindings for the FFmpeg AV libraries.

### Features

- Static build (does still require some system libraries, such as `m`).
- Linux (arm64, amd64) and macOS (arm64, amd64) support.
- Auto generated.
- Thin layer over the C API.
- Doc comments retained (where supported by clang).
- GPL version of FFmpeg (allows for x264 etc).

### TODO

- [ ] Expose more headers.
- [ ] Ensure acceleration enabled in builds.
- [ ] Expose platform specific headers.
- [ ] Cleanup internal packages.

### Setup

```
go get github.com/csnewman/ffmpeg-go
```

### Example

![example.gif](example.gif)

(The asciiplayer demo playing [Big Buck Bunny](https://en.wikipedia.org/wiki/Big_Buck_Bunny), with some GIF artifacts)

The `examples` directory contains some ported and some custom examples based on the C docs.

An example of printing a file's metadata:

```go
package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/csnewman/ffmpeg-go"
)

func main() {
	slog.Info("Metadata")

	var ctx *ffmpeg.AVFormatContext

	url := ffmpeg.ToCStr(os.Args[1])
	defer url.Free()

	_, err := ffmpeg.AVFormatOpenInput(&ctx, url, nil, nil)
	if err != nil {
		log.Panicln(err)
	}

	defer ffmpeg.AVFormatFreeContext(ctx)

	if _, err := ffmpeg.AVFormatFindStreamInfo(ctx, nil); err != nil {
		log.Panicln(err)
	}

	ffmpeg.AVDumpFormat(ctx, 0, url, 0)

	streams := ctx.Streams()

	for i := uintptr(0); i < uintptr(ctx.NbStreams()); i++ {
		s := streams.Get(i)

		slog.Info("  Stream", "i", i, "id", s.Id(), "dur", s.Duration())

		meta := s.Metadata()

		var entry *ffmpeg.AVDictionaryEntry

		for {
			entry = ffmpeg.AVDictGet(meta, ffmpeg.GlobalCStr(""), entry, ffmpeg.AVDictIgnoreSuffix)
			if entry == nil {
				break
			}

			slog.Info("    Meta", "key", entry.Key(), "value", entry.Value())
		}
	}
}
```

```
2024/01/15 21:39:43 INFO Metadata
Input #0, mov,mp4,m4a,3gp,3g2,mj2, from './bbb.mp4':
  Metadata:
    major_brand     : isom
    minor_version   : 1
    compatible_brands: isomavc1
    creation_time   : 2013-12-16T17:59:32.000000Z
    title           : Big Buck Bunny, Sunflower version
    artist          : Blender Foundation 2008, Janus Bager Kristensen 2013
    comment         : Creative Commons Attribution 3.0 - http://bbb3d.renderfarming.net
    genre           : Animation
    composer        : Sacha Goedegebure
  Duration: 00:10:34.57, start: 0.000000, bitrate: 4486 kb/s
  Stream #0:0[0x1](und): Video: h264 (High) (avc1 / 0x31637661), yuv420p(progressive), 1920x1080 [SAR 1:1 DAR 16:9], 4001 kb/s, 60 fps, 60 tbr, 60k tbn (default)
    Metadata:
      creation_time   : 2013-12-16T17:59:32.000000Z
      handler_name    : GPAC ISO Video Handler
      vendor_id       : [0][0][0][0]
  Stream #0:1[0x2](und): Audio: mp3 (mp4a / 0x6134706D), 48000 Hz, stereo, fltp, 160 kb/s (default)
    Metadata:
      creation_time   : 2013-12-16T17:59:37.000000Z
      handler_name    : GPAC ISO Audio Handler
      vendor_id       : [0][0][0][0]
  Stream #0:2[0x3](und): Audio: ac3 (ac-3 / 0x332D6361), 48000 Hz, 5.1(side), fltp, 320 kb/s (default)
    Metadata:
      creation_time   : 2013-12-16T17:59:37.000000Z
      handler_name    : GPAC ISO Audio Handler
      vendor_id       : [0][0][0][0]
    Side data:
      audio service type: main
2024/01/15 21:39:43 INFO   Stream i=0 id=1 dur=38072000
2024/01/15 21:39:43 INFO     Meta key=creation_time value=2013-12-16T17:59:32.000000Z
2024/01/15 21:39:43 INFO     Meta key=language value=und
2024/01/15 21:39:43 INFO     Meta key=handler_name value="GPAC ISO Video Handler"
2024/01/15 21:39:43 INFO     Meta key=vendor_id value=[0][0][0][0]
2024/01/15 21:39:43 INFO   Stream i=1 id=2 dur=30441600
2024/01/15 21:39:43 INFO     Meta key=creation_time value=2013-12-16T17:59:37.000000Z
2024/01/15 21:39:43 INFO     Meta key=language value=und
2024/01/15 21:39:43 INFO     Meta key=handler_name value="GPAC ISO Audio Handler"
2024/01/15 21:39:43 INFO     Meta key=vendor_id value=[0][0][0][0]
2024/01/15 21:39:43 INFO   Stream i=2 id=3 dur=30438912
2024/01/15 21:39:43 INFO     Meta key=creation_time value=2013-12-16T17:59:37.000000Z
2024/01/15 21:39:43 INFO     Meta key=language value=und
2024/01/15 21:39:43 INFO     Meta key=handler_name value="GPAC ISO Audio Handler"
2024/01/15 21:39:43 INFO     Meta key=vendor_id value=[0][0][0][0]
```

### FAQ

- Q: Why is FFmpeg statically bundled?
  - A: The version of FFmpeg varies by platform, making a one-size fits all style binding difficult with the Go build
    system and its lack of a build.go file.
- Q: Why are the arm packages not built by CI?
  - A: Ideally they would be. However, the free GitHub plan does not include arm builders.
- Q: Why is library `xyz` not enabled?
  - A: Libraries have been added on a as-needed basis. Feel free to create a PR enabling any additional libraries.
- Q: I have some C code that I can't convert to Go, due to a missing binding feature
  - A: Create an issue with a C sample, and we can go from there.

### Library versions

| Library  | Version |
|----------|---------|
| FFmpeg   | 6.1     |
| aom      | 3.8.1   |
| ass      | 0.17.1  |
| brotli   | 1.1     |
| bz2      | 1.0.8   |
| freetype | 2.13.2  |
| fribidi  | 1.0.13  |
| harfbuzz | 8.3.0   |
| mp3lame  | 3.1000  |
| ogg      | 1.3.1   |
| opus     | 1.4     |
| png      | head    |
| speex    | 1.2.1   |
| theora   | 1.1.1   |
| unibreak | 5.1     |
| vpx      | 1.14.0  |
| x264     | head    |
| zlib     | 1.3     |
