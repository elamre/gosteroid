package main

import (
	"fmt"
	"github.com/atolVerderben/tentsuyu"
	"github.com/elamre/Gosteroid/common"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
	"sync"
)

const (
	WIDTH  = 640
	HEIGHT = 480
)

//GameDrawHelperFunction is meant to draw something on the passed ebiten.Image
func TopDrawer(screen *ebiten.Image) error {
	backGround.Update()
	backGround.Draw(screen)
	_ = ebitenutil.DebugPrint(screen, fmt.Sprintf("Hello world: %f", ebiten.CurrentFPS()))
	return nil
}

var imageManager *tentsuyu.ImageManager
var initImageManager sync.Once

var audioManager *tentsuyu.AudioPlayer
var initAudioManager sync.Once

var backGround StarBackground

func GetImageManager() *tentsuyu.ImageManager {
	initImageManager.Do(func() {
		imageManager = tentsuyu.NewImageManager()
	})
	return imageManager
}

func GetAudioManager() *tentsuyu.AudioPlayer {
	initAudioManager.Do(func() {
		p, _ := tentsuyu.NewAudioPlayer()
		audioManager = p
	})
	return audioManager
}

func main() {
	common.MathsTest()
	p, _ := ebiten.NewImage(1, 1, ebiten.FilterNearest)
	_ = p.Fill(color.RGBA{R: 255, G: 255, B: 255, A: 255})

	tentsuyu.Pixel = p
	backGround = NewStarBackground(150, WIDTH, HEIGHT)
	gg, _ := tentsuyu.NewGame(WIDTH, HEIGHT)

	gg.LoadImages(GetImageManager)
	gg.LoadAudio(GetAudioManager)
	gg.ImageManager = GetImageManager()
	gg.AudioPlayer = GetAudioManager()

	gg.SetGameDrawLoop(TopDrawer)
	NewStateLogic(gg, common.STATE_MENU)

	if err := ebiten.RunGame(gg); err != nil {
		log.Fatal(err)
	}
}
