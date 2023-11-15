package main

import (
	"GoWasProject/scenes"

	"github.com/rocco-gossmann/GoWas"
)

func main() {

	e := GoWas.Init(GoWas.EngineSetup{
		WindowWidth:  160, // 160
		WindowHeight: 144, // 144

		AutoClearPixels: false,
		//		AutoClearColor:  0x00333333,

		TileMapWidth:  22,
		TileMapHeight: 22,
	})

	e.Run(&scenes.Debug)
}
