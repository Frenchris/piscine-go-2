package main

import "fmt"

type ptr struct {
	x int
	y int
}

func setPoint(ptr *ptr) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	point := ptr{}
	points := &point

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
