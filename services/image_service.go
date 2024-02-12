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

	imageOpacity100Cache := make(map[assetKey]*ebiten.Image)
	imageOpacity80Cache := make(map[assetKey]*ebiten.Image)
	imageOpacity60Cache := make(map[assetKey]*ebiten.Image)
	imageOpacity40Cache := make(map[assetKey]*ebiten.Image)
	imageOpacity20Cache := make(map[assetKey]*ebiten.Image)

	for _, l := range letters {
		opacity100 := toAssetKey(l, static.LetterAssetFolder)
		opacity80 := toAssetKey(l, static.LetterAssetOpacity80Folder)
		opacity60 := toAssetKey(l, static.LetterAssetOpacity60Folder)
		opacity40 := toAssetKey(l, static.LetterAssetOpacity40Folder)
		opacity20 := toAssetKey(l, static.LetterAssetOpacity20Folder)
		imageOpacity100Cache[opacity100] = loadImage(opacity100)
		imageOpacity80Cache[opacity80] = loadImage(opacity80)
		imageOpacity60Cache[opacity60] = loadImage(opacity60)
		imageOpacity40Cache[opacity40] = loadImage(opacity40)
		imageOpacity20Cache[opacity20] = loadImage(opacity20)
	}

	return &imageService{imageOpacity100Cache, imageOpacity80Cache, imageOpacity60Cache, imageOpacity40Cache, imageOpacity20Cache}
}

type ImageService interface {
	PickRandom(opacity Opacity) (string, *ebiten.Image)
	FindByName(name string, opacity Opacity) *ebiten.Image
	DrawRandom(screen *ebiten.Image, options *ebiten.DrawImageOptions, opacity Opacity) (string, *ebiten.Image)
	Draw(screen *ebiten.Image, image *ebiten.Image, options *ebiten.DrawImageOptions)
	DrawWithOpacity(screen *ebiten.Image, imageName string, options *ebiten.DrawImageOptions, opacity Opacity) (string, *ebiten.Image)
}

type imageService struct {
	imageOpacity100Cache map[assetKey]*ebiten.Image
	imageOpacity80Cache  map[assetKey]*ebiten.Image
	imageOpacity60Cache  map[assetKey]*ebiten.Image
	imageOpacity40Cache  map[assetKey]*ebiten.Image
	imageOpacity20Cache  map[assetKey]*ebiten.Image
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

func (i *imageService) PickRandom(opacity Opacity) (string, *ebiten.Image) {
	index := rand.Int31n(int32(len(i.imageOpacity20Cache)))
	var idx int32 = 0

	//TODO hate this, find better way to index with opacity
	switch opacity {
	case Opacity100:
		for k := range i.imageOpacity100Cache {
			if idx == index {
				return k.key, i.imageOpacity100Cache[k]
			}
			idx++
		}
	case Opacity80:
		for k := range i.imageOpacity80Cache {
			if idx == index {
				return k.key, i.imageOpacity80Cache[k]
			}
			idx++
		}
	case Opacity60:
		for k := range i.imageOpacity60Cache {
			if idx == index {
				return k.key, i.imageOpacity60Cache[k]
			}
			idx++
		}
	case Opacity40:
		for k := range i.imageOpacity40Cache {
			if idx == index {
				return k.key, i.imageOpacity40Cache[k]
			}
			idx++
		}
	case Opacity20:
		for k := range i.imageOpacity20Cache {
			if idx == index {
				return k.key, i.imageOpacity20Cache[k]
			}
			idx++
		}
	}

	panic("Failed to pick random image")
}

func (i *imageService) DrawRandom(screen *ebiten.Image, options *ebiten.DrawImageOptions, opacity Opacity) (string, *ebiten.Image) {
	key, image := i.PickRandom(opacity)
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
		return i.imageOpacity100Cache[assetKey{name, static.LetterAssetFolder}]
	case Opacity80:
		return i.imageOpacity80Cache[assetKey{name, static.LetterAssetOpacity80Folder}]
	case Opacity60:
		return i.imageOpacity60Cache[assetKey{name, static.LetterAssetOpacity60Folder}]
	case Opacity40:
		return i.imageOpacity40Cache[assetKey{name, static.LetterAssetOpacity40Folder}]
	case Opacity20:
		return i.imageOpacity20Cache[assetKey{name, static.LetterAssetOpacity20Folder}]
	default:
		panic(fmt.Sprintf("Failed to get image by name:opacity %v:%v", name, opacity))
	}
}

func (i *imageService) DrawWithOpacity(screen *ebiten.Image, imageName string, options *ebiten.DrawImageOptions, opacity Opacity) (string, *ebiten.Image) {
	//TODO hate this, find better way to index with opacity
	switch opacity {
	case Opacity100:
		aKey := toAssetKey(imageName, static.LetterAssetFolder)
		image := i.imageOpacity100Cache[aKey]
		screen.DrawImage(image, options)
		return aKey.key, image
	case Opacity80:
		aKey := toAssetKey(imageName, static.LetterAssetOpacity80Folder)
		image := i.imageOpacity80Cache[aKey]
		screen.DrawImage(image, options)
		return aKey.key, image
	case Opacity60:
		aKey := toAssetKey(imageName, static.LetterAssetOpacity60Folder)
		image := i.imageOpacity60Cache[aKey]
		screen.DrawImage(image, options)
		return aKey.key, image
	case Opacity40:
		aKey := toAssetKey(imageName, static.LetterAssetOpacity40Folder)
		image := i.imageOpacity40Cache[aKey]
		screen.DrawImage(image, options)
		return aKey.key, image
	case Opacity20:
		aKey := toAssetKey(imageName, static.LetterAssetOpacity20Folder)
		image := i.imageOpacity20Cache[aKey]
		screen.DrawImage(image, options)
		return aKey.key, image
	default:
		panic("Failed to pick opacity file")
	}
}

func loadImage(key assetKey) *ebiten.Image {
	file, _, err := ebitenutil.NewImageFromFile(key.ToPath())
	if err != nil {
		panic(fmt.Sprintf("Failed to load image : %v", err))
	}
	return file
}
