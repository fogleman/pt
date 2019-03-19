package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"log"
	"math"
	"os"

	"github.com/hborntraeger/pt/pt"
)

var Threshold float64
var Count int

func init() {
	flag.Float64Var(&Threshold, "t", 0.2, "speckle threshold")
	flag.IntVar(&Count, "n", 2, "neighbors below threshold")
}

func imageToRGBA(src image.Image) *image.RGBA {
	dst := image.NewRGBA(src.Bounds())
	draw.Draw(dst, dst.Rect, src, image.ZP, draw.Src)
	return dst
}

type Fix struct {
	X, Y  int
	Color color.RGBA
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	args := flag.Args()
	if len(args) != 2 {
		log.Println("usage: despeckle input output [-n COUNT] [-t THRESHOLD]")
		flag.PrintDefaults()
		os.Exit(1)
	}
	image, err := pt.LoadImage(args[0])
	if err != nil {
		log.Fatalf("error reading %s: %v", args[0], err)
	}
	im := imageToRGBA(image)
	w := im.Bounds().Size().X
	h := im.Bounds().Size().Y
	var fixes []Fix
	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			c1 := im.RGBAAt(x, y)
			r1, g1, b1, _ := c1.RGBA()
			var count int
			var tr, tg, tb uint32
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dx == 0 && dy == 0 {
						continue
					}
					c2 := im.RGBAAt(x+dx, y+dy)
					r2, g2, b2, _ := c2.RGBA()
					tr += r2
					tg += g2
					tb += b2
					dr := (float64(r1) - float64(r2)) / 65535
					dg := (float64(g1) - float64(g2)) / 65535
					db := (float64(b1) - float64(b2)) / 65535
					e := math.Sqrt(dr*dr + dg*dg + db*db)
					if e < Threshold {
						// neighbor is not too different
						count++
					}
				}
			}
			if count >= Count {
				// if at least Count neighbors were not too different
				continue
			}
			c := color.RGBA{
				uint8(tr / 8 / 256),
				uint8(tg / 8 / 256),
				uint8(tb / 8 / 256),
				255,
			}
			fixes = append(fixes, Fix{x, y, c})
		}
	}
	for _, fix := range fixes {
		im.SetRGBA(fix.X, fix.Y, fix.Color)
	}
	log.Printf("fixed %d speckles", len(fixes))
	if err := pt.SavePNG(args[1], im); err != nil {
		log.Fatalf("error writing %s: %v", args[1], err)
	}
}
