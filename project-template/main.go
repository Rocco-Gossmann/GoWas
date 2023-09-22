package main

import (
	"GoWasTest/scenes"

	"github.com/rocco-gossmann/GoWas"
)

func main() {

	e := GoWas.Init(GoWas.EngineSetup{
		WindowWidth:  256,
		WindowHeight: 192,
	})

	e.Run(&scenes.Debug)

}
