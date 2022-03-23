package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

// z^4 - 1 = 0
var roots = []complex128{complex(1, 0), complex(-1, 0), complex(0, 1), complex(0, -1)}
var colors = []color.RGBA{
	{255, 0, 0, 255},
	{0, 255, 0, 255},
	{0, 0, 255, 255},
	{128, 128, 128, 255}}

func main() {
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
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 255
	const contrast = 15

	for n := uint8(0); n < iterations; n++ {
		z -= function(z) / derivative(z)
		for i, root := range roots {
			diff := cmplx.Abs(z - root)
			if math.Abs(diff) <= 0.01 {
				finalColor := colors[i]
				return color.RGBA{finalColor.R, finalColor.G, finalColor.B, n}
			}
		}
	}
	return color.Black
}

func function(z complex128) complex128 {
	return cmplx.Pow(z, 4)
}

func derivative(z complex128) complex128 {
	return 4 * cmplx.Pow(z, 3)
}
