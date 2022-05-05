package pinpol

import (
	"testing"
)

// test orientation function
func TestOrientation(t *testing.T) {
	/*
		Given this points:
		  y
		2 │  q
		  │
		1 │  p  r  z
		  │
		  └───────────x
		     1  2  3
	*/
	p := Point{x: 1, y: 1}
	q := Point{x: 1, y: 2}
	r := Point{x: 2, y: 1}
	z := Point{x: 3, y: 1}

	o0 := orientation(p, r, z) // collinear
	o1 := orientation(p, q, r) // clockwise 1
	o2 := orientation(p, r, q) // counterclockwise 2

	if o0 != 0 || o1 != 1 || o2 != 2 {
		t.Errorf("Error on orientation() function.")
	}
}

// test doIntersect function
func TestDoIntersect(t *testing.T) {
	/*
		Given this points:
		y
		2 │  q     z
		│
		1 │  p  r
		│
		└───────────x
		1  2  3
	*/
	p := Point{x: 1, y: 1}
	q := Point{x: 1, y: 2}
	r := Point{x: 2, y: 1}
	z := Point{x: 3, y: 2}

	// vector1 = p->z
	// vector2 = q->r
	i1 := doIntersect(p, z, q, r) // intersect

	// vector1 = p->q
	// vector2 = r->z
	i2 := doIntersect(p, q, r, z) // not intersect

	if i1 != true || i2 != false {
		t.Errorf("Error on doIntersect() function.")
	}
}

// func IsInside(polygon []Point, n int, p Point) bool {
// test IsInside function
func TestIsInside(t *testing.T) {
	/*
		Given this points:
		y
		3 │  q     z
		2 │     r
		1 │  p     y
		│
		└───────────x
		1  2  3
	*/
	p := Point{x: 1, y: 1}
	q := Point{x: 1, y: 3}
	r := Point{x: 2, y: 2}
	z := Point{x: 3, y: 3}
	y := Point{x: 3, y: 1}

	var polygon1 []Point
	polygon1 = append(polygon1, p)
	polygon1 = append(polygon1, q)
	polygon1 = append(polygon1, z)
	polygon1 = append(polygon1, y)
	inside1 := IsInside(polygon1, 4, r) // is inside

	var polygon2 []Point
	polygon2 = append(polygon2, p)
	polygon2 = append(polygon2, q)
	polygon2 = append(polygon2, y)
	inside2 := IsInside(polygon2, 3, z) // is not inside

	if inside1 != true || inside2 != false {
		t.Errorf("Error on IsInside() function.")
	}
}
