# Bitmap

> [!warning]
> If all you want is to load grapics, from Images, please read the [Graphics & Sound](../Graphics_and_Sound.md) Section first

Sub Package: `GoWas/core`

A `struct` holding a slice of `uint32` pixels and acompanying MetaData

## Constructors

Their are 2 ways to create A Bitmap.

```go
func CreateBitmap(
	pixelsPerLine uint16, 	// the PixelWidth of the Bitmap
	pixelMemory *[]uint32	// The Memory containing all individual pixels
) *Bitmap 

//and 

func CreateBitmapFromCompressed(
	pixelsPerLine uint16, 		// the PixelWidth of the Bitmap
	uncompressedLength int, 	// The Total number of pixels in the uncompressed version
	compressedDate *[]uint32	// The Memory containing complressed version of the Bitmap 
) *Bitmap
```

These 2 are not meant to be used directly for now.



## Methods

Once you have access to an Instance of a `core.Bitmap` you can use the following Methods.

```go

// Get the Full Pixel-Width and Height of the Bitmap
func (me *core.Bitmap) Width() uint16  
func (me *core.Bitmap) Height() uint16 

// To Create a Renderabel Entity out of the Bitmap
func (me *core.Bitmap) MakeEntity() *core.BitmapEntity 
```

## Usage Guide

In order for GoWas to be able to do anything usefull with a Bitmap (Like rendering it), it needs to be either converted.
Or passed to a Processor (A Tileset consturctor for example).


### Bitmap-Enties

If you want to use the Bitmap by itself, it must be
converted into a [`core.BitmapEntity`](./BitmapEntity.md).  
you can do this via the `MakeEntity()` method on the Bitmap.

```go

var bmp *coreBitmap /* = loaded from somewhere */

var entity *core.BitmapEntity = bmp.MakeEntity()

entity.MoveTo(20, 20); //<--Entity should be drawn on coordinates 20, 20
entity.ToCanvas(canvasinstance /*<- also defined somewhere else */)
// See the BitmapEntity Documentation for more options 
```



### TileSets

Bitmaps can also act as a source for a [`gfx.TileSet`](./TileSets.md)  
, which can then be further converted into a [`gfx.TileMap`](./TileMap.md) (WIP)


```go
var bmp *coreBitmap /* = loaded from somewhere */

var tileSet gfx.TileSet;
tileSet.InitFromMapSheet(bmp, 8, 8) //<-create a tileset where each tile is 8x8 pixels

// ... Further process TIleset into a TileMap
var tileMap gfx.TileMap;
tileMap.Init(&tileSet, 10, 10) //<- create a map of 10x10 tiles


```

