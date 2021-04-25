package main

import (
	"image/color"
	"math"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bar struct {
	Object
}

func (b *Bar) UpdatePos(pos Vector) {
	b.Pos = pos
	b.Pos.X = math.Ceil(b.Pos.X)
	b.Pos.Y = math.Ceil(b.Pos.Y)
	b.Op.GeoM.Reset()
	b.Op.GeoM.Translate(b.Pos.X, b.Pos.Y)
}

func NewBar(initialPos Vector, clr color.Color) *Bar {
	img := ebiten.NewImage(200, 400)
	ebitenutil.DrawRect(img, 0, 0, 200, 400, clr)
	dir := Vector{
		X: 1,
		Y: 1,
	}
	return &Bar{
		Object: Object{
			img:    img,
			Width:  200,
			Height: 400,
			Op:     &ebiten.DrawImageOptions{},
			Speed:  20,
			Dir:    dir,
			Pos:    initialPos,
		},
	}
}
