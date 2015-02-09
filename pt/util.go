package pt

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"path"
	"strconv"
	"time"
)

// TODO: don't export this stuff

func Radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func Degrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

func SavePNG(path string, im image.Image) error {
	fmt.Printf("Writing %s... ", path)
	defer fmt.Println("OK")
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, im)
}

func LoadJPG(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return jpeg.Decode(file)
}

func LoadPNG(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return png.Decode(file)
}

func Median(items []float64) float64 {
	n := len(items)
	switch {
	case n == 0:
		return 0
	case n%2 == 1:
		return items[n/2]
	default:
		a := items[n/2-1]
		b := items[n/2]
		return (a + b) / 2
	}
}

func DurationString(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%d:%02d:%02d", h, m, s)
}

func ParseFloats(items []string) []float64 {
	result := make([]float64, len(items))
	for i, item := range items {
		f, _ := strconv.ParseFloat(item, 64)
		result[i] = f
	}
	return result
}

func ParseInts(items []string) []int {
	result := make([]int, len(items))
	for i, item := range items {
		f, _ := strconv.ParseInt(item, 0, 0)
		result[i] = int(f)
	}
	return result
}

func RelativePath(path1, path2 string) string {
	dir, _ := path.Split(path1)
	return path.Join(dir, path2)
}

func Fract(x float64) float64 {
	_, x = math.Modf(x)
	return x
}
