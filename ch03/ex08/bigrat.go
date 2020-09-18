// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
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
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(x, y))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(x float64, y float64) color.Color {
	const iterations = 3
	const contrast = 15

	zx := new(big.Rat).SetFloat64(x)
	zy := new(big.Rat).SetFloat64(y)
	vx := new(big.Rat).SetFloat64(0.0)
	vy := new(big.Rat).SetFloat64(0.0)
	bf4 := new(big.Rat).SetFloat64(4.0)

	for n := uint8(0); n < iterations; n++ {
		vx, vy = step(vx, vy, zx, zy)
		if brNorm2(vx, vy).Cmp(bf4) == 1 {
			return color.RGBA{50, 255 - contrast*n, 50, 255}
		}
	}
	return color.Black
}

func step(vx *big.Rat, vy *big.Rat, zx *big.Rat, zy *big.Rat) (*big.Rat, *big.Rat) {
	bf2 := new(big.Rat).SetFloat64(2.0)
	vxtmp := new(big.Rat).Sub(new(big.Rat).Mul(vx, vx), new(big.Rat).Mul(vy, vy))
	vytmp := new(big.Rat).Mul(new(big.Rat).Mul(bf2, vx), vy)
	vxtmp = new(big.Rat).Add(vxtmp, zx)
	vytmp = new(big.Rat).Add(vytmp, zy)
	return vxtmp, vytmp
}

func brNorm2(vx *big.Rat, vy *big.Rat) *big.Rat {
	tmp := new(big.Rat).Add(new(big.Rat).Mul(vx, vx), new(big.Rat).Mul(vy, vy))
	return tmp
}
