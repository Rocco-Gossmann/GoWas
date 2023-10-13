# TileSets

A Tileset is a [Bitmap](./Bitmap.md), that is split into multiple areas.

Each Area can be drawn independently

## Creation

You first must define, where your Tileset Lives in your application.
You do so, by defining a `var` of the `gfx.TileSet` struct type.

After that you have 2 options to initialize it.

```go
import "github.com/rocco-gossmann/GoWas/gfx"

var myTileSet = gfx.TileSet{}

// When all your Tiles have the same Size and are arranged in a Grid use:
myTileSet.InitFromMapSheet(...)

// Otherwise use to define each tiles position and size by hand 
myTileSet.InitFromSpriteSheet(...) //<- yet to be implemented
```

## Things to Note

> [!important]  
> Things to keep in mind, when using/drawing TileSets
>
> - Tile `0` (most top left) is always blank and will never be drawn to the screen.


# Functions

## InitFromMapSheet(bmp, tilewidth, tileheight)

```go
func (ts *TileSet) InitFromMapSheet(bmp *core.Bitmap, tilewidth, tileheight uint16)
```

This function build a TileSet from a given [Bitmap](./Bitmap.md).

### Parameters:

|            | type           | description                                                               |
|------------|----------------|---------------------------------------------------------------------------|
| bmp        | `*core.Bitmap` | The[Bitmap](./Bitmap.md) that defines all pixels available to the tileset |
| tilewidth  | `uint16`       | the width of an individual tile in pixels                                 |
| tileheight | `uint16`       | the height of an individual tile in pixels                                |


## BlitTo(canvas, tileindex, opts)

```go
func (pTs *TileSet) BlitTo(canvas *core.Canvas, tileindex uint, opts *TilesetBlitOptions) core.CanvasCollisionLayers {
```

Draw a given Tile to the Canvas

### Parameters:

|           | type           | description                                     |
|-----------|----------------|-------------------------------------------------|
| canvas    | `*core.Canvas` | the target, the Map is drawn to                 |
| tileindex | `int`          | see [Tile List Structure](#tile-list-structure) |
| opts | [`*TilesetBlitOptions`](#blitoptions) |  Defines how and where the tile is drawn| 

# Tile List Structure

The Index of a Tile depends on the Type of map you created.
(Right now their are only Map-TileSets so don't worry)

## Map Tilesets
### created via [InitFromMapSheet](#initfrommapsheetbmp-tilewidth-tileheight)
The tiles are arranged left to right first. Then top to bottom.
```
-------------------------------------
|  1  |  2  |  3  |  4  |  5  |  6  |
-------------------------------------
|  7  |  8  |  9  | 10  | 11  | 12  |
-------------------------------------
| 13  | 14  | 15  | 16  | 17  | 18  |
-------------------------------------
```



# Structures
## BlitOptions

```go
type TilesetBlitOptions struct {
	X, Y      int32     // Target X, Y coordinates on the canvas
	Alpha     byte      // transparency of the drawn pixels 0x01 to 0xff (0x00 is equal to 0xff unless Alphazero is set to true)
	Alphazero bool      // If set to true, an Alpha of 0x00 will result in nothing being drawn.
	Layers    core.CanvasCollisionLayers // The collision-Layers, the drawn tile will occupi
}
```

