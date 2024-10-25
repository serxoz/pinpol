package pinpol

import (
	"math"
)

// INF Define Infinite
var INF float64 = 10000

// Point type for points...
type Point struct {
	X float64
	Y float64
}

// onSegment checks if a point is on "pr" vector
func onSegment(p Point, q Point, r Point) bool {
	// floats because math.Max() and math.Min()
	if float64(q.X) <= math.Max(float64(p.X), float64(r.X)) &&
		float64(q.X) >= math.Min(float64(p.X), float64(r.X)) &&
		float64(q.Y) <= math.Max(float64(p.Y), float64(r.Y)) &&
		float64(q.Y) >= math.Min(float64(p.Y), float64(r.Y)) {
		return true
	}
	return false
}

// orientation find the orientation of 3 points (p, q, r).
// returns:
// 0 -> p, q and r are collinear
// 1 -> clockwise
// 2 -> counterclockwise
func orientation(p Point, q Point, r Point) int {
	val := (q.Y-p.Y)*(r.X-q.X) - (q.X-p.X)*(r.Y-q.Y)

	// collinear
	if val == 0 {
		return 0
	}
	// clock or counterclock wise
	if val > 0 {
		return 1
	}
	return 2
}

// doIntersect returns true if line segment "p1q1" and "p2q2" intersect.
func doIntersect(p1 Point, q1 Point, p2 Point, q2 Point) bool {
	// Find the four orientations needed for general and
	// special cases
	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
	o3 := orientation(p2, q2, p1)
	o4 := orientation(p2, q2, q1)

	// General case
	if o1 != o2 && o3 != o4 {
		return true
	}

	// Special Cases
	// p1, q1 and p2 are collinear and p2 lies on segment p1q1
	if o1 == 0 && onSegment(p1, p2, q1) {
		return true
	}

	// p1, q1 and p2 are collinear and q2 lies on segment p1q1
	if o2 == 0 && onSegment(p1, q2, q1) {
		return true
	}

	// p2, q2 and p1 are collinear and p1 lies on segment p2q2
	if o3 == 0 && onSegment(p2, p1, q2) {
		return true
	}

	// p2, q2 and q1 are collinear and q1 lies on segment p2q2
	if o4 == 0 && onSegment(p2, q1, q2) {
		return true
	}

	return false // Doesn't fall in any of the above cases
}

// IsInside returns true if the point p lies inside the polygon[] with n vertices
func IsInside(polygon []Point, n int, p Point) bool {

	// There must be at least 3 vertices in polygon[]
	if n < 3 {
		return false
	}

	// Create a point for line segment from p to infinite
	extreme := Point{X: INF, Y: p.Y}

	// Count intersections of the above line with sides of polygon
	count := 0
	i := 0
	for {
		next := (i + 1) % n

		// Check if the line segment from 'p' to 'extreme' intersects
		// with the line segment from 'polygon[i]' to 'polygon[next]'
		if doIntersect(polygon[i], polygon[next], p, extreme) {
			// If the point 'p' is collinear with line segment 'i-next',
			// then check if it lies on segment. If it lies, return true,
			// otherwise false
			if orientation(polygon[i], p, polygon[next]) == 0 {
				return onSegment(polygon[i], p, polygon[next])
			}

			count++
		}
		i = next

		if i == 0 {
			break
		}
	}

	// Return true if count is odd, false otherwise
	return count%2 == 1
}
