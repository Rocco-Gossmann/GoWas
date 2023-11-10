# Engine-State

a structure provided to all called methods of a [Scene](../Scenes.md).\
It provides additional pieces of information and functions, that the Scene can use to operate.

```go
type EngineState struct {
    DeltaTime float64
    Mouse     MouseState
    Keyboard  KeyboardState
    // ...
}
```
<!-- TOC -->

- [Engine-State](#engine-state)
- [Misc Properties](#misc-properties)
    - [Properties](#properties)
- [IO Properties](#io-properties)
    - [Properties](#properties)
- [Display Layer Control](#display-layer-control)
    - [Properties](#properties)
    - [Functions](#functions)

<!-- /TOC -->


# Utility 
## Properties
| Property    | Type      | Description                                                                                |
|-------------|-----------|--------------------------------------------------------------------------------------------|
| `DeltaTime` | `float64` | how many seconds have passed since the last Tick (1 = 1 second, .5 = half a second, etc. ) |



# IO
## Properties
| Property   | Type                                                  | Description                                                                                                                                         |
|------------|-------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| `Mouse`    | [`MouseState`](../MouseInput.md#mouse-state)          | see [Mouse - Input](../MouseInput.md#mouse-state)<br>Mouse-Cursor Position and what Buttons are Pressed/Held/Released                               |
| `Keyboard` | [`KeyboardState`](../KeyboardInput.md#keyboard-state) | see [Keyboard - Input](../KeyboardInput.md)<br>What Keys have been Pressed/Held/Released<br>Also contains history of the last 64 entered characters |




# Display Layer Control

## Properties
| Property | Type                              | Description                                                                                                |
|----------|-----------------------------------|------------------------------------------------------------------------------------------------------------|
| `Text`   | [`TextDisplay`](./TextDisplay.md) | Direct access to the Text-Display controlling the [UI/Text-Layer](../Graphics_and_Sound.md#display-layers) |
| `Map1`   | [`TileMap`](./TileMap.md)         | Direct access to the [Map1/Backround TileMap](../Graphics_and_Sound.md#display-layers)                     |
| `Map2`   | [`TileMap`](./TileMap.md)         | Direct access to the [Map2/Foreground TileMap](../Graphics_and_Sound.md#display-layers)                    |


## Functions
| Function                         | Description                      |
|----------------------------------|----------------------------------|
| `EnableTextLayer()`              | Makes the Text-Layer visible     |
| `DisableTextLayer()`             | Make the  Text-Layer invisible   |
|                                  |                                  |
| `EnableMap2Layer(*core.TileSet)` | Makes the Map2-Layer visible     |
| `DisableMap2Layer()`             | Make the  Map2-Layer invisible   |
|                                  |                                  |
| `EnableSpriteLayer()`            | Makes the Sprite-Layer visible   |
| `DisableSpriteLayer()`           | Make the  Sprite-Layer invisible |
|                                  |                                  |
| `EnableMap1Layer(*core.TileSet)` | Makes the Map1-Layer visible     |
| `DisableMap1Layer()`             | Make the  Map1-Layer invisible   |
