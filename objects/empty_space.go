package objects

import "github.com/hajimehoshi/ebiten/v2"

func NewEmptySpace() Letter {
	return &emptySpace{}
}

type emptySpace struct {
}

func (e *emptySpace) Update() error {
	return nil
}
func (e *emptySpace) Draw(screen *ebiten.Image) {

}
