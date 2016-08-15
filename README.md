# pt: a golang path tracer

[![Build Status](https://travis-ci.org/fogleman/pt.png?branch=master)](https://travis-ci.org/fogleman/pt) [![GoDoc](https://godoc.org/github.com/fogleman/pt/pt?status.svg)](https://godoc.org/github.com/fogleman/pt/pt)

This is a CPU-only, unidirectional [path tracing](http://en.wikipedia.org/wiki/Path_tracing) engine written in Go. The Go gopher below was rendered using it, and [here's the code](https://github.com/fogleman/pt/blob/master/examples/gopher.go) that was used to do it. The gopher 3D model was found [here](https://github.com/golang-samples/gopher-3d).

![Go Gopher](http://i.imgur.com/PhUUcTe.png)

### Installation

    go get -u github.com/fogleman/pt/pt

### Examples

The are [lots of examples](https://github.com/fogleman/pt/tree/master/examples) to learn from! To try them, just run, e.g.

    cd go/src/github.com/fogleman/pt
    go run examples/gopher.go

### Features

* Supports OBJ and STL
* Supports textures, bump maps and normal maps
* Supports volume rendering from image slices
* Supports various material properties
* Supports configurable depth of field
* Supports iterative rendering
* Uses k-d trees to accelerate ray intersection tests
* Uses all CPU cores in parallel
* 100% pure Go with no dependencies besides the standard library

### TODO

Here are things that I'm hoping to add.

* bidirectional path tracing
* true BRDFs
* subsurface scattering
* atmosphere
* constructive solid geometry

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

![Go](http://i.imgur.com/LMNUoaM.jpg)

![Cornell](https://www.michaelfogleman.com/static/gallery/853.png)

![Spheres](https://www.michaelfogleman.com/static/gallery/dof.png)

![Suzanne](http://i.imgur.com/iw32US1.png)

![Dragon](http://i.imgur.com/woBoPFx.png)
