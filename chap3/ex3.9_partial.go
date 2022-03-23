package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveMandelborot)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func serveMandelborot(w http.ResponseWriter, r *http.Request) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 255
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			var val uint8 = contrast * n
			return color.RGBA{uint8(math.Pow(float64(val), 1)) % 255, uint8(math.Pow(float64(val), 2)) % 255, uint8(math.Pow(float64(val), 3)) % 255, 255}
		}
	}
	return color.Black
}
