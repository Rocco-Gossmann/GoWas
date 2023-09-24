package scenes

import (
	"GoWasProject/bmps"
	"fmt"

	"github.com/rocco-gossmann/GoWas"
	"github.com/rocco-gossmann/GoWas/canvas"
)

type debugScene struct {
	cursorBMP canvas.Bitmap
	fontBMP   canvas.Bitmap
}

//func (s debugScene) Tick(e *GoWas.Engine, dt float64) bool {
//	fmt.Println("Tick")
//	return true
//}

func (s debugScene) Load(e *GoWas.Engine) {
	e.Canvas().FillColorA(0x00333333, 0xff, GoWas.CANV_CL_ALL)
	fmt.Println("Debug-Scene loaded")
}

func (s debugScene) Draw(e *GoWas.Engine, ca *GoWas.EngineCanvas) {
	ca.FillColorA(0x00333333, 0x10, GoWas.CANV_CL_ALL)

	ca.BlitBitmap(&(s.fontBMP), 0, 0, 0xff, GoWas.CANV_CL_NONE)

	mouse := ca.Mouse
	if mouse.X > 0 || mouse.Y > 0 {
		ca.BlitBitmap(&(s.cursorBMP), int32(mouse.X), int32(mouse.Y), 0xff, GoWas.CANV_CL_NONE)
	}
}

var Debug = debugScene{
	cursorBMP: bmps.CursorBMP,
	fontBMP:   bmps.AsciiFontBMP,
}
