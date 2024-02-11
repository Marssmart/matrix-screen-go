package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"matrix-screen-go/static"
)

func main() {

	ebiten.SetWindowTitle("Hello, matrix!")
	ebiten.SetWindowSize(static.ScreenWidth, static.ScreenHeight)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
