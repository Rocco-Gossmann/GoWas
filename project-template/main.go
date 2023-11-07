package main

import (
	"GoWasProject/scenes"

	"github.com/rocco-gossmann/GoWas"
)

func main() {

	e := GoWas.Init(GoWas.EngineSetup{
		WindowWidth:  160,
		WindowHeight: 144,

		AutoClearPixels: true,
		AutoClearColor:  0x00333333,
	})

	e.Run(&scenes.Debug)
}
