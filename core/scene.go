package core

type Drawable interface {
	Draw(e *EngineState, ec *Canvas)
}
type Tickable interface {
	Tick(e *EngineState) bool
}
type Loadable interface {
	Load(e *EngineState, ec *Canvas)
}
type Unloadable interface {
	Unload(e *EngineState) *struct{}
}

type SceneTickFunction func(e *EngineState) bool
type SceneDrawFunction func(e *EngineState, pixelCount uint32, width, height, uint16, pixels *[]uint32)
type SceneLoadFunction func(e *EngineState, ec *Canvas)
type SceneUnloadFunction func(e *EngineState) *struct{}

type DefaultScene struct{}

func (s DefaultScene) Tick(e *EngineState) bool        { return true }
func (s DefaultScene) Draw(e *EngineState, ec *Canvas) {}
func (s DefaultScene) Unload(e *EngineState) *struct{} { return nil }
