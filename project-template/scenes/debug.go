package scenes

import (
	"GoWasProject/bmps"
	"fmt"

	"github.com/rocco-gossmann/GoWas/core"
	"github.com/rocco-gossmann/GoWas/gfx"
	"github.com/rocco-gossmann/GoWas/types"
)

type debugScene struct {
	tma float64

	cursorBMP *core.Bitmap
	text      *gfx.TextDisplay

	bgSet    *gfx.TileSet
	bgMap    *gfx.TileMap
	bgScroll types.Point

	fpsTime float64
	fpsCnt  int
	fps     int
}

func (s *debugScene) Load(e *core.Engine) {
	fmt.Println("Debug-Scene loaded")

	// Setting the Initial background to light #333333
	//-------------------------------------------------------------------------
	e.Canvas().FillColorA(0x00333333, 0xff, core.CANV_CL_ALL)

	// Setting up a Text-Display
	//-------------------------------------------------------------------------
	s.text = gfx.InitTextDisplay(e) // Initialize a Text-Display (You can have as many as you want)

	s.text. //<- Starting the Text change on a Display
		SetCursor(0, 0).Echo("@ Test {}()<|>"). //<- Settting a Cursor position and Printing Text, starting from that location
		SetCursor(-2, 8).Echo("->").            //<- negative coordinates mean "From the Bottom" and/or "From the Right"

		SetCursor(0, 7).Echo("Hey you!"). //<- Positive coordinate = "From the Top" and/or "From the Left"
		Echo("\nYou are finally\nawake.") // <- If you don't specifiy a location, the text
		//									     continues where the last character was printed
		//									     You can use \n to force a line break and carriage return from within the text

	// Preparing a part on the bottom line for showing a constantly changing value
	s.text.
		SetWrap(false). //<- Disable automatic wrapping, as the the printed value would wrap arround to the first line
		//					 otherwise
		SetCursor(0, -2).Echo("FPS:\nTimer:")

	// Text stays persistent per Text-Display you don't need to reset it each frame

	// Setup Background Map
	//-------------------------------------------------------------------------
	// > Init Background TileSet
	s.bgSet = &gfx.TileSet{}
	s.bgSet.InitFromMapSheet(bmps.DebugTiles, 8, 8)

	// > Init Background Map
	s.bgMap = &gfx.TileMap{}
	s.bgMap.Init(s.bgSet, 22, 22).SetMap([]byte{
		12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7,
		7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12,
		12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7,
		7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12,
		12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7,
		7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12,
		12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7,
		7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12,
		12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7,
		7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12,
		12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7,
		7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12,
		12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7,
		7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12,
		12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7,
		7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12,
		12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7,
		7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12,
		12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7,
		7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12,
		12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7,
		7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12, 7, 12,
	})

}

func (me *debugScene) Tick(e *core.Engine, dt float64) bool {

	// Update Timer
	me.tma += 24 * dt
	me.text.SetCursor(7, -1).Echo(fmt.Sprint(me.tma)) // <- update the Text display with the current timer value

	// Update FPS
	me.fpsTime += dt
	if me.fpsTime >= 1 {

		me.text.SetCursor(7, -2).Echo("        ").
			SetCursor(7, -2).Echo(fmt.Sprint(me.fpsCnt))

		me.fpsTime = 0
		me.fpsCnt = 0
	}

	// Update Background-Scroll
	me.bgScroll.X = uint16(me.tma)
	me.bgScroll.Y = uint16(me.tma)

	return true
}

func (s *debugScene) Draw(e *core.Engine, ca *core.Canvas) {

	// Update FPS Counter
	s.fpsCnt++

	s.bgMap.ToCanvas(ca, &gfx.ToCanvasOpts{
		Scroll: s.bgScroll,
	})

	ca.FillColorA(0x00000000, 0x80, core.CANV_CL_ALL) // Filling the canvas with a half transparent black
	//													 to darken the backaground a bit

	s.text.ToCanvas(ca) //<- Tell the engine to display the TextDisplay

	// Draw the Mouse Cursor
	//-------------------------------------------------------------------------
	mouse := ca.Mouse
	if mouse.X > 0 || mouse.Y > 0 {
		ca.Blit(&core.BlitSettings{
			Bmp: s.cursorBMP,
			X:   int32(mouse.X),
			Y:   int32(mouse.Y),
		})
	}

}

func (s *debugScene) Unload(e *core.Engine) *struct{} {
	s.text = nil
	s.bgMap = nil
	s.bgSet = nil
	return nil
}

var Debug = debugScene{
	cursorBMP: bmps.CursorBMP,
}
