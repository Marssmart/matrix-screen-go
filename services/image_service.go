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
	PickRandom() *ebiten.Image
	DrawRandom(screen *ebiten.Image, options *ebiten.DrawImageOptions) *ebiten.Image
	Draw(screen *ebiten.Image, image *ebiten.Image, options *ebiten.DrawImageOptions)
}

type imageService struct {
	images []*ebiten.Image
}

func (i *imageService) PickRandom() *ebiten.Image {
	return i.images[rand.Int31n(int32(len(i.images)))]
}

func (i *imageService) DrawRandom(screen *ebiten.Image, options *ebiten.DrawImageOptions) *ebiten.Image {
	image := i.PickRandom()
	screen.DrawImage(image, options)
	return image
}

func (i *imageService) Draw(screen *ebiten.Image, image *ebiten.Image, options *ebiten.DrawImageOptions) {
	screen.DrawImage(image, options)
}

func loadImage(path string) *ebiten.Image {
	file, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(fmt.Sprintf("Failed to load image : %v", err))
	}
	return file
}
