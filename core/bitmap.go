package core

import (
	"fmt"
	"github.com/rocco-gossmann/GoWas/types"
)

type Bitmap struct {
	width, height uint16
	MemoryBuffer  Buffer
}

type BitmapFlag uint32

const (
	BMP_OPAQUE BitmapFlag = 0x01000000 // Set on the Bitmap itself. A Pixel can either be drawn or not, their is no true Alpha channel

	// ColorChannels oBitmapFlag f a BMP
	BMP_CHANNEL_RED   BitmapFlag = 0x00ff0000
	BMP_CHANNEL_GREEN BitmapFlag = 0x0000ff00
	BMP_CHANNEL_BLUE  BitmapFlag = 0x000000ff
)

func (b *Bitmap) Width() uint16  { return b.width }
func (b *Bitmap) Height() uint16 { return b.height }
func (b *Bitmap) PPL() uint16    { return b.MemoryBuffer.PixelPerLine }
func (b *Bitmap) Pixels() int    { return len((*(*b).MemoryBuffer.Memory)) }

func (b *Bitmap) CopyFrom(src *Bitmap, dstx, dsty int32, alpha byte, srcclip *types.Rect) {
	panic("//TODO: Implement")
}


func (bmp *Bitmap) Init(pixelsPerLine uint16, pixelMemory *[]uint32) {
	if bmp == nil {
		panic("bmp can't be nil")
	}

	if pixelsPerLine == 0 {
		panic("bitmap must have at least 1 pixelPerLine")
	}

	memoryLen := len(*pixelMemory)
	missaligment := memoryLen % int(pixelsPerLine)

	if missaligment > 0 {
		panic(fmt.Sprintf(
			"Memory does not align width given pixels per line. %v more pixels required to fill out last line",
			pixelsPerLine-uint16(missaligment),
		))
	}

	bmp.width = pixelsPerLine
	bmp.height = uint16(memoryLen / int(pixelsPerLine))
	bmp.MemoryBuffer.Memory = pixelMemory
	bmp.MemoryBuffer.PixelPerLine = pixelsPerLine
}

func CreateBitmap(pixelsPerLine uint16, pixelMemory *[]uint32) Bitmap {
	bmp := Bitmap{}
	bmp.Init(pixelsPerLine, pixelMemory)
	return bmp
}
