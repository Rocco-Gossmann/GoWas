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

	// If set, all Pixels and Collisions on screen are reset to 0 each frame
	AutoClearPixels bool

	AutoClearColor uint32 // 0x00 RR GG BB   24bit color that the screen clears to if AutoClearPixels is true

	// [Default: 32] How many tiles in width and height the Maps for Layers Map1 and Map2 will have
	TileMapWidth, TileMapHeight uint32
}

type Engine struct {
	canvas       *Canvas
	currentScene any

	Tick   *Tickable
	Draw   *Drawable
	Unload *Unloadable

	Run func(scene any)

	textDisplay *TextDisplay
	tileSet     TileSet
}

type EngineState struct {
	Mouse     io.MouseState
	Keyboard  io.KeyboardState
	DeltaTime float64
	Text      *TextDisplay
	Map1      TileMap
	Map2      TileMap

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

func (me *EngineState) EnableMap1Layer(tileSet *TileSet) {
	if tileSet != nil {
		me.Map1.SetTileSet(tileSet)
	}

	if me.Map1.HasTileSet() {
		me.canvas.enableLayer(CANV_RL_MAP1)
	}
}
func (me *EngineState) DisableMap1Layer() {
	me.canvas.disableLayer(CANV_RL_MAP1)
}

func (me *EngineState) EnableMap2Layer(tileSet *TileSet) {
	if tileSet != nil {
		me.Map2.SetTileSet(tileSet)
	}

	if me.Map2.HasTileSet() {
		me.canvas.enableLayer(CANV_RL_MAP2)
	}
}
func (me *EngineState) DisableMap2Layer() {
	me.canvas.disableLayer(CANV_RL_MAP2)
}

func (me *EngineState) SetLayerOrder(
	topMost CanvasRenderLayers,
	belowTop CanvasRenderLayers,
	middle CanvasRenderLayers,
	belowMiddle CanvasRenderLayers,
	last CanvasRenderLayers,
) {
	me.canvas.layerOrder[5] = last
	me.canvas.layerOrder[4] = belowMiddle
	me.canvas.layerOrder[3] = middle
	me.canvas.layerOrder[2] = belowTop
	me.canvas.layerOrder[1] = topMost

	//fmt.Println(me.canvas.layerOrder, me.canvas.layerEnable)

	me.canvas.reorderLayers()
}

// ==============================================================================
// Methods
// ==============================================================================
func (e *Engine) Init(s *EngineSetup) {
	//fmt.Println("[core.engine.Init] run")
	if e == nil {
		panic("'engine' can't be nil")
	}
	if s == nil {
		panic("'setup' can't be nil")
	}

	if s.TileMapWidth == 0 {
		s.TileMapWidth = 32
	}
	if s.TileMapHeight == 0 {
		s.TileMapHeight = 32
	}

	engineState.Map1.Init(nil, s.TileMapWidth, s.TileMapHeight)
	engineState.Map2.Init(nil, s.TileMapWidth, s.TileMapHeight)

	e.canvas = CreateCanvas(e, (*s).WindowWidth, (*s).WindowHeight)
	e.textDisplay = InitTextDisplay(e.canvas)
	e.canvas.layers[CANV_RL_TEXT] = e.textDisplay

	engineState.Text = e.textDisplay

	if s.AutoClearPixels {
		e.canvas.enableLayer(0)
		e.canvas.clearlayer.color = s.AutoClearColor
	}
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
			e.canvas.enableLayer(CANV_RL_SCENE)
		} else {
			i, _ := interface{}(defaultScene).(Drawable)
			e.Draw = &i
			e.canvas.disableLayer(CANV_RL_SCENE)
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
