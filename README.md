### Summary

[![Build Status](https://travis-ci.org/fogleman/pt.png)](https://travis-ci.org/fogleman/pt) [![GoDoc](https://godoc.org/github.com/fogleman/pt/pt?status.svg)](https://godoc.org/github.com/fogleman/pt/pt)

I am writing a [path tracer](http://en.wikipedia.org/wiki/Path_tracing) in Go. The Go gopher below was rendered using it, and [here's the code](https://github.com/fogleman/pt/blob/master/examples/gopher.go) that was used to do it. It's 2560x1440px... so you can make it your wallpaper! The gopher 3D model was found [here](https://github.com/golang-samples/gopher-3d).

Disclaimer: This is my first time using Go.

![Go Gopher](http://i.imgur.com/buSF7m5.png)

    go run examples/gopher.go

### Features

* Supports OBJ, Binary STL and ASCII STL
* Supports textures, bump maps and normal maps
* Uses k-d trees to accelerate ray intersection tests
* Supports various material properties
* Light sources with configurable attenuation
* Supports configurable depth of field
* Supports iterative rendering
* Uses all CPU cores in parallel

### TODO

Here are things that I'm planning, or at least hoping, to do.

* subsurface scattering
* atmosphere
* constructive solid geometry
* input files to define scene
* animation support?

### Links

Here are some resources that I have found useful.

* [WebGL Path Tracing - Evan Wallace](http://madebyevan.com/webgl-path-tracing/)
* [Global Illumination in a Nutshell](http://www.thepolygoners.com/tutorials/GIIntro/GIIntro.htm)
* [Simple Path Tracing - Iñigo Quilez](http://www.iquilezles.org/www/articles/simplepathtracing/simplepathtracing.htm)
* [Realistic Raytracing - Zack Waters](http://web.cs.wpi.edu/~emmanuel/courses/cs563/write_ups/zackw/realistic_raytracing.html)
* [Reflections and Refractions in Ray Tracing - Bram de Greve](http://graphics.stanford.edu/courses/cs148-10-summer/docs/2006--degreve--reflection_refraction.pdf)
* [Better Sampling - Rory Driscoll](http://www.rorydriscoll.com/2009/01/07/better-sampling/)
* [Ray Tracing for Global Illumination - Nelson Max at UC Davis](https://www.youtube.com/playlist?list=PLslgisHe5tBPckSYyKoU3jEA4bqiFmNBJ)
* [Physically Based Rendering - Matt Pharr, Greg Humphreys](http://www.amazon.com/Physically-Based-Rendering-Second-Edition/dp/0123750792)

### Samples

![HDRI](http://i.imgur.com/z1SUVrr.png)

![Sponza](http://i.imgur.com/wjNZJPT.png)

![Dragon](http://i.imgur.com/woBoPFx.png)

![Cubes](http://i.imgur.com/Ypn3WCI.png)

![Suzanne](http://i.imgur.com/eI5yLu7.png)

![Cylinders](http://i.imgur.com/yVeil5G.png)

![Dinosaur](http://i.imgur.com/fx8Cgvy.png)

![Balls](http://i.imgur.com/2PNvTgE.png)

![Pencil](http://i.imgur.com/m6drd9s.png)
