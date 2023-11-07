package scenes

import (
	"GoWasProject/bmps"
	"fmt"

	"github.com/rocco-gossmann/GoWas/core"
	"github.com/rocco-gossmann/GoWas/io"
	"github.com/rocco-gossmann/GoWas/ui"
)

const validMouseButtons = io.MOUSE_BTN1 |
	io.MOUSE_BTN2 |
	io.MOUSE_BTN3

type debugScene struct {
	tma       float64
	totaltime float64
	fpsTime   float64
	fpsCnt    int
	fps       int

	fpsLabel  *ui.Label
	timeLabel *ui.Label

	exampleRessource core.RessourceHandle
}

func (s *debugScene) Load(e *core.EngineState, ca *core.Canvas) {

	s.initTextDisplay(e)
	s.initMaps(e)

	e.SetLayerOrder(
		core.CANV_RL_MAP2,  // Top most layer is drawn above all (Map2 has transparency to tint everything)
		core.CANV_RL_TEXT,  // Then the text
		core.CANV_RL_SCENE, // Scene.Draw below the text
		core.CANV_RL_MAP1,  // Map 1 as the background

		core.CANV_RL_SPRITES, // Sprites don't matter for now, so they are covered by Map 1
	)

	// Load in Mouse Button Display-Tileset
	//-------------------------------------------------------------------------
	//mbts := core.TileSet{}
	//mbts.InitFromMapSheet(bmps.BMPmouse, 32, 32)
	//s.mouseButtonDisplay = &mbts

	// Setting up a Text-Display
	//-------------------------------------------------------------------------
	// s.text = core.InitTextDisplay(ca) // Initialize a Text-Display (You can have as many as you want)

	//	s.text. //<- Starting the Text change on a Display
	//		SetCursor(0, 0).Echo("--- start typing ---"). //<- Settting a Cursor position and Printing Text, starting from that location
	//		SetCursor(-15, 2).Echo("<- last key").        //<- negative coordinates mean "From the Bottom" and/or "From the Right"
	//
	//		SetCursor(0, 6).Echo("Hey you!").                                           //<- Positive coordinate = "From the Top" and/or "From the Left"
	//		Echo(" Move the\nmouse over this\nScreen and press\none of it's Buttons."). // <- If you don't specifiy a location, the text
	//		//																				  continues where the last character was printed
	//		//									     										  You can use \n to force a line break and carriage return from within the text
	//		SetCursor(5, 13).Echo(fmt.Sprintf("Pressed / Held:"))
	//
	// Text stays persistent per Text-Display you don't need to reset it each frame

	// > Init Background Map
	//	s.bgMap = &core.TileMap{}
	//	s.bgMap.

	fmt.Println("Debug-Scene loaded")
}

func (me *debugScene) Tick(e *core.EngineState) bool {

	me.updateTimes(e)

	return true
}

func (s *debugScene) Draw(e *core.EngineState, ca *core.Canvas) {

	s.countFPS()

	//s.bgMap.
	//	MoveTo(s.bgScroll, s.bgScroll).
	//	ToCanvas(ca)

	//	s.text.
	//		SetCursor(5, 14).Clear(8).
	//		SetCursor(5, 11).Clear(9).
	//		Echo(fmt.Sprintf("%v, %v\n\n\n     %v", e.Mouse.X, e.Mouse.Y, e.Mouse.PressedOrHeld)).
	//		ToCanvas(ca) //<- Tell the engine to display the TextDisplay
	//
	//	// Draw the Mouse Button Display
	//	//-------------------------------------------------------------------------
	//	// This is not going to stay as it is. It will be replaced by propper sprites, that function similar to
	//	// How BitmapEntitys, Maps and TextDisplays do
	//	s.mouseButtonDisplay.BlitTo(ca, int(e.Mouse.PressedOrHeld&validMouseButtons), &core.TilesetBlitOptions{
	//		X: 0, Y: 88,
	//	})

	// Draw the Mouse Cursor
	//-------------------------------------------------------------------------
	// s.CursorEntity.MoveTo(int32(e.Mouse.X), int32(e.Mouse.Y))
	// if e.Mouse.X > 0 || e.Mouse.Y > 0 {
	// 	s.CursorEntity.ToCanvas(ca)
	// }

}

func (s *debugScene) Unload(e *core.EngineState) *struct{} {
	e.FreeRessource(s.exampleRessource)
	return nil
}

var Debug = debugScene{
	// CursorEntity: bmps.BMPcursor.MakeEntity(),
}

// Helpers
// ------------------------------------------------------------------------------

// /*Map Test
// ------------------------------------------------------------------------------
func (s *debugScene) initMaps(e *core.EngineState) {
	// Setup Background Map
	//-------------------------------------------------------------------------
	// > Init Background TileSet
	tileSet := core.TileSet{}
	tileSet.InitFromMapSheet(bmps.BMPdebugtiles, 8, 8)

	e.EnableMap1Layer(&tileSet)
	e.Map1.SetAlpha(0x40).SetMap([]byte{
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

	e.EnableMap2Layer(&tileSet)
	e.Map2.SetAlpha(0x70).SetMap([]byte{
		11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5,
		5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11,
		11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5,
		5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11,
		11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5,
		5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11,
		11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5,
		5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11,
		11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5,
		5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11,
		11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5,
		5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11,
		11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5,
		5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11,
		11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5,
		5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11,
		11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5,
		5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11,
		11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5,
		5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11,
		11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5,
		5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11, 5, 11,
	})
}

// /*Text Display Test
// ------------------------------------------------------------------------------
func (s *debugScene) initTextDisplay(e *core.EngineState) {
	// Text-Layer
	//-------------------------------------------------------------------------
	e.EnableTextLayer()
	e.Text.SetCursor(0, -2).Echo("FPS:\nTime:") //<-- Print some static text, that won't change
	// Create A UI-Label that we can change without having to worry about the current TextDisplay-State
	s.fpsLabel = ui.CreateLabel(e.Text, 6, -2, 4)
	s.timeLabel = ui.CreateLabel(e.Text, 6, -1, 10)
}

func (me *debugScene) updateTimes(e *core.EngineState) {
	// Update Timers
	me.fpsTime += e.DeltaTime
	me.totaltime += e.DeltaTime
	me.tma += 24 * e.DeltaTime // <-background scroll timer

	me.timeLabel.Text(fmt.Sprint(me.totaltime))

	// Update FPS
	if me.fpsTime >= 1 {

		me.fpsLabel.Text(fmt.Sprint(me.fpsCnt)) //<-display the text via the created ui.Label

		me.fpsTime = 0
		me.fpsCnt = 0
	}
}

func (me *debugScene) countFPS() {
	me.fpsCnt++
}

//*/
