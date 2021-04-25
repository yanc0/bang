package main

import (
	"fmt"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type MovableObject interface {
	UpdatePos(Vector)
	IsColliding(Object)
}

type Object struct {
	img    *ebiten.Image
	Op     *ebiten.DrawImageOptions
	Width  int
	Height int
	Dir    Vector
	Pos    Vector
	Speed  float64
}

func (a Object) IsColliding(b Object) bool {
	return a.Pos.X <= b.Pos.X+float64(b.Width) && a.Pos.X+float64(a.Width) >= b.Pos.X &&
		a.Pos.Y <= b.Pos.Y+float64(b.Height) && a.Pos.Y+float64(a.Height) >= b.Pos.Y
}

func (a Object) IsCollidingFromLeft(b Object) bool {
	fmt.Println(a.Pos.X, b.Pos.X+float64(b.Width))
	return a.Pos.X <= b.Pos.X+float64(b.Width)
}

func (a Object) IsCollidingFromRight(b Object) bool {
	fmt.Println(a.Pos.X+float64(a.Width),  b.Pos.X)
	return a.Pos.X+float64(a.Width) >= b.Pos.X
}

func (a Object) IsCollidingFromTop(b Object) bool {
	return a.Pos.Y >= b.Pos.Y+float64(b.Height)
}

func (a Object) IsCollidingFromBottom(b Object) bool {
	return a.Pos.Y+float64(a.Height) >= b.Pos.Y
}
