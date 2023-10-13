# Engine-State

a structure provided to all called methods of a [Scene](./Scene.md).\
It provides additional informations, that the Scene can use to opperate.

```go
type EngineState struct {
    Mouse     MouseState
    DeltaTime float64
}
```

| Property    | Type                            | Description                                                                                |
| ----------- | ------------------------------- | ------------------------------------------------------------------------------------------ |
| `Mouse`     | [`MouseState`](./MouseState.md) | Mouse-Cursor Position and what Buttons are Pressed/Held/Released                           |
| `DeltaTime` | `float64`                       | how many seconds have passed since the last Tick (1 = 1 second, .5 = half a second, etc. ) |
