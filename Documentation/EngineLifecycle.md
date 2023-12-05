# The Engine-Lifecycle

![Engine-Lifecycle-img](./_img/Engine-Lifecycle.png)

The Engines Lifecycle may look complicated at first but you will only interact
at a few points with it. (All the Blue-Methods in the graphic above)

To make things easier to understand, I'll use a Vehicle (like a Car or Bike) as an
analogy.

## 1.) Initialization (Building the Frame)

```
GoWas.Init(GoWas.EngineSetup)
```

Before we can drive something, we must first define what we want to drive. In a sense, 
we are defining the Frame of our Vehicle. You define the size of your
Viewport here. The Viewport could be seen as the Windshield of a car. Letting
the driver see what is happening.

This is done via:

```go
import "github.com/rocco-gossmann/GoWas"

func main() {
    var engine := GoWas.Init(GoWas.EngineSetup{
        //WindowWidth uint16,
        //WindowHeight uint16,
    })
}
```

For more details on how to initialize the GoWas-Engine, please read the
[Engine-Reference](./Engine.md)

## 2.) Initial-Scene-Setup (Defining your Vehicle)

Now that we have a Frame, we need to define, what our Vehicle will be filled
with. What levers or wheels are on it?

This is where the [`core.Scene`](./Scenes.md) comes in.

The GoWas Core defines 4 Interfaces.

```go
type Loadable interface {
    Load(e *EngineState, ec *Canvas)
}
```

```go
type Tickable interface {
    Tick(e *EngineState) bool
}
```

```go
type Drawable interface {
    Draw(e *EngineState, ec *Canvas)
}
```

```go
type Unloadable interface {
    Unload(e *EngineState) *struct{}
}
```

> [!important]\
> Any struc{} that implements the `Drawable` and/or `Tickable` interface Will be
> a valid `core.Scene`
>
> You don't have to implemet all 4 interfaces, if you don't need them.

### 2a.) Loading the Scene (Getting into the Vehicle and turning the ignition key)

&nbsp;<pre>Load(*[core.EngineState](./reference/EngineState.md), *[core.Canvas](./reference/Canvas.md))</pre>  
Is called before the Engine starts running its main loop.

Here is, where you should initialize your resources, that the Engine will use
when it is running.

You can also use it to define the initial State of the Viewport.

In our analogy, this is the point, where you fill the Cockpit or driver
cabin of your Vehicle.\
Do you want to steer it with a Wheel or a Bike-Handle?

### 2b.) Ticking (Defining what happens with each stroke of the Engine)

&nbsp;<pre>Tick(*[core.EngineState](./reference/EngineState.md)) bool </pre>  

The Cockpit is set up, you turned the Key, and now the Engine is running and each tick/stroke.
your Scenes `Tick(*core.EngineState) bool` - Method is called.

Here, you can do whatever you want with the resources you set up in your Scenes `Load(...)`-Method.

The `Tick(...)` method must however return a `boolean` to tell the engine
if it should keep running like this.

> [!warning]
> This method must return `true` unless you want to switch to another Scene or Stop the Engine  
> (see [Scene-Transitions](./SceneTransitions.md) for more details)



### 2c.) Drawing (Showing the driver what is going on)

&nbsp;<pre>Draw(*[core.EngineState](./reference/EngineState.md))</pre>  

After each Tick, your Scenes `Draw(*core.EngineState, *core.Canvas)` Method is called.
`*core.Canvas` is our viewport, that can be manipulated in this step.

It is more likely, however, that you will just pass it through to some components, you 
set up in your scenes `Load(...)` Method.

> [!note]  
> Drawing and Ticking are separated.
> once your scenes `Tick(...)` method returns `false` the call to `Draw(...)` is skipped. 

### 3.) Unloading (Shutting down the Vehicle and clearing the driver's cabin) 

&nbsp;<pre>Unload(*[core.EngineState](./reference/EngineState.md)) *[core.Scene](./Scenes.md) </pre>  

Once your scenes `Tick(...)`-Method returns a `false`, then Engine will call your scenes `Unload(...)` method.  
Should your Scene not have an `Unload(...)` method, the Engine will shut down.

The unload method is meant to clear and free all resources, that were set up/created during your Scenes 
`Load(...)`, `Tick(...)` or `Draw(...)` method calls.

The `Unload(...)` method itself has to return either a `core.Scene` pointer or `nil`.

if a `*core.Scene` is returned, then the Engine will first call its `Load(...)` Method (should you define one) and then continue to run its cycle. 

Should you want to stop the Engines execution, then `Unload(...)` must return `nil`.