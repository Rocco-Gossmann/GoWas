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
## Properties
| Property    | Type                                                  | Description                                                                                                                                         |
|-------------|-------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| `Mouse`     | [`MouseState`](../MouseInput.md#mouse-state)          | see [Mouse - Input](../MouseInput.md#mouse-state)<br>Mouse-Cursor Position and what Buttons are Pressed/Held/Released                               |
| `Keyboard`  | [`KeyboardState`](../KeyboardInput.md#keyboard-state) | see [Keyboard - Input](../KeyboardInput.md)<br>What Keys have been Pressed/Held/Released<br>Also contains history of the last 64 entered characters |
| `DeltaTime` | `float64`                                             | how many seconds have passed since the last Tick (1 = 1 second, .5 = half a second, etc. )                                                          |
| `Text` | [`TextDisplay`](./TextDisplay.md) | Direct access to the Text-Display controlling the [UI/Text-Layer](../Graphics_and_Sound.md#display-layers)


## Functions

### Layer-Control

| Function               | Description                    |
|------------------------|--------------------------------|
| `EnableTextLayer()`    | Makes the Text-Layer visible   |
| `DisableTextLayer()`   | Make the  Text-Layer invisible |
|                        |                                |
| `EnableMap2Layer()`    | Makes the Map2-Layer visible   |
| `DisableMap2Layer()`   | Make the  Map2-Layer invisible |
|                        |                                |
| `EnableSpriteLayer()`  | Makes the Sprite-Layer visible   |
| `DisableSpriteLayer()` | Make the  Sprite-Layer invisible |
|                        |                                |
| `EnableMap1Layer()`    | Makes the Map1-Layer visible   |
| `DisableMap1Layer()`   | Make the  Map1-Layer invisible |
