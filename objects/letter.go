package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"math/rand"
	"matrix-screen-go/services"
	"matrix-screen-go/static"
)

type Letter interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type letter struct {
	x, y, scale, step                    float64
	last                                 *imageRef
	trailDropLimit, lastUsed, stampLimit int
	trail                                []*trail
	options                              *ebiten.DrawImageOptions
	container                            services.ServiceContainer
	effects                              []Effect
}

type trail struct {
	ref     *imageRef
	options *ebiten.DrawImageOptions
	x, y    float64
}

type imageRef struct {
	image *ebiten.Image
	key   string
}

func NewLetterAtScale(x float64, y float64, scale float64, speed float64, container services.ServiceContainer) Letter {
	trails := make([]*trail, 0)
	step := static.SpeedToMovement(1, speed) + float64(rand.Int31n(5))
	var stampLimit = int(static.IconHeight / math.Round(step)) //+ static.IconSpacingInColumn
	var trailDropLimit = int(rand.Int31n(static.MaxTrailLength) + static.MinTrailLength)
	return &letter{
		x, y, scale, step, nil, trailDropLimit, 0, stampLimit, trails, &ebiten.DrawImageOptions{}, container, []Effect{NewPulsatingEffect()},
	}
}

func (l *letter) Update() error {
	hitBottom := l.y+l.step > static.ResolutionHeight

	if hitBottom {
		l.y = 0
	} else {
		l.y = l.y + l.step
	}

	return nil
}

func (l *letter) Draw(screen *ebiten.Image) {

	//draw the trail
	for _, memory := range l.trail {
		resetScaleAndTranslate(memory.options, l.scale, memory.x, memory.y)
		ref := memory.ref
		memory.options.ColorScale.ScaleAlpha(0.995 - float32(len(l.trail))*0.0004)
		l.container.ImageService().Draw(screen, ref.image, memory.options)
	}

	//draw the lead
	resetScaleAndTranslate(l.options, l.scale, l.x, l.y)
	if l.lastUsed == 0 || l.lastUsed == l.stampLimit {
		key, image := l.container.ImageService().DrawRandom(screen, l.options)
		l.last = &imageRef{image, key}
		l.trail = append([]*trail{{l.last, &ebiten.DrawImageOptions{}, l.x, l.y}}, l.trail...)
		l.lastUsed = 1

		for i := len(l.trail) - 1; i > l.trailDropLimit; i-- {
			ref := l.trail[i].ref
			ref.image = l.container.ImageService().FindByName(ref.key)
		}
	} else {
		l.container.ImageService().Draw(screen, l.last.image, l.options)
		l.lastUsed = l.lastUsed + 1
	}

	if len(l.trail) == l.trailDropLimit {
		l.trail = l.trail[:len(l.trail)-1]
	}
}

func resetScaleAndTranslate(options *ebiten.DrawImageOptions, scale float64, x float64, y float64) {
	options.GeoM.Reset()
	if scale != 1 {
		options.GeoM.Scale(scale, scale)
	}
	options.GeoM.Translate(x, y)
}
