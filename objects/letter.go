package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"matrix-screen-go/services"
	"matrix-screen-go/static"
)

type Letter interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type letter struct {
	x         float64
	y         float64
	scale     float64
	options   *ebiten.DrawImageOptions
	container services.ServiceContainer
}

func NewLetterAtScale(x float64, y float64, scale float64, container services.ServiceContainer) Letter {
	return &letter{
		x, y, scale, &ebiten.DrawImageOptions{}, container,
	}
}

func (l *letter) Update() error {
	change := static.NormalSpeed(1)
	hitBottom := l.y+change > static.ScreenHeight

	if hitBottom {
		l.y = 0
	} else {
		l.y = l.y + change
	}

	return nil
}

func (l *letter) Draw(screen *ebiten.Image) {
	l.options.GeoM.Reset()
	if l.scale != 1 {
		l.options.GeoM.Scale(l.scale, l.scale)
	}
	l.options.GeoM.Translate(l.x, l.y)

	l.container.ImageService().DrawRandom(screen, l.options)
}
