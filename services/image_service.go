package services

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math/rand"
	"matrix-screen-go/static"
)

type Opacity int

const (
	Opacity100 Opacity = 100
	Opacity80  Opacity = 80
	Opacity60  Opacity = 60
	Opacity40  Opacity = 40
	Opacity20  Opacity = 20
	Opacity0   Opacity = 0
)

func LowerOpacity(opacity Opacity) Opacity {
	return opacity - 20
}

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

	cache := make(map[assetKey]*ebiten.Image)
	for _, l := range letters {
		opacity100 := toAssetKey(l, static.LetterAssetFolder)
		opacity80 := toAssetKey(l, static.LetterAssetOpacity80Folder)
		opacity60 := toAssetKey(l, static.LetterAssetOpacity60Folder)
		opacity40 := toAssetKey(l, static.LetterAssetOpacity40Folder)
		opacity20 := toAssetKey(l, static.LetterAssetOpacity20Folder)
		cache[opacity100] = loadImage(opacity100)
		cache[opacity80] = loadImage(opacity80)
		cache[opacity60] = loadImage(opacity60)
		cache[opacity40] = loadImage(opacity40)
		cache[opacity20] = loadImage(opacity20)
	}

	return &imageService{imageCache: cache}
}

type ImageService interface {
	PickRandom() (string, *ebiten.Image)
	FindByName(name string, opacity Opacity) *ebiten.Image
	DrawRandom(screen *ebiten.Image, options *ebiten.DrawImageOptions) (string, *ebiten.Image)
	Draw(screen *ebiten.Image, image *ebiten.Image, options *ebiten.DrawImageOptions)
	DrawWithOpacity(screen *ebiten.Image, imageName string, options *ebiten.DrawImageOptions, opacity Opacity) (string, *ebiten.Image)
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

// TODO unify the folder building logic
func (i *imageService) FindByName(name string, opacity Opacity) *ebiten.Image {
	switch opacity {
	case Opacity100:
		return i.imageCache[assetKey{name, static.LetterAssetFolder}]
	case Opacity80:
		return i.imageCache[assetKey{name, static.LetterAssetOpacity80Folder}]
	case Opacity60:
		return i.imageCache[assetKey{name, static.LetterAssetOpacity60Folder}]
	case Opacity40:
		return i.imageCache[assetKey{name, static.LetterAssetOpacity40Folder}]
	case Opacity20:
		return i.imageCache[assetKey{name, static.LetterAssetOpacity20Folder}]
	default:
		panic(fmt.Sprintf("Failed to get image by name:opacity %v:%v", name, opacity))
	}
}

func (i *imageService) DrawWithOpacity(screen *ebiten.Image, imageName string, options *ebiten.DrawImageOptions, opacity Opacity) (string, *ebiten.Image) {
	var aKey assetKey
	switch opacity {
	case Opacity100:
		aKey = toAssetKey(imageName, static.LetterAssetFolder)
	case Opacity80:
		aKey = toAssetKey(imageName, static.LetterAssetOpacity80Folder)
	case Opacity60:
		aKey = toAssetKey(imageName, static.LetterAssetOpacity60Folder)
	case Opacity40:
		aKey = toAssetKey(imageName, static.LetterAssetOpacity40Folder)
	case Opacity20:
		aKey = toAssetKey(imageName, static.LetterAssetOpacity20Folder)
	default:
		panic("Failed to pick opacity file")
	}

	image := i.imageCache[aKey]
	screen.DrawImage(image, options)
	return aKey.key, image
}

func loadImage(key assetKey) *ebiten.Image {
	file, _, err := ebitenutil.NewImageFromFile(key.ToPath())
	if err != nil {
		panic(fmt.Sprintf("Failed to load image : %v", err))
	}
	return file
}
