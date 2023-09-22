package scenes

import (
	"GoWasProject/bmps"
	"fmt"

	"github.com/rocco-gossmann/GoWas"
	"github.com/rocco-gossmann/GoWas/canvas"
)

type debugScene struct {
	cursorBMP canvas.Bitmap
}

//func (s debugScene) Tick(e *GoWas.Engine, dt float64) bool {
//	fmt.Println("Tick")
//	return true
//}

func (s debugScene) Load(e *GoWas.Engine) {
	e.Canvas().FillColorA(0x00333333, 0xff)
	fmt.Println("Debug-Scene loaded")
}

func (s debugScene) Draw(e *GoWas.Engine, ca *GoWas.EngineCanvas) {
	ca.FillColorA(0x00333333, 0x10)
	mouse := ca.Mouse
	if mouse.X > 0 || mouse.Y > 0 {
		ca.BlitBitmap(&(s.cursorBMP), int32(mouse.X), int32(mouse.Y))
	}
}

var Debug = debugScene{
	cursorBMP: bmps.CursorBMP,
}
