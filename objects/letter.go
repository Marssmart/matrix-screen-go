package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
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
	last      *ebiten.Image
	trail     []*trail
	lastUsed  int
	scale     float64
	speed     float64
	options   *ebiten.DrawImageOptions
	container services.ServiceContainer
}

type trail struct {
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
}

func NewLetterAtScale(x float64, y float64, scale float64, speed float64, container services.ServiceContainer) Letter {
	trails := make([]*trail, rand.Int31n(static.MaxTrailLength)+static.MinTrailLength)
	for i := 0; i < len(trails); i++ {
		trails[i] = &trail{container.ImageService().PickRandom(), &ebiten.DrawImageOptions{}}
	}
	return &letter{
		x, y, nil, trails, 0, scale, speed, &ebiten.DrawImageOptions{}, container,
	}
}

func (l *letter) Update() error {
	change := static.SpeedToMovement(1, l.speed)
	hitBottom := l.y+change-float64(static.IconHeight*len(l.trail)) > static.ResolutionHeight

	if hitBottom {
		l.y = 0
	} else {
		l.y = l.y + change
	}

	return nil
}

func (l *letter) Draw(screen *ebiten.Image) {
	//draw the lead
	l.options = resetScaleAndTranslate(l.options, l.scale, l.x, l.y)
	if l.lastUsed == 0 || l.lastUsed == 25 {
		l.last = l.container.ImageService().DrawRandom(screen, l.options)
		l.lastUsed = 1
	} else {
		l.container.ImageService().Draw(screen, l.last, l.options)
		l.lastUsed = l.lastUsed + 1
	}

	//draw the trail
	for idx, memory := range l.trail[1:] {
		memory.options = resetScaleAndTranslate(memory.options, l.scale, l.x, l.y-float64((idx+1)*(static.IconHeight-static.IconOverlap)))
		l.container.ImageService().Draw(screen, memory.image, memory.options)
	}
}

func resetScaleAndTranslate(options *ebiten.DrawImageOptions, scale float64, x float64, y float64) *ebiten.DrawImageOptions {
	options.GeoM.Reset()
	if scale != 1 {
		options.GeoM.Scale(scale, scale)
	}
	options.GeoM.Translate(x, y)
	return options
}
