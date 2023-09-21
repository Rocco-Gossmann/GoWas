package GoWas

type Drawable interface {
	Draw(e *Engine, ec *EngineCanvas)
}
type Tickable interface {
	Tick(e *Engine, dt float64) bool
}
type Loadable interface{ Load(e *Engine) }
type Unloadable interface{ Unload(e *Engine) *struct{} }

type SceneTickFunction func(e *Engine, dt float64) bool
type SceneDrawFunction func(e *Engine, pixelCount uint32, width, height, uint16, pixels *[]uint32)
type SceneLoadFunction func(e *Engine)
type SceneUnloadFunction func(e *Engine) *struct{}
