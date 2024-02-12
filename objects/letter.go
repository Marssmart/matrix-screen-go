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
	x             float64
	y             float64
	last          *imageRef
	trail         []*trail
	lastUsed      int
	scale         float64
	speed         float64
	speedVariance int
	opacity       services.Opacity
	options       *ebiten.DrawImageOptions
	container     services.ServiceContainer
}

type trail struct {
	ref     *imageRef
	options *ebiten.DrawImageOptions
	x       float64
	y       float64
}

type imageRef struct {
	image   *ebiten.Image
	key     string
	opacity services.Opacity
}

func NewLetterAtScale(x float64, y float64, scale float64, speed float64, opacity services.Opacity, container services.ServiceContainer) Letter {
	trails := make([]*trail, 0)
	return &letter{
		x, y, nil, trails, 0, scale, speed, int(rand.Int31n(5)), opacity, &ebiten.DrawImageOptions{}, container,
	}
}

func (l *letter) Update() error {
	change := static.SpeedToMovement(1, l.speed) + float64(l.speedVariance)
	hitBottom := l.y+change > static.ResolutionHeight

	if hitBottom {
		l.y = 0
	} else {
		l.y = l.y + change
	}

	return nil
}

func (l *letter) Draw(screen *ebiten.Image) {

	//draw the trail
	for _, memory := range l.trail {
		memory.options = resetScaleAndTranslate(memory.options, l.scale, memory.x, memory.y)
		ref := memory.ref
		l.container.ImageService().DrawWithOpacity(screen, ref.key, memory.options, ref.opacity)
	}

	//draw the lead
	l.options = resetScaleAndTranslate(l.options, l.scale, l.x, l.y)
	if l.lastUsed == 0 || l.lastUsed == 25 {
		key, image := l.container.ImageService().DrawRandom(screen, l.options)
		l.last = &imageRef{image, key, l.opacity}
		l.trail = append([]*trail{{l.last, &ebiten.DrawImageOptions{}, l.x, l.y}}, l.trail...)
		l.lastUsed = 1

		for i := len(l.trail) - 1; i > static.MaxTrailLength-(int(l.opacity)/20); i-- {
			ref := l.trail[i].ref
			ref.opacity = services.LowerOpacity(ref.opacity)
			ref.image = l.container.ImageService().FindByName(ref.key, ref.opacity)
		}
	} else {
		l.container.ImageService().Draw(screen, l.last.image, l.options)
		l.lastUsed = l.lastUsed + 1
	}

	if len(l.trail) == static.MaxTrailLength {
		l.trail = l.trail[:len(l.trail)-1]
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
