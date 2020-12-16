package main

import "math"

func Per(rec Rectangle) float64 {
	return 2 * (rec.L + rec.A)
}

func Are(rec Rectangle) float64 {
	return rec.L * rec.A
}

type Rectangle struct {
	L float64
	A float64
}
type Circle struct {
	Thunder float64
}

func (r Rectangle) Area() float64 {
	return r.L * r.A
}

func (c Circle) Area() float64 {
	return math.Pi * c.Thunder * c.Thunder
}

type Form interface {
	Area() float64
}
