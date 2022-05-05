# Point IN POLygon

This package checks if a point is inside a polygon.

## Example
```golang
package main

import (
	"fmt"

	pin "github.com/serxoz/pinpol"
)

func main() {
	// Point
	p := pin.Point{X: 2, Y: 2}

	// Polygon
	var poly []pin.Point
	poly = append(poly, pin.Point{X: 1, Y: 1})
	poly = append(poly, pin.Point{X: 1, Y: 3})
	poly = append(poly, pin.Point{X: 3, Y: 3})
	poly = append(poly, pin.Point{X: 3, Y: 1})

	// Check if point "p" is inside polygon "poly"
	in := pin.IsInside(poly, len(poly), p)

	fmt.Println(in)
}
```

## Doc
Based in [Patrick Prosser](http://www.dcs.gla.ac.uk/~pat/) [Geometric Algorithms](http://www.dcs.gla.ac.uk/~pat/52233/slides/Geometry1x1.pdf)
saved in doc/ for the future :D
