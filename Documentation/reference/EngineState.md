# Engine-State

a structure provided to all called methods of a [Scene](../Scenes.md).\
It provides additional informations, that the Scene can use to opperate.

```go
type EngineState struct {
    Mouse     MouseState
    Keyboard  KeyboardState
    DeltaTime float64
}
```

| Property    | Type                                                  | Description                                                                                                                                         |
| ----------- | ----------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Mouse`     | [`MouseState`](../MouseInput.md#mouse-state)          | see [Mouse - Input](../MouseInput.md#mouse-state)<br>Mouse-Cursor Position and what Buttons are Pressed/Held/Released                               |
| `Keyboard`  | [`KeyboardState`](../KeyboardInput.md#keyboard-state) | see [Keyboard - Input](../KeyboardInput.md)<br>What Keys have been Pressed/Held/Released<br>Also contains history of the last 64 entered characters |
| `DeltaTime` | `float64`                                             | how many seconds have passed since the last Tick (1 = 1 second, .5 = half a second, etc. )                                                          |
