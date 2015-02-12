### Summary

[![Build Status](https://travis-ci.org/fogleman/pt.png)](https://travis-ci.org/fogleman/pt) [![GoDoc](https://godoc.org/github.com/fogleman/pt/pt?status.svg)](https://godoc.org/github.com/fogleman/pt/pt)

I am writing a [path tracer](http://en.wikipedia.org/wiki/Path_tracing) in Go. The Go gopher below was rendered using it, and [here's the code](https://github.com/fogleman/pt/blob/master/examples/gopher.go) that was used to do it. It's 2560x1440px... so you can make it your wallpaper! The gopher 3D model was found [here](https://github.com/golang-samples/gopher-3d).

Disclaimer: This is my first time using Go.

![Go Gopher](http://i.imgur.com/buSF7m5.png)

    go run examples/gopher.go

### TODO

Here are things that I'm planning, or at least hoping, to do.

* bump / normal maps
* subsurface scattering
* atmosphere
* depth of field
* constructive solid geometry
* input files to define scene
* animation support?

### Links

Here are some resources that I have found useful.

* http://madebyevan.com/webgl-path-tracing/
* http://www.thepolygoners.com/tutorials/GIIntro/GIIntro.htm
* http://www.iquilezles.org/www/articles/simplepathtracing/simplepathtracing.htm
* http://web.cs.wpi.edu/~emmanuel/courses/cs563/write_ups/zackw/realistic_raytracing.html
* http://graphics.stanford.edu/courses/cs148-10-summer/docs/2006--degreve--reflection_refraction.pdf
* http://www.rorydriscoll.com/2009/01/07/better-sampling/
* https://www.youtube.com/playlist?list=PLslgisHe5tBPckSYyKoU3jEA4bqiFmNBJ
* http://www.amazon.com/Physically-Based-Rendering-Second-Edition/dp/0123750792

### Samples

![Sponza](http://i.imgur.com/wjNZJPT.png)

![Suzanne](http://i.imgur.com/eI5yLu7.png)

![Pencil](http://i.imgur.com/m6drd9s.png)

![Balls](http://i.imgur.com/2PNvTgE.png)

![Teapot](http://i.imgur.com/2bVB9PL.png)

![Earth](http://i.imgur.com/zCPDKbt.png)

![Balls](http://i.imgur.com/zHRmmeP.png)
