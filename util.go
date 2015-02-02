package pt

import (
	"image"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"sort"
)

func Radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func Degrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

func Save(path string, im image.Image) error {
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

func Distinct(items []float64) []float64 {
	result := make([]float64, 0)
	seen := make(map[float64]bool)
	for _, item := range items {
		if _, ok := seen[item]; !ok {
			result = append(result, item)
			seen[item] = true
		}
	}
	sort.Float64s(result)
	return result
}
