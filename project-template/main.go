package main

import (
	"GoWasTest/scenes"

	"github.com/rocco-gossmann/GoWas"
)

func main() {

	e := GoWas.Init(GoWas.EngineSetup{
		WindowWidth:  160,
		WindowHeight: 144,
	})

	e.Run(&scenes.Debug)

}
