# Bitmap

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

These 2 are less meant to be used directly.

Instead you should create your Bitmaps via the 
[png2gowasbmp.go](./tools/png2gowasbmp.go) Tool in this repos `tools` folder



