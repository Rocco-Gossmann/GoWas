//go:build js && wasm

package GoWas

type EngineSetup struct {
	WindowHeight, WindowWidth uint16
}

type Engine struct {
	canvas       *EngineCanvas
	currentScene any

	Tick   *Tickable
	Draw   *Drawable
	Unload *Unloadable

	Run func(scene any)

	Canvas func() *EngineCanvas
}

var engineActive *Engine = nil

func Init(setup EngineSetup) *Engine {

	if engineActive != nil {
		panic("Engine.Init was called multiple times. Only one time allowed")
	}

	engine := Engine{}
	engine.canvas = CreateCanvas(&engine, setup.WindowWidth, setup.WindowHeight)
	engine.Canvas = func() *EngineCanvas { return engine.canvas }

	engineActive = &engine

	engine.Run = func(scene any) {
		engine.switchScene(scene)
		engine.canvas.Run()
	}

	return &engine
}

func (e *Engine) switchScene(scene any) {
	if scene != nil {
		hasInterface := false

		if l, ok := interface{}(scene).(Loadable); ok {
			l.Load(e)
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

type DefaultScene struct{}

func (s DefaultScene) Tick(e *Engine, dt float64) bool  { return true }
func (s DefaultScene) Draw(e *Engine, ec *EngineCanvas) {}
func (s DefaultScene) Unload(e *Engine) *struct{}       { return nil }

var defaultScene = DefaultScene{}
