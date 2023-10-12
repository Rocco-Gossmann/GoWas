//go:build js && wasm

package GoWas

import "github.com/rocco-gossmann/GoWas/core"

var engineActive *core.Engine

type EngineSetup core.EngineSetup

func Init(setup EngineSetup) *core.Engine {

	if engineActive != nil {
		panic("Engine.Init was called multiple times. Only one time allowed")
	}

	es := core.EngineSetup(setup)
	engine := core.Engine{}
	engine.Init(&es)
	engineActive = &engine

	engine.Run = func(scene any) {
		engine.SwitchScene(scene)
		engine.Canvas().Run()
	}

	return &engine
}
