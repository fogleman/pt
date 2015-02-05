### Summary

I am writing a [path tracer](http://en.wikipedia.org/wiki/Path_tracing) in Go.

![Sample](http://i.imgur.com/zHRmmeP.png)

    1280 x 720 pixels, 25 x 64 = 1600 samples per pixel, 8 bounces max
    3 shapes, 1 light
    Render time: 3m52s on 3.4 GHz Intel Core i5 (4 cores)

### TODO

Here are things that I'm planning, or at least hoping, to do.

* material properties *(partially done)*
* OBJ / STL models *(partially done)*
* textured objects *(partially done)*
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
* http://graphics.stanford.edu/courses/cs148-10-summer/docs/2006--degreve--reflection_refraction.pdf
* http://www.rorydriscoll.com/2009/01/07/better-sampling/
* https://www.youtube.com/playlist?list=PLslgisHe5tBPckSYyKoU3jEA4bqiFmNBJ
* http://www.amazon.com/Physically-Based-Rendering-Second-Edition/dp/0123750792

### Samples

![Sample](http://i.imgur.com/eI5yLu7.png)

![Sample](http://i.imgur.com/2PNvTgE.png)

![Sample](http://i.imgur.com/2bVB9PL.png)

![Sample](http://i.imgur.com/zCPDKbt.png)

![Sample](http://i.imgur.com/7nJieKd.png)
