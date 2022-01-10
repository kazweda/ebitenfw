package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 300
	screenHeight = 300

	barAccelConstant = float64(0.00000015)
	barWidth         = 50
	barHeight        = 10
)

type Game struct {
	pressedKeys []ebiten.Key
}

var (
	barPosX        = float64(screenWidth/2 - barWidth/2)
	barPosY        = float64(screenHeight/2 - barHeight/2)
	barAccelX      = float64(0)
	prevUpdateTime = time.Now()
	timeDelta      = float64(0)
)

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	timeDelta = float64(time.Since(prevUpdateTime))
	prevUpdateTime = time.Now()

	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])

	acc := barAccelConstant

	for _, key := range g.pressedKeys {
		switch key.String() {
		case "ArrowRight":
			barAccelX = acc
		case "ArrowLeft":
			barAccelX = -acc
		}
	}

	barPosX += barAccelX * timeDelta

	if barPosX < 0 {
		barPosX = 0
	}

	if barPosX > screenWidth-50 {
		barPosX = screenWidth - 50
	}
	return nil
}

func (g *Game) drawBar(screen *ebiten.Image) {

	ebitenutil.DrawRect(
		screen,
		barPosX,
		barPosY,
		barWidth,
		barHeight,
		color.RGBA{0x80, 0xa0, 0xc0, 0xff},
	)

}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawBar(screen)

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf(
			"FPS: %0.2f x: %0.2f y: %0.2f ",
			ebiten.CurrentFPS(),
			barPosX,
			barPosY,
		),
	)
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ebiten Test")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
