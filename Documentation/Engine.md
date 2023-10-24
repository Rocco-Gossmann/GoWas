# The Engine

The `core.Engine` is the entrypoint to every application using it.

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
For further information see:
- [Engine - Lifecycle](./EngineLifecycle.md)
- [Engine - Reference](./reference/Engine.md)
- [Engine - State](./reference/EngineState.md)




