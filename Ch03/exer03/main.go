package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	maxHeight, minHeight := getMaxAndMinHeight()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			color := getPolygonColor(f(ax, ay), minHeight, maxHeight)

			// Skip invalid polygons.
			if isFinite(ax) && isFinite(ay) && isFinite(bx) && isFinite(by) && isFinite(cx) && isFinite(cy) && isFinite(dx) && isFinite(dy) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// isFinite returns true if f is finite.
func isFinite(f float64) bool {
	return !math.IsInf(f, 0) && !math.IsNaN(f)
}

func getMaxAndMinHeight() (float64, float64) {
	maxHeight := math.Inf(-1)
	minHeight := math.Inf(1)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
			z := f(x, y)
			if isFinite(z) {
				if z > maxHeight {
					maxHeight = z
				}
				if z < minHeight {
					minHeight = z
				}
			}
		}
	}
	return maxHeight, minHeight
}

// Exercise 3.3 Color each polygon based on its height,
// so that the peaks are colored red (#ff0000) and the valleys blue (#0000ff).
func getPolygonColor(height, maxHeight, minHeight float64) string {
	if !isFinite(height) || !isFinite(maxHeight) || !isFinite(minHeight) {
		return "#0000FF"
	}
	n := int((height - minHeight) / (maxHeight - minHeight) * 255)
	r := fmt.Sprintf("%02x", n)
	g := "00"
	b := fmt.Sprintf("%02x", 255-n)

	return fmt.Sprintf("#%s%s%s", r, g, b)
}
