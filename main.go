package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

var (
	screenWidth  int
	screenHeight int
	ball         *Ball
	game         *Game
)

// Game implements ebiten.Game interface.
type Game struct {
	Bars []*Bar
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.

	// walls
	if ball.Pos.Y+float64(ball.Height) >= float64(screenHeight) || ball.Pos.Y <= float64(screenWidth)*0 {
		ball.Dir.Y = ball.Dir.Y * -1
	}
	if ball.Pos.X+float64(ball.Height) >= float64(screenWidth) || ball.Pos.X <= float64(screenWidth)*0 {
		ball.Dir.X = ball.Dir.X * -1
	}

	for _, bar := range game.Bars {
		if ball.IsColliding(bar.Object) {
			if ball.IsCollidingFromLeft(bar.Object) {
				fmt.Println("from left")
				ball.Dir.X = ball.Dir.X * -1
			} else if ball.IsCollidingFromRight(bar.Object) {
				fmt.Println("from right")
				ball.Dir.X = ball.Dir.X * -1
			} else if ball.IsCollidingFromTop(bar.Object) {
				fmt.Println("from top")
				ball.Dir.Y = ball.Dir.Y * -1
			} else if ball.IsCollidingFromBottom(bar.Object) {
				fmt.Println("from bottom")
				ball.Dir.Y = ball.Dir.Y * -1
			}
		}
	}

	dirSpeed := ball.Dir.normalize().multiply(ball.Speed)
	ball.UpdatePos(ball.Pos.add(dirSpeed))
	game.Bars[0].UpdatePos(Vector{X: 0, Y: 0})
	game.Bars[1].UpdatePos(Vector{X: float64(screenWidth) - float64(game.Bars[1].Width), Y: 0})

	// fmt.Printf("%+v\n", ball.Pos)
	return nil

}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})

	screen.DrawImage(ball.img, ball.Op)

	for _, bar := range game.Bars {
		screen.DrawImage(bar.img, bar.Op)
	}

	// fmt.Println(ebiten.CursorPosition())
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1800, 900
}

func main() {
	game = &Game{}

	screenWidth = 1800
	screenHeight = 900

	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Bang")

	ball = NewBall()
	game.Bars = make([]*Bar, 2)
	game.Bars[0] = NewBar(Vector{X: 0, Y: 0}, color.Black)
	game.Bars[1] = NewBar(Vector{X: 0, Y: 0}, color.White)

	fmt.Printf("%+v ||||||||| %+v\n", game.Bars[0], game.Bars[1])

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
