package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"matrix-screen-go/services"
	"matrix-screen-go/static"
)

type Layer interface {
	Update() error
	Draw(screen *ebiten.Image)
}

func NewLayer(x float64, y float64, scale float64, speed float64, container services.ServiceContainer) Layer {
	letterCount := static.LetterCount(x, scale)
	letters := make([]Letter, letterCount)
	for i := 0; i < letterCount; i++ {
		if rand.Int31n(8) == 1 {
			letters[i] = NewEmptySpace()
		} else {
			randomShiftY := rand.Int31n(static.ResolutionHeight / 3)
			letters[i] = NewLetterAtScale(x+float64((static.IconWidth-static.IconOverlapInRow)*i), y+float64(i)+float64(randomShiftY), scale, speed, container)
		}
	}
	return &layer{x, y, letters, container}
}

type layer struct {
	x, y      float64
	letters   []Letter
	container services.ServiceContainer
}

func (l *layer) Update() error {
	for _, le := range l.letters {
		err := le.Update()
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *layer) Draw(screen *ebiten.Image) {
	for _, le := range l.letters {
		le.Draw(screen)
	}
}
