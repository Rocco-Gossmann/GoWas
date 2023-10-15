# Mouse-Input

The Mouse-Input is bound to the [Engine-State](./reference/EngineState.md) as
the `Mouse` property.
Thus it is provided to your [Scene](./Scenes.md) through the
[Engine-Lifecycle](./EngineLifecycle.md).

## Mouse-State

The [Engine-States](./reference/EngineState.md) `Mouse` Propert provides the
Following Structure.

```go
type MouseState struct {
    X, Y          uint16

    Pressed       uint32
    Held          uint32
    Released      uint32
    PressedOrHeld uint32
}
```

## Reading the Position

The Cursors position is based on your Viewports/Canvases resolution provided.
You can set them in the [`GoWas.Init(...)`](./Engine.md) call in your `main`
function.

The following Example created a `DemoScene` as described in
[Scenes](./Scenes.md).\
It registers a [BitmapEntity](./reference/BitmapEntity.md) called
`CursorEntity`. to show the Mouse-Cursor.

```go
type demoScene struct {
    CursorEntity *core.BitmapEntity
}
// ...implement the core.Drawable interface
func (s *demoScene) Draw(e *core.EngineState, c *core.Canvas) {

    //...
    s.CursorEntity.MoveTo(
        int32(e.Mouse.X), //<- Taking the x/y from the Mouse-State
        int32(e.Mouse.Y), //   And moving the CursorEntity to it
    )

    if e.Mouse.X > 0 || e.Mouse.Y > 0 { // showing the entity only
        s.CursorEntity.ToCanvas(ca)     // if it is not stuck to the
                                        // top left cornor
    }
}
```

## Reading the Buttons

Mouse Buttons are provided via the following 4 Properties on the
[Mouse-State](#mousestate).

| Property        | Description                                                       |
| --------------- | ----------------------------------------------------------------- |
| `Pressed`       | Buttons that have not been pressed or held last cycle and are now |
| `Held`          | Buttons that have been pressed last cycle and are still pressed   |
| `Released`      | Buttons that where pressed or held last cycle and no longer are   |
| `PressedOrHeld` | the Pressed and Held property combined for convinience            |

You interact with these via `Bitwise` opperations. For that you are provided 
the `core.MOUSE_BTN1` to `core.MOUSE_BTN32` constants.

```go


func (me *demoScene) Tick(e *core.EngineState) bool {

    if(e.Mouse.Pressed & core.MOUSE_BTN1 > 0) {
        // Do something if the Left mouse-button was just pressed this cycle 
        //...
    }

    if(e.Mouse.PressedOrHeld & core.MOUSE_BTN1 > 0) {
        // Do something if the Left mouse-button is currently held down 
        //...
    }

    if(e.Mouse.Released & core.MOUSE_BTN1 > 0) {
        // Do something if the Left mouse-button was released this cycle
        //...
    }

    // Combined Button Presses
    const leftAndRight = core.MOUSE_BTN1 | core.MOUSE_BTN3

    if(e.Mouse.Held & leftAndRight > 0) {
        // Do something if either the Left OR Right mouse-button are held
        //...
    }

    if(e.Mouse.Held & leftAndRight == leftAndRight) {
        // Do something if the Left AND Right mouse-button are held together
        //...
    }

}

```



