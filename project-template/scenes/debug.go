package scenes

import (
	"GoWasProject/bmps"
	"GoWasTest/tilesets"
	"fmt"

	"github.com/rocco-gossmann/GoWas/core"
	"github.com/rocco-gossmann/GoWas/gfx"
)

type debugScene struct {
	cursorBMP *core.Bitmap
	tsFont    *gfx.TileSet
	mapText   *gfx.TileMap
}

//func (s debugScene) Tick(e *GoWas.Engine, dt float64) bool {
//	fmt.Println("Tick")
//	return true
//}

func (s debugScene) Load(e *core.Engine) {
	fmt.Println("Debug-Scene loaded")
	e.Canvas().FillColorA(0x00333333, 0xff, core.CANV_CL_ALL)

	var text = gfx.TileMap{}
	text.Init(s.tsFont, 20, 5).SetMap([]byte{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	})

	s.mapText = &text
}

func (s debugScene) Unload(e *core.Engine) *struct{} {
	s.mapText = nil
	return nil
}

func (s debugScene) Draw(e *core.Engine, ca *core.Canvas) {
	ca.FillColorA(0x00333333, 0x10, core.CANV_CL_ALL)

	s.tsFont.BlitTo(ca, 65-32, nil) // Print A

	mouse := ca.Mouse
	if mouse.X > 0 || mouse.Y > 0 {
		ca.Blit(&core.BlitSettings{
			Bmp: s.cursorBMP,
			X:   int32(mouse.X),
			Y:   int32(mouse.Y),
		})
	}
}

var Debug = debugScene{
	cursorBMP: &bmps.CursorBMP,
	tsFont:    &tilesets.TsFont,
}
