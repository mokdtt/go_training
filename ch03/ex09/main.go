// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"os"
	"strconv"
)

var (
	xmin, ymin, xmax, ymax = -2., -2., +2., +2.
	width, height          = 256, 256
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if query, ok := r.URL.Query()["x"]; ok {
			xQuery, err := strconv.Atoi(query[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Got query %s:  %v\n", query, err)
				os.Exit(1)
			}
			if xQuery < 0 {
				fmt.Fprintf(os.Stderr, "Got negative value %s:  %v\n", query)
				os.Exit(1)
			}
			xmin = -float64(xQuery)
			xmax = float64(xQuery)
		}
		if query, ok := r.URL.Query()["y"]; ok {
			yQuery, err := strconv.Atoi(query[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Got query %s:  %v\n", query, err)
				os.Exit(1)
			}
			if yQuery < 0 {
				fmt.Fprintf(os.Stderr, "Got negative value %s:  %v\n", query)
				os.Exit(1)
			}
			ymin = -float64(yQuery)
			ymax = float64(yQuery)
		}
		if query, ok := r.URL.Query()["scale"]; ok {
			scale, err := strconv.Atoi(query[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Got query %s:  %v\n", query, err)
				os.Exit(1)
			}
			width *= scale
			height *= scale
		}
		plot(w)
	}
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func plot(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{50, 255 - contrast*n, 50, 255}
		}
	}
	return color.Black
}
