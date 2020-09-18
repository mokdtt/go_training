// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
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
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	imgNew := supersampling(img, 2)
	png.Encode(os.Stdout, imgNew) // NOTE: ignoring errors
}

func supersampling(img image.Image, ratio int) image.Image {
	rct := img.Bounds()
	width := rct.Dx()
	height := rct.Dy()
	imgNew := image.NewRGBA(image.Rect(0, 0, width*ratio, height*ratio))
	for py := 0; py < height*2; py++ {
		for px := 0; px < width*2; px++ {
			cx, cy := int(px/2), int(py/2)
			newColor := img.At(cx, cy)
			imgNew.Set(px, py, newColor)
		}
	}
	imgSuper := image.NewRGBA(image.Rect(0, 0, width*ratio, height*ratio))
	for py := 0; py < height*2; py++ {
		for px := 0; px < width*2; px++ {
			newColor := getAveragedColor(imgNew, px, py)
			imgSuper.Set(px, py, newColor)
		}
	}
	return imgSuper
}

func getAveragedColor(img image.Image, cx int, cy int) color.Color {
	rct := img.Bounds()
	width := rct.Dx()
	height := rct.Dy()
	if cx-1 < 0 || cy-1 < 0 || cx+1 > width-1 || cy+1 > height-1 {
		return img.At(cx, cy)
	}
	var rSum, gSum, bSum, aSum int64
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			r, g, b, a := getRGBA(img, cx+i, cy+j)
			rSum += r
			gSum += g
			bSum += b
			aSum += a
		}
	}
	rAve := rSum / 9
	gAve := gSum / 9
	bAve := bSum / 9
	aAve := aSum / 9
	return color.RGBA{uint8(rAve), uint8(gAve), uint8(bAve), uint8(aAve)}
}

func getRGBA(img image.Image, x int, y int) (int64, int64, int64, int64) {
	r, g, b, a := img.At(x, y).RGBA()
	r, g, b, a = r>>8, g>>8, b>>8, a>>8
	return int64(r), int64(g), int64(b), int64(a)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{50, 255 - contrast*n, 50, 200}
		}
	}
	return color.RGBA{205, 92, 92, 200}
}
