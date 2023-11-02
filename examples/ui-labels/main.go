package main

import (
	"fmt"

	"github.com/rocco-gossmann/GoWas"
	"github.com/rocco-gossmann/GoWas/core"
	"github.com/rocco-gossmann/GoWas/ui"
)

type ExampleScene struct {
	totalTime float64

	label *ui.Label
}

// ==============================================================================
// Setup
// ==============================================================================
func (me *ExampleScene) Load(state *core.EngineState, canvas *core.Canvas) {

	state.EnableTextLayer() // 1.) Tell the Engine to enable Text / UI on screen
	//                             Otherwise nothing from the ui package will be displayed

	state.Text. // 2.) Add some static Text, that will not be touched
			SetCursor(1, 1).
			Echo("Time:")

	me.label = ui.CreateLabel(state.Text, 7, 1, 12) // 3.) Reseerve the porition of the screen,
	//                                                    That holds the changing time text
	//                                                    (12 chracters at location x=7 and y=1)
}

// ==============================================================================
// Main-Loop
// ==============================================================================
func (me *ExampleScene) Tick(state *core.EngineState) bool {

	me.totalTime += state.DeltaTime // 1.) Update total time by means of DeltaTime (Time passed since last Tick)

	me.label.Text(fmt.Sprint(me.totalTime)) // 2.) Update the Label to show the time passed on next Draw

	return true
}

// ==============================================================================
// Go-Main-Function
// ==============================================================================
func main() {

	scene := ExampleScene{}

	GoWas.Init(GoWas.EngineSetup{
		WindowWidth:     160,
		WindowHeight:    144,
		AutoClearPixels: true,
	}).Run(&scene)
}
