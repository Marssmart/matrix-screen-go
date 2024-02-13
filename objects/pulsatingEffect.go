package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

const (
	StartRed   float64 = 55
	StartGreen float64 = 183
	StartBlue  float64 = 69
	EndRed     float64 = 255
	EndGreen   float64 = 255
	EndBlue    float64 = 255
)

type pulsatingEffect struct {
	CurrentRed, CurrentGreen, CurrentBlue float64
	direction                             bool
}

func NewPulsatingEffect() Effect {
	return &pulsatingEffect{StartRed, StartGreen, StartBlue, true}
}

func (p *pulsatingEffect) Apply(options *ebiten.DrawImageOptions) {
	if (p.direction && p.CurrentRed == EndRed) || (!p.direction && p.CurrentRed == StartRed) {
		p.direction = !p.direction
	}

	if p.direction {
		p.CurrentRed = math.Min(p.CurrentRed+1, EndRed)
		p.CurrentGreen = math.Min(p.CurrentGreen+1, EndGreen)
		p.CurrentBlue = math.Min(p.CurrentBlue+1, EndBlue)
	} else {
		p.CurrentRed = math.Max(p.CurrentRed-1, StartRed)
		p.CurrentGreen = math.Max(p.CurrentGreen-1, StartGreen)
		p.CurrentBlue = math.Max(p.CurrentBlue-1, StartBlue)
	}
	options.ColorScale.Scale(float32(p.CurrentRed), float32(p.CurrentGreen), float32(p.CurrentBlue), options.ColorScale.A())
}
