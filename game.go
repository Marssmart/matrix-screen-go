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

	layers[0] = objects.NewLayer(0, static.ResolutionHeight/35, 0.1, 1, container)
	layers[1] = objects.NewLayer(static.IconWidth, static.ResolutionHeight/7, 0.2, 1.1, container)
	layers[2] = objects.NewLayer(static.IconWidth*2, static.ResolutionHeight/15, 0.3, 1, container)
	layers[3] = objects.NewLayer(static.IconWidth*0.5, static.ResolutionHeight/9, 0.6, 1.25, container)
	layers[4] = objects.NewLayer(static.IconWidth*1.5, static.ResolutionHeight/2, 0.75, 1, container)
	layers[5] = objects.NewLayer(static.IconWidth/2.5, static.ResolutionHeight/3, 0.85, 0.95, container)

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
