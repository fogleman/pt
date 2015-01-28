### Summary

I am writing a [path tracer](http://en.wikipedia.org/wiki/Path_tracing) in Go.
I just started, so there isn't much to see yet.

![Sample](http://i.imgur.com/J17Py6l.png)

2560 x 1440 pixels, 16 x 64 = 1024 samples per pixel, 8 bounces max
Render time: 32m42s on 3.4 GHz Intel Core i5 (4 cores)

### TODO

Here are things that I'm planning, or at least hoping, to do.

* material properties
* more shapes
* more example scenes
* OBJ / STL models
* textured objects
* constructive solid geometry
* depth of field
* input files to define scene
* animation support

### Links

Here are some resources that I have found useful.

* http://madebyevan.com/webgl-path-tracing/
* http://www.thepolygoners.com/tutorials/GIIntro/GIIntro.htm
* http://www.iquilezles.org/www/articles/simplepathtracing/simplepathtracing.htm
* http://web.cs.wpi.edu/~emmanuel/courses/cs563/write_ups/zackw/realistic_raytracing.html
* http://www.rorydriscoll.com/2009/01/07/better-sampling/
* https://www.youtube.com/playlist?list=PLslgisHe5tBPckSYyKoU3jEA4bqiFmNBJ
* http://www.amazon.com/Physically-Based-Rendering-Second-Edition/dp/0123750792

### Samples

![Sample](http://i.imgur.com/zCPDKbt.png)

![Sample](http://i.imgur.com/7nJieKd.png)
