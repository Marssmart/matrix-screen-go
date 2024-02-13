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

	imageCache := make(map[assetKey]*ebiten.Image)

	for _, l := range letters {
		key := toAssetKey(l, static.LetterAssetFolder)
		imageCache[key] = loadImage(key)
	}

	return &imageService{imageCache}
}

type ImageService interface {
	PickRandom() (string, *ebiten.Image)
	FindByName(name string) *ebiten.Image
	DrawRandom(screen *ebiten.Image, options *ebiten.DrawImageOptions) (string, *ebiten.Image)
	Draw(screen *ebiten.Image, image *ebiten.Image, options *ebiten.DrawImageOptions)
}

type imageService struct {
	imageCache map[assetKey]*ebiten.Image
}

type assetKey struct {
	key    string
	folder string
}

func (k *assetKey) ToPath() string {
	return k.folder + k.key
}

func toAssetKey(key string, folder string) assetKey {
	return assetKey{key, folder}
}

func (i *imageService) PickRandom() (string, *ebiten.Image) {
	index := rand.Int31n(int32(len(i.imageCache)))
	var idx int32 = 0
	for k := range i.imageCache {
		if idx == index {
			return k.key, i.imageCache[k]
		}
		idx++
	}

	panic("Failed to pick random image")
}

func (i *imageService) DrawRandom(screen *ebiten.Image, options *ebiten.DrawImageOptions) (string, *ebiten.Image) {
	key, image := i.PickRandom()
	screen.DrawImage(image, options)
	return key, image
}

func (i *imageService) Draw(screen *ebiten.Image, image *ebiten.Image, options *ebiten.DrawImageOptions) {
	screen.DrawImage(image, options)
}

func (i *imageService) FindByName(name string) *ebiten.Image {
	return i.imageCache[assetKey{name, static.LetterAssetFolder}]
}

func loadImage(key assetKey) *ebiten.Image {
	file, _, err := ebitenutil.NewImageFromFile(key.ToPath())
	if err != nil {
		panic(fmt.Sprintf("Failed to load image : %v", err))
	}
	return file
}
