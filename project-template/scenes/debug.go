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

	CursorEntity *core.BitmapEntity
	bg2          *core.BitmapEntity

	text *gfx.TextDisplay

	bgSet              *gfx.TileSet
	mouseButtonDisplay *gfx.TileSet

	bgMap    *gfx.TileMap
	bgScroll types.Point

	fpsTime float64
	fpsCnt  int
	fps     int
}

func (s *debugScene) Load(e *core.EngineState, ca *core.Canvas) {
	fmt.Println("Debug-Scene loaded")

	// Setting the Initial background to light #333333
	//-------------------------------------------------------------------------
	ca.FillColorA(0x00333333, 0xff, core.CANV_CL_ALL)

	// Load in Mouse Button Display-Tileset
	//-------------------------------------------------------------------------
	mbts := gfx.TileSet{}
	mbts.InitFromMapSheet(bmps.BMPmouse, 32, 32)
	s.mouseButtonDisplay = &mbts

	// Setting up a Text-Display
	//-------------------------------------------------------------------------
	s.text = gfx.InitTextDisplay(ca) // Initialize a Text-Display (You can have as many as you want)

	s.text. //<- Starting the Text change on a Display
		SetCursor(0, 0).Echo("@ Test {}()<|>"). //<- Settting a Cursor position and Printing Text, starting from that location
		SetCursor(-2, 0).Echo("->").            //<- negative coordinates mean "From the Bottom" and/or "From the Right"

		SetCursor(0, 6).Echo("Hey you!").                                          //<- Positive coordinate = "From the Top" and/or "From the Left"
		Echo(" Move the\nmouse over this\nScreen and press\none of it's Buttons.") // <- If you don't specifiy a location, the text
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
	s.bgSet.InitFromMapSheet(bmps.BMPdebugtiles, 8, 8)

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

	s.bg2.MoveBy(-4, -4).Alpha(0x80, false)
}

func (me *debugScene) Tick(e *core.EngineState) bool {

	// Update Timer
	me.tma += 24 * e.DeltaTime
	me.text.SetCursor(7, -1).Echo(fmt.Sprint(me.tma)) // <- update the Text display with the current timer value

	// Update FPS
	me.fpsTime += e.DeltaTime
	if me.fpsTime >= 1 {

		me.text.
			SetCursor(7, -2).
			Clear(8). //<-- overwrite the next 8 characters from the cursor with spaces
			Echo(fmt.Sprint(me.fpsCnt))

		me.fpsTime = 0
		me.fpsCnt = 0
	}

	// Update Background-Scroll
	me.bgScroll.X = uint16(me.tma)
	me.bgScroll.Y = uint16(me.tma)

	return true
}

func (s *debugScene) Draw(e *core.EngineState, ca *core.Canvas) {
	// Update FPS Counter
	s.fpsCnt++

	s.bgMap.ToCanvas(ca, &gfx.ToCanvasOpts{
		Scroll: s.bgScroll,
	})

	s.bg2.ToCanvas(ca)

	ca.FillColorA(0x00000000, 0xc8, core.CANV_CL_ALL) // Filling the canvas with a half transparent black
	//													 to darken the backaground a bit

	s.text.
		SetCursor(5, 13).
		Clear(9).
		Echo(fmt.Sprintf("%v, %v", e.Mouse.X, e.Mouse.Y)).
		ToCanvas(ca) //<- Tell the engine to display the TextDisplay

	// Draw the Mouse Button Display
	//-------------------------------------------------------------------------
	s.mouseButtonDisplay.BlitTo(ca, int(e.Mouse.PressedOrHeld), &gfx.TilesetBlitOptions{
		X: 0, Y: 88,
	})

	// Draw the Mouse Cursor
	//-------------------------------------------------------------------------
	s.CursorEntity.MoveTo(int32(e.Mouse.X), int32(e.Mouse.Y))
	if e.Mouse.X > 0 || e.Mouse.Y > 0 {
		s.CursorEntity.ToCanvas(ca)
	}

}

func (s *debugScene) Unload(e *core.EngineState) *struct{} {
	s.text = nil
	s.bgMap = nil
	s.bgSet = nil
	s.mouseButtonDisplay = nil
	return nil
}

var Debug = debugScene{
	CursorEntity: bmps.BMPcursor.MakeEntity(),
	bg2:          bmps.BMPoversized.MakeEntity(),
}
