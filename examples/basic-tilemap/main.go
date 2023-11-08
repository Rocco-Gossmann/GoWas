package main

import (
	"GoWasProject/bmps"

	"github.com/rocco-gossmann/GoWas"
	"github.com/rocco-gossmann/GoWas/core"
)

type ExampleScene struct {
	movetiming float64
}

// ==============================================================================
// Setup
// ==============================================================================
func (me *ExampleScene) Load(state *core.EngineState, canvas *core.Canvas) {

	// 1.) Before we can enable any Tilemap for the first time,
	//     we must define what TileSet is is going to use
	tileSet := core.TileSet{}
	tileSet.InitFromMapSheet(bmps.BMPdebugtiles, 8, 8)

	// 2.) now we can tell the Engine to enable a Tilemap for the first time
	//     Map 1 is the Background-Map, so that is what we are going to use here
	state.EnableMap1Layer(
		// 				   You can also pass `nil` here, but that would only keep
		// 				   the currently active TileSet, since this is the FIRST ENABLE
		//                 we MUST pass a Tileset, otherwise nothing would happen
		&tileSet,
	)

	// 3.) Next we fill the Map with values. (By default each tile is set to 0 first)
	//     0 means, the tile is not rendered
	state.Map1.SetMap([]byte{
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

// ==============================================================================
// Main-Loop
// ==============================================================================
func (me *ExampleScene) Tick(state *core.EngineState) bool {

	me.movetiming += 32 * state.DeltaTime

	state.Map1.MoveTo(
		int32(me.movetiming),
		int32(me.movetiming),
	)

	return true
}

// ==============================================================================
// Go-Main-Function
// ==============================================================================
func main() {

	scene := ExampleScene{}

	GoWas.Init(GoWas.EngineSetup{
		WindowWidth:  160, //<- classic Gameboy resolution
		WindowHeight: 144,

		AutoClearPixels: false, // <- Don't need it, since the
		//                            tilemap will cover the entire screen

		TileMapWidth:  22, //<- Setting Width and Height to 22 will make it easy to create an infinitly
		TileMapHeight: 22, //   Looping background
	}).Run(&scene)
}
