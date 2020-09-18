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
	const iterations = 100
	const contrast = 15

	zx := new(big.Float).SetFloat64(x)
	zy := new(big.Float).SetFloat64(y)
	vx := new(big.Float).SetFloat64(0.0)
	vy := new(big.Float).SetFloat64(0.0)
	bf2 := new(big.Float).SetFloat64(2.0)

	for n := uint8(0); n < iterations; n++ {
		vx, vy = step(vx, vy, zx, zy)
		if bfAbs(vx, vy).Cmp(bf2) == 1 {
			return color.RGBA{50, 255 - contrast*n, 50, 255}
		}
	}
	return color.Black
}

func step(vx *big.Float, vy *big.Float, zx *big.Float, zy *big.Float) (*big.Float, *big.Float) {
	bf2 := new(big.Float).SetFloat64(2.0)
	vxtmp := new(big.Float).Sub(new(big.Float).Mul(vx, vx), new(big.Float).Mul(vy, vy))
	vytmp := new(big.Float).Mul(new(big.Float).Mul(bf2, vx), vy)
	vxtmp = new(big.Float).Add(vxtmp, zx)
	vytmp = new(big.Float).Add(vytmp, zy)
	return vxtmp, vytmp
}

func bfAbs(vx *big.Float, vy *big.Float) *big.Float {
	tmp := new(big.Float).Add(new(big.Float).Mul(vx, vx), new(big.Float).Mul(vy, vy))
	tmp = tmp.Sqrt(tmp)
	return tmp
}
