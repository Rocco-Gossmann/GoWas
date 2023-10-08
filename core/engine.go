package core

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
}

type EngineState struct {
	Mouse     MouseState
	DeltaTime float64
}

func (e *Engine) Init(s *EngineSetup) {
	if e == nil {
		panic("'engine' can't be nil")
	}
	if s == nil {
		panic("'setup' can't be nil")
	}

	e.canvas = CreateCanvas(e, (*s).WindowWidth, (*s).WindowHeight)
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
