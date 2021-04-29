package main

import "strconv"

type Point struct {
	coordinateX int
	coordinateY int
}

func (point Point) show() string {
	return strconv.Itoa(point.coordinateX) + ", " + strconv.Itoa(point.coordinateY)
}
