package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

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
			// exercise 3.6
			z1 := mandelbrot(complex(x, y))
			z2 := mandelbrot(complex(x, (float64(py)+0.5)/height*(ymax-ymin)+ymin))
			z3 := mandelbrot(complex((float64(px)+0.5)/width*(xmax-xmin)+xmin, y))
			z4 := mandelbrot(complex((float64(px)+0.5)/width*(xmax-xmin)+xmin, (float64(py)+0.5)/height*(ymax-ymin)+ymin))
			rAvg := uint8((z1.R + z2.R + z3.R + z4.R) / 4)
			gAvg := uint8((z1.G + z2.G + z3.G + z4.G) / 4)
			bAvg := uint8((z1.B + z2.B + z3.B + z4.B) / 4)
			zAvg := color.RGBA{rAvg, gAvg, bAvg, 255}
			// Image point (px, py) represents complex value z.
			img.Set(px, py, zAvg)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.RGBA {
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
	return color.RGBA{0, 0, 0, 255}
}
