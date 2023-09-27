package main

import (
	"GoWasProject/scenes"

	"github.com/rocco-gossmann/GoWas"
	"github.com/rocco-gossmann/GoWas/core"
)

func main() {

	e := GoWas.Init(core.EngineSetup{
		WindowWidth:  160,
		WindowHeight: 144,
	})

	e.Run(&scenes.Debug)

}
