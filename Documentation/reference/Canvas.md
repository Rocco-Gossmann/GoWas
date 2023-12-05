# Canvas

<!-- TOC depthfrom:2 -->

- [GetWidth / GetHeight](#getwidth--getheight)
    - [Parameters:](#parameters)

<!-- /TOC -->

The Canvas is the visible area of your application.

It is mainly provided during the [engine-lifescycle](../EngineLifecycle.md) 
via the [Load](../Scenes.md#loadable)- and [Draw](../Scenes.md#drawable)-Hooks.

# Methods

## GetWidth / GetHeight
```go 
func (ca *Canvas) GetWidth() uint16 

func (ca *Canvas) GetHeight() uint16
```

These two functions receive the Dimensions of the current Canvas.

# FillRGBA
```go
func (ec *Canvas) FillRGBA(r, g, b, alpha byte, layerReset CanvasCollisionLayers) 
```
Fills the entire canvas with a given color, constructed of the RGB values.

### Parameters:

| Parameter    | Type                                                | Description                                                                                                      |
|--------------|-----------------------------------------------------|------------------------------------------------------------------------------------------------------------------|
| `r`          | `byte`                                              | Red color channel strength as a value from 0 - 255  (0 being no color, 255 full strength)                        |
| `g`          | `byte`                                              | Green color channel strength as a value from 0 - 255  (0 being no color, 255 full strength)                      |
| `b`          | `byte`                                              | Blue color channel strength as a value from 0 - 255  (0 being no color, 255 full strength)                       |
| `alpha`      | `byte`                                              | How strongly the fill covers the existing pixels on the canvas from 0-255 (0 = no coverage, 255 = full coverage) |
| `layerReset` | [`core.CanvasCollisionLayer`](./CollisionLayers.md) | if set, the given layer are reset to 0                                                                           |



# FillColorA
```go 
func (ec *Canvas) FillColorA(color uint32, alpha byte, layerReset CanvasCollisionLayers)
```
The same as [FillRGBA](#fillrgba) but takes a combined RGB value instead of individual R, G and B channel values.
Therefore it should be a bit more performant.

| Parameter    | Type                                                | Description                                                                                                      |
|--------------|-----------------------------------------------------|------------------------------------------------------------------------------------------------------------------|
| `color`      | `uint32`                                            | The new Color in the format `0xRRGGBB` (The highest byte of the uint32 are not used)
| `alpha`      | `byte`                                              | How strongly the fill covers the existing pixels on the canvas from 0-255 (0 = no coverage, 255 = full coverage) |
| `layerReset` | [`core.CanvasCollisionLayer`](./CollisionLayers.md) | if set, the given layer are reset to 0                                                                           |
