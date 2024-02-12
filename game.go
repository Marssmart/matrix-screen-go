package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"matrix-screen-go/objects"
	"matrix-screen-go/services"
	"matrix-screen-go/static"
)

type Game struct {
	layers []objects.Layer
}

func NewGame() *Game {
	container := services.NewServiceContainer()

	layers := make([]objects.Layer, 6)

	layers[0] = objects.NewLayer(5, 0, 0.1, 1, services.Opacity20, container)
	layers[1] = objects.NewLayer(1, 0, 0.2, 1.1, services.Opacity40, container)
	layers[2] = objects.NewLayer(3, 0, 0.3, 1, services.Opacity60, container)
	layers[3] = objects.NewLayer(0, 0, 0.5, 1.25, services.Opacity80, container)
	layers[4] = objects.NewLayer(0, static.ResolutionHeight/2, 0.7, 1, services.Opacity100, container)
	layers[5] = objects.NewLayer(0, static.ResolutionHeight/3, 0.8, 0.95, services.Opacity100, container)

	return &Game{layers}
}

func (g *Game) Update() error {
	for _, l := range g.layers {
		err := l.Update()
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, l := range g.layers {
		l.Draw(screen)
	}
}

func (g *Game) Layout(int, int) (screenWidth int, screenHeight int) {
	return static.ResolutionWidth, static.ResolutionHeight
}
