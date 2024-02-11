package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {

	ebiten.SetWindowTitle("Hello, matrix!")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
