package main

import (
	"GoWasProject/scenes"

	"github.com/rocco-gossmann/GoWas"
)

func main() {

	e := GoWas.Init(GoWas.EngineSetup{
		WindowWidth:  128, // 160
		WindowHeight: 8,   // 144

		AutoClearPixels: true,
		//		AutoClearColor:  0x00333333,

		TileMapWidth:  22,
		TileMapHeight: 22,
	})

	e.Run(&scenes.Debug)
}
