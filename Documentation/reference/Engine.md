# Engine

This is the entry point to your project. It provides 2 Methods to get started.

## Init

```go
Init(setup core.EngineSetup) *core.Engine
```

Sets the base parameters and environment for to run the Entire thing.

#### Parameters

| parameter | type                               | description                            |
| --------- | ---------------------------------- | -------------------------------------- |
| `setup`   | [`core.EngineSetup`](#enginesetup) | Base-Parameters for running the Engine |

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

| parameter | type                       | description                                |
| --------- | -------------------------- | ------------------------------------------ |
| `scene`   | [`core.Scene`](./Scene.md) | The very first Scene to run in the Project |

# EngineSetup

a struct passed to `core.Engine.Init` to define the basic parameters.

```go
type core.EngineSetup struct {
	WindowHeight, WindowWidth uint16 //<- The dimensions of the canvas to draw on
}
```
