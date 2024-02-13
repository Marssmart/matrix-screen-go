package objects

import "github.com/hajimehoshi/ebiten/v2"

type Effect interface {
	Apply(options *ebiten.DrawImageOptions)
}
