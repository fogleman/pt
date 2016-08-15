package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/fogleman/pt/pt"
)

func main() {
	log.SetFlags(0)
	args := os.Args[1:]
	if len(args) < 3 || len(args)%2 != 1 {
		log.Fatalln("Usage: combine path1.png weight1 path2.png weight2 ... output.png")
	}
	var images []image.Image
	var weights []int64
	var totalWeight int64
	for i := 0; i < len(args)-1; i += 2 {
		path := args[i]
		weight, err := strconv.ParseInt(args[i+1], 0, 0)
		if err != nil {
			log.Fatalf("Invalid weight: %s\n", args[i+1])
		}
		fmt.Printf("Loading %s...\n", path)
		im, err := pt.LoadImage(path)
		if err != nil {
			log.Fatalf("Error loading %s: %v", path, err)
		}
		images = append(images, im)
		weights = append(weights, weight)
		totalWeight += weight
	}
	for i := 1; i < len(images); i++ {
		s1 := images[0].Bounds().Max
		s2 := images[i].Bounds().Max
		if s1 != s2 {
			log.Fatalf("Images must all have the same dimensions")
		}
	}
	fmt.Println("Combining images...")
	size := images[0].Bounds().Max
	w, h := size.X, size.Y
	result := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			fr, fg, fb := 0.0, 0.0, 0.0
			for i, im := range images {
				weight := float64(weights[i])
				r, g, b, _ := im.At(x, y).RGBA()
				fr += weight * float64(r) / 65535
				fg += weight * float64(g) / 65535
				fb += weight * float64(b) / 65535
			}
			fr /= float64(totalWeight)
			fg /= float64(totalWeight)
			fb /= float64(totalWeight)
			r := uint8(math.Min(255, fr*255))
			g := uint8(math.Min(255, fg*255))
			b := uint8(math.Min(255, fb*255))
			result.SetRGBA(x, y, color.RGBA{r, g, b, 255})
		}
	}
	path := args[len(args)-1]
	if err := pt.SavePNG(path, result); err != nil {
		log.Fatalf("Error saving %s: %v", path, err)
	}
}
