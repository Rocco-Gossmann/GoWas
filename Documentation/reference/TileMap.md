# Tile Maps
 
Tilemaps are a 2-dimensional grid of bytes, where each entry points to an Entry of a [TileSet](./TileSets.md)

![TileMap Graphic](../_img/tilemap.png)


<!-- TOC -->

- [Tile Maps](#tile-maps)
- [Build in Map Layers](#build-in-map-layers)
- [Creating your own Layers](#creating-your-own-layers)

<!-- /TOC -->

# Build in Map Layers

In the [ Display Layers](../Graphics_and_Sound.md#display-layers), You get access to two TileMaps out of the Box.

The Maps default Size is 32x32 tiles.
You probably want to adjust that, depending on your chosen resolution and the size 
of the [TileSet](./TileSets.md) you plan to use. 
(See [Engine-Setup](./Engine.md#enginesetup) for more info on how to change the Maps size)

Before you can enable any of the two Build in Maps, you must first define a [TileSet](./TileSets.md), that your chosen layer will use

The following would need to happen during your Scenen [Load-Lifecycle-Hook](../Scenes.md#loadable)

```go
func (me *ExampleScene) Load(state *core.EngineState, canvas *core.Canvas) {
    var tileSet := core.TileSet{}
    tileSet.InitFromMapSheet( 
        /* your core.Bitmap */, 
        /* PixelWidth of each Tile */,
        /* PixelHeight of each Tile */,
    );
    // ...
```
with That, you can then enable your chosen Map-Layer through the [Engine](./EngineState.md#layer-control) State](./EngineState.md#layer-control)
```go
    // ... 
    state.EnableMap1Layer(&tileSet)
    // ...
```
And Start filling it with Data
```go
    //...
    state.Map1.SetMap([]byte{ /* your Maps Tile definitions goes here */ })
}
```




# Creating your own Layers

