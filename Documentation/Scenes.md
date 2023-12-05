# Scenes

<!-- TOC depthfrom:1 -->

- [Scenes](#scenes)
    - [Lifecycle-Hooks](#lifecycle-hooks)
        - [Tickable](#tickable)
        - [Drawable](#drawable)
        - [Loadable](#loadable)
        - [Unloadable](#unloadable)
    - [Barebones Example](#barebones-example)
    - [Example for Using a Scene](#example-for-using-a-scene)

<!-- /TOC -->

A scene is a `struct{}` you define yourself. It hooks into the
[Engines Lifecycle](./EngineLifecycle.md) to manipulate and change it in various
ways.

```go
type myScene struct {}
```

## Lifecycle-Hooks

In total, there are 4 Interfaces a Scene `struct{}` can implement. By Implementing
at least either the `Tickable` and/or the `Drawable`` Interface, The Scene
becomes a valid `core.Scene`.

### Tickable

If a Scene implements it, its `Tick` method is called every time the Engine hits
a new cycle.

> [!important] 
> If your Scene does not implement the `core.Tickable` interface,
> The Engine will run indefinitly on this Scene.

```go
type Tickable interface {
    Tick(e *EngineState) bool
}

// Example:
// func (s *myScene) Tick(e *core.EngineState) bool { return true }
```

### Drawable

If a Scene implements it, its `Draw` method is called after a Tick returned
`true` (should your scene not implement the `Tickable` interface, this is run
each Cycle)

```go
type Drawable interface {
    Draw(e *EngineState, ec *Canvas)
}

// Example:
// func (s *myScene) Draw(e *core.EngineState, e *core.Canvas) {}
```

### Loadable

If a Scene implements it, its `Load` method is called before the Engine 
starts/continues ticking.

It is only called once and should therefore be used to load/reserve/define 
resources that the. `Tick(...)` and/or `Draw(...)` methods may need.

```go
type Loadable interface {
    Load(e *EngineState, ec *Canvas)
}
// Example:
// func (s *myScene) Load(e *core.EngineState, e *core.Canvas) {}
```

### Unloadable

If a Scene implements it, its `Unload` method is called before the Engine
continues ticking.

Like the `Loadable` interface, this is also only called once but only after a 
scene's `Tick` method. returned `false`.

Therefore it should be used to unload/free/reset the resources created in the `Load` method.

The `Unload` method returns a `struct{}` - pointer, that points to the next
`core.Scene`, that the Engine should load and run next.

If it returns `nil` instead, the Engine will stop running and the program will
end.

```go
type Unloadable interface {
    Unload(e *EngineState) *struct{}
}
// Example:
// func (s *myScene) Unload(e *core.EngineState) *struct{} { return nil }
```

## Barebones Example

Here is an example of a barebones Scene definition

```go
package myproject_scenes

import (
    "github.com/rocco-gossmann/GoWas/core"
)

// Scene Properties
//==============================================================================
type exampleScene struct {
    secondsSinceStart float64
}

// Implementing core.Scene - Interface(s)
//==============================================================================
func (me *exampleScene) Tick(e *core.EngineState) bool {
    me.secondsSinceStart += e.DeltaTime
    return true
}

// Granting the rest of the project access to the Scene
//==============================================================================
var ExampleScene = exampleScene{
    secondsSinceStart: 0    //<- initial Values go here
}
```

## Example for Using a Scene 

```go
package main
import (
    "MyProject/myproject_scenes"
    "github.com/rocco-gossmann/GoWas"
)

func main() {

    e := GoWas.Init(GoWas.EngineSetup{
        WindowWidth:  160,
        WindowHeight: 144,
    })

    e.Run(&myproject_scenes.ExampleScene)
}
```

