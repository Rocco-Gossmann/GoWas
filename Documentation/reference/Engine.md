# Engine

This is the entry point to your project. It provides 2 Methods to get started.

## Init

```go
Init(setup GoWas.EngineSetup) *core.Engine
```

Sets the base parameters and environment for to run the Entire thing.

#### Parameters

| parameter | type                               | description                            |
| --------- | ---------------------------------- | -------------------------------------- |
| `setup`   | [`GoWas.EngineSetup`](#enginesetup) | Base-Parameters for running the Engine |

#### Returns

`*core.Engine` - The Pointer to the Instance of the Engine, that was just
created.

## Run

```go
Run(scene *core.Scene)
```

Tells the Engine what Scene to run First.\
The Function will run, until it no longer receives a Scene to Run

(See [Engine-Lifecycle](../Engine.md#enginelifecycle) for further information)

#### Parameters

| parameter | type                         | description                                                              |
| --------- | ---------------------------- | ------------------------------------------------------------------------ |
| `scene`   | \*[`core.Scene`](../Scenes.md) | The Memory address/pointer to the very first Scene to run in the Project |

# EngineSetup

a struct passed to `GoWas.Init` to define the basic parameters for the Engine to Run

```go
type GoWas.EngineSetup struct {
	WindowHeight, WindowWidth uint16 //<- The dimensions of the canvas to draw on

	AutoClearPixels bool //<- [Default: false] If true, the Pixels of the screen are reset 
	//                        to black or a  at the start of each drawing cycle

	AutoClearColor uint32 // <- [Default: 0x00000000] 0x00RRGGBB 24bit color 
	//                          that the screen clears to if AutoClearPixels is true

	TileMapWidth, TileMapHeight uint32 // <- [Default: 32] How many tiles in 
	//                                       width and height the Maps for Layers 
	//                                       Map1 and Map2 will have
}
```
