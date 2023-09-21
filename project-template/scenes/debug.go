package scenes

import (
	"github.com/rocco-gossmann/GoWas"
	"github.com/rocco-gossmann/GoWas/canvas"

	"GoWasProject/bmps"
)

type debugScene struct {
	player    GoWas.Point
	playerBmp canvas.Bitmap
}

func (s debugScene) Tick(e *GoWas.Engine, dt float64) bool {
	//fmt.Println("Tick")
	return true
}

func (s debugScene) Draw(e *GoWas.Engine, ca *GoWas.EngineCanvas) {
	ca.FillColorA(0x00333333, 0xff)
	mouse := ca.Mouse
	ca.BlitBitmap(&(s.playerBmp), int32(mouse.X), int32(mouse.Y))
}

var Debug = debugScene{
	playerBmp: bmps.PlayerBMP,
}
