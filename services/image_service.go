package services

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math/rand"
	"matrix-screen-go/static"
)

func NewImageService() ImageService {
	var letters = [15]string{
		static.LetterAlpha,
		static.LetterBeta,
		static.LetterGamma,
		static.LetterH,
		static.LetterHiragana,
		static.LetterKsi,
		static.LetterLambda,
		static.LetterMu,
		static.LetterOmega,
		static.LetterPi,
		static.LetterR,
		static.LetterS,
		static.LetterSigma,
		static.LetterU,
		static.LetterW,
	}

	letterImages := make([]*ebiten.Image, len(letters))

	for i, l := range letters {
		letterImages[i] = loadImage(l)
	}

	return &imageService{images: letterImages}
}

type ImageService interface {
	DrawRandom(screen *ebiten.Image, options *ebiten.DrawImageOptions)
}

type imageService struct {
	images []*ebiten.Image
}

func (i *imageService) DrawRandom(screen *ebiten.Image, options *ebiten.DrawImageOptions) {
	screen.DrawImage(i.images[rand.Int31n(int32(len(i.images)))], options)
}

func loadImage(path string) *ebiten.Image {
	file, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(fmt.Sprintf("Failed to load image : %v", err))
	}
	return file
}
