package main

import (
	"image/color"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ball struct {
	Object
}

func (b *Ball) UpdatePos(pos Vector) {
	b.Pos = pos
	b.Op.GeoM.Reset()
	b.Op.GeoM.Translate(b.Pos.X, b.Pos.Y)
}

func NewBall() *Ball {
	img := ebiten.NewImage(40, 40)
	ebitenutil.DrawRect(img, 0, 0, 40, 40, color.White)
	dir := Vector{
		X: 1,
		Y: 1,
	}
	pos := Vector{
		X: 200,
		Y: 45,
	}
	return &Ball{
		Object: Object{
			img: img,
			Width:  40,
			Height: 40,
			Op:     &ebiten.DrawImageOptions{},
			Speed:  10,
			Dir:    dir,
			Pos:    pos,
		},
	}
}
