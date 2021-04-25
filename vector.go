package main

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v Vector) multiply(coef float64) Vector {
	return Vector{
		X: v.X * coef,
		Y: v.Y * coef,
	}
}

func (v Vector) add(b Vector) Vector {
	return Vector{
		X: v.X + b.X,
		Y: v.Y + b.Y,
	}
}
func (v *Vector) pos() (x, y int)   { return int(math.Floor(v.X + 0.5)), int(math.Floor(v.Y + 0.5)) }
func (v *Vector) normalize() Vector { l := v.len(); return Vector{v.X / l, v.Y / l} }
func (v *Vector) len() float64      { return math.Sqrt(v.X*v.X + v.Y*v.Y) }
