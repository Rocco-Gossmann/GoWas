package core

import (
	"syscall/js"

	"github.com/rocco-gossmann/GoWas/io"
)

// ==============================================================================
// Types
// ==============================================================================
type EngineSetup struct {
	WindowHeight, WindowWidth uint16
}

type Engine struct {
	canvas       *Canvas
	currentScene any

	Tick   *Tickable
	Draw   *Drawable
	Unload *Unloadable

	Run func(scene any)

	textDisplay *TextDisplay
}

type EngineState struct {
	Mouse     io.MouseState
	Keyboard  io.KeyboardState
	DeltaTime float64
	Text      *TextDisplay

	ressources map[RessourceHandle]Ressource

	canvas *Canvas
	engine *Engine
}

// ==============================================================================
// Ressource Processing Framework
// ==============================================================================

type _ResourceProcessorFunc func(*Ressource, []js.Value)

func (me *EngineState) _ProcessRessource(args []js.Value, handler _ResourceProcessorFunc) interface{} {
	ressourceHandle := RessourceHandle(args[0].Int())
	if res, ok := me.ressources[ressourceHandle]; ok {
		handler(&res, args)
		me.ressources[ressourceHandle] = res
	}

	return nil
}

// ==============================================================================
// Ressource Processing Functions
// ==============================================================================
func (me *EngineState) _ProcessReadRessource(res *Ressource, args []js.Value) {
	res.state = RESSTATE_PROCESSING
	res.jsData = args[1]
	res._Process()
}

func (me *EngineState) _ProcessNotFoundRessource(res *Ressource, _ []js.Value) {
	res.state = RESSTATE_NOTFOUND
}

// ==============================================================================
// Ressource Ressource Functions
// ==============================================================================
func (me *EngineState) RequestRessource(ressourceType RessourceType, fileName string) RessourceHandle {
	ressource := _RequestRessource(ressourceType, fileName)
	me.ressources[ressource.handle] = ressource
	return ressource.handle
}

func (me *EngineState) reseiveRessource(args []js.Value) interface{} {
	return me._ProcessRessource(args, me._ProcessReadRessource)
}

func (me *EngineState) markRessourceNotFound(args []js.Value) interface{} {
	return me._ProcessRessource(args, me._ProcessNotFoundRessource)
}

func (me *EngineState) FreeRessource(handle RessourceHandle) {
	delete(me.ressources, handle)
}

// ==============================================================================
// TextDisplay
// ==============================================================================

func (me *EngineState) EnableTextLayer() {
	me.canvas.enableLayer(CANV_RL_TEXT)
}

func (me *EngineState) DisableTextLayer() {
	me.canvas.disableLayer(CANV_RL_TEXT)
}

// ==============================================================================
// Methods
// ==============================================================================
func (e *Engine) Init(s *EngineSetup) {
	if e == nil {
		panic("'engine' can't be nil")
	}
	if s == nil {
		panic("'setup' can't be nil")
	}

	e.canvas = CreateCanvas(e, (*s).WindowWidth, (*s).WindowHeight)
	e.textDisplay = InitTextDisplay(e.canvas)
}

func (e *Engine) Canvas() *Canvas { return e.canvas }

func (e *Engine) SwitchScene(scene any) {
	if scene != nil {
		hasInterface := false

		if l, ok := interface{}(scene).(Loadable); ok {
			l.Load(&engineState, e.canvas)
		}

		if t, ok := interface{}(scene).(Tickable); ok {
			e.Tick = &t
			hasInterface = true
		} else {
			i, _ := interface{}(defaultScene).(Tickable)
			e.Tick = &i
		}

		if d, ok := interface{}(scene).(Drawable); ok {
			e.Draw = &d
			hasInterface = true
		} else {
			i, _ := interface{}(defaultScene).(Drawable)
			e.Draw = &i
		}

		if u, ok := interface{}(scene).(Unloadable); ok {
			e.Unload = &u
		} else {
			i, _ := interface{}(defaultScene).(Unloadable)
			e.Unload = &i
		}

		if hasInterface {
			e.currentScene = scene

		} else {
			e.currentScene = nil
			panic("Given Scene must implement Tickable and/or Drawable Interface")

		}
	}
}

var defaultScene = DefaultScene{}
