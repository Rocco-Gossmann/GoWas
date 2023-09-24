//go:build js && wasm

package GoWas

import "github.com/rocco-gossmann/GoWas/core"

var engineActive *core.Engine

func Init(setup core.EngineSetup) *core.Engine {

	if engineActive != nil {
		panic("Engine.Init was called multiple times. Only one time allowed")
	}

	engine := core.Engine{}
	engine.Init(&setup)
	engineActive = &engine

	engine.Run = func(scene any) {
		engine.SwitchScene(scene)
		engine.Canvas().Run()
	}

	return &engine
}
