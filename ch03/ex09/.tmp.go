// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

var (
	width, height = 600, 320                     // canvas size in pixels
	cells         = 100                          // number of grid cells
	xyrange       = 30.0                         // axis ranges (-xyrange..+xyrange)
	xyscale       = float64(width) / 2 / xyrange // pixels per x or y unit
	zscale        = float64(height) * 0.4        // pixels per z unit
	angle         = math.Pi / 6                  // angle of x, y axes (=30°)
	color         = "jet"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		if query, ok := r.URL.Query()["height"]; ok {
			heightQuery, err := strconv.Atoi(query[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Got query %s:  %v\n", query, err)
				os.Exit(1)
			}
			height = heightQuery
		}
		if query, ok := r.URL.Query()["width"]; ok {
			widthQuery, err := strconv.Atoi(query[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Got query %s:  %v\n", query, err)
				os.Exit(1)
			}
			width = widthQuery
		}
		if query, ok := r.URL.Query()["color"]; ok {
			color = query[0]
		}
		illustrate(w)
	}
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func illustrate(out io.Writer) {
	zlist := []float64{}
	pointslist := [][]float64{}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)
			if math.IsInf(z, 0) {
				continue
			}
			zlist = append(zlist, z)
			pointslist = append(pointslist, []float64{ax, ay, bx, by, cx, cy, dx, dy, z})
		}
	}
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	zmin, zmax := getMinMax(zlist)
	for _, points := range pointslist {
		ax, ay := points[0], points[1]
		bx, by := points[2], points[3]
		cx, cy := points[4], points[5]
		dx, dy := points[6], points[7]
		z := points[8]
		if color == "jet" {
			r, g, b := getColor(z, zmin, zmax)
			str := fmt.Sprintf("rgb(%d,%d,%d)", r, g, b)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%v;'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, str)
		} else {
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%v;'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}

	}
	fmt.Fprintln(out, "</svg>")
}

func getMinMax(list []float64) (float64, float64) {
	vmin, vmax := list[0], list[0]
	for _, v := range list {
		if v < vmin {
			vmin = v
		}
		if v > vmax {
			vmax = v
		}
	}
	return vmin, vmax
}

func getColor(v, vmin, vmax float64) (int64, int64, int64) {
	if v < vmin {
		v = vmin
	}
	if v > vmax {
		v = vmax
	}
	r, g, b := 1., 1., 1.
	dv := vmax - vmin
	if v < (vmin + 0.15*dv) {
		r = 0
		g = 4 * (v - vmin) / dv
	} else if v < (vmin + 0.3*dv) {
		r = 0
		b = 1 + 4*(vmin+0.25*dv-v)/dv
	} else if v < (vmin + 0.5*dv) {
		r = 4 * (v - vmin - 0.5*dv) / dv
		b = 0
	} else {
		g = 1 + 4*(vmin+0.75*dv-v)/dv
		b = 0
	}
	r = r * 255.
	g = g * 255.
	b = b * 255.
	return int64(r), int64(g), int64(b)
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
