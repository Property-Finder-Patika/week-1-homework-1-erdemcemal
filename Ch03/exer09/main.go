package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

// Exercise 3.9: Write a web server that renders fractals and writes the image data to the client.
// Allow the client to specify the x, y, and zoom values as parameters to the HTTP request.
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderFractal(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	return
}

// write a function that parse a string to a float with the default value of 0
func parseFloat(s string, defaultValue float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defaultValue
	}
	return f
}
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.Black
}

func renderFractal(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	const (
		width, height = 1024, 1024
	)
	x := parseFloat(r.FormValue("x"), 0)
	y := parseFloat(r.FormValue("y"), 0)
	zoom := parseFloat(r.FormValue("zoom"), 0)

	m := math.Exp2(1 - zoom)
	xmin, ymin, xmax, ymax := x-m, y-m, x+m, y+m

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	err := png.Encode(w, img)
	if err != nil {
		return
	}
}
