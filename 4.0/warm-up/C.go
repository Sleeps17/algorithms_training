package main

import (
	"fmt"
	"math"
)

func distance(xa, ya, xb, yb int) float64 {
	return math.Sqrt(math.Pow(float64(xb-xa), 2) + math.Pow(float64(yb-ya), 2))
}

func swap(xa, ya, xb, yb *int) {
	tmp := *xa
	*xa = *xb
	*xb = tmp

	tmp = *ya
	*ya = *yb
	*yb = tmp
}

func polarAngle(x, y int) float64 {
	return math.Atan2(float64(y), float64(x))
}

func minPolarAngle(x1, y1, x2, y2 int) float64 {
	angle1 := polarAngle(x1, y1)
	angle2 := polarAngle(x2, y2)

	angleDiff := math.Abs(angle1 - angle2)
	minAngle := math.Min(angleDiff, 2*math.Pi-angleDiff)

	return minAngle
}

func mainC() {
	var xa, ya, xb, yb int
	fmt.Scan(&xa, &ya, &xb, &yb)

	if distance(xa, ya, 0, 0) > distance(xb, yb, 0, 0) {
		swap(&xa, &ya, &xb, &yb)
	}

	ra := distance(xa, ya, 0, 0)
	rb := distance(xb, yb, 0, 0)
	dist1 := ra + rb

	angle := minPolarAngle(xa, ya, xb, yb)
	dist2 := ra*angle + (rb - ra)

	fmt.Printf("%.8f\n", min(dist2, dist1))
}
