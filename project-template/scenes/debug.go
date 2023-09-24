package scenes

import (
	"GoWasProject/bmps"
	"fmt"

	"github.com/rocco-gossmann/GoWas"
	"github.com/rocco-gossmann/GoWas/canvas"
)

type debugScene struct {
	cursorBMP *canvas.Bitmap
	fontBMP   *canvas.Bitmap
}

//func (s debugScene) Tick(e *GoWas.Engine, dt float64) bool {
//	fmt.Println("Tick")
//	return true
//}

func (s debugScene) Load(e *GoWas.Engine) {
	fmt.Println("Debug-Scene loaded")

	e.Canvas().FillColorA(0x00333333, 0xff, GoWas.CANV_CL_ALL)

}

func (s debugScene) Draw(e *GoWas.Engine, ca *GoWas.EngineCanvas) {
	ca.FillColorA(0x00333333, 0x10, GoWas.CANV_CL_ALL)

	ca.Blit(GoWas.BlitSettings{Bmp: s.fontBMP})
	//	ca.BlitBitmap(s.fontBMP, 0, 0, 0xff, GoWas.CANV_CL_NONE)

	mouse := ca.Mouse
	if mouse.X > 0 || mouse.Y > 0 {
		ca.Blit(GoWas.BlitSettings{
			Bmp: s.cursorBMP,
			X:   int32(mouse.X),
			Y:   int32(mouse.Y),
		})
	}
}

var Debug = debugScene{
	cursorBMP: &bmps.CursorBMP,
	fontBMP:   &bmps.AsciiFontBMP,
}
