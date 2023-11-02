# The Engine

The Entry Point for the Engine is the `GoWas.Init` - function.
It returns a [`core.Engine`](./reference/Engine.md) that  can then be run

This is the barebones skeleton you should use in your `main.go` file.

```go
package main

import (
    "MyProject/scenes"  //<- This must be defined by yourself

    "github.com/rocco-gossmann/GoWas"
)

func main() {
    e := GoWas.Init(
        GoWas.EngineSetup{
            WindowWidth:  160, // <- Width of the canvas in pixels
            WindowHeight: 144, // <- Height of the canvas in pixels
        }
    )

    // Starting the Engine by telling it, what the first Scene is
    e.Run(&scenes.Debug)
}
```

A shorter form would be:

```go
package main

import (
    "MyProject/scenes"  //<- This must be defined by yourself
    "github.com/rocco-gossmann/GoWas"
)

func main() {
    GoWas.Init( 
        GoWas.EngineSetup{
            WindowWidth:  160, // <- Width of the canvas in pixels
            WindowHeight: 144, // <- Height of the canvas in pixels
        }
    ).Run(&scenes.Debug)
}
```




For further information see:
- [Engine - Lifecycle](./EngineLifecycle.md)
- [Engine - Reference](./reference/Engine.md)
- [Engine - State](./reference/EngineState.md)




