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

// ------------------------------------------------------------------------------
// Constructors
// ------------------------------------------------------------------------------
func (bmp *Bitmap) constructor(pixelsPerLine uint16, pixelMemory *[]uint32) {
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

func (me *Bitmap) MakeEntity() *BitmapEntity {
	var bmp BitmapEntity
	bmp.canvasBlitOpts.Bmp = me
	bmp.canvasBlitOpts.Clip = &types.Rect{X: 0, Y: 0, W: me.width, H: me.height}
	bmp.visible = true
	return &bmp
}

func (me *Bitmap) Width() uint16  { return me.width }
func (me *Bitmap) Height() uint16 { return me.height }

func CreateBitmap(pixelsPerLine uint16, pixelMemory *[]uint32) *Bitmap {
	bmp := Bitmap{}
	bmp.constructor(pixelsPerLine, pixelMemory)
	return &bmp
}

func CreateBitmapFromCompressed(pixelsPerLine uint16, uncompressedLength int, compressedDate *[]uint32) *Bitmap {
	if pixelsPerLine == 0 {
		panic("bitmap must have at least 1 pixelPerLine")
	}

	compressedLen := len(*compressedDate)
	if compressedLen <= 0 {
		panic("no compressed data given")
	}

	bmp := Bitmap{}
	memory := make([]uint32, uncompressedLength)

	// Decomplress compressed
	memIndex := 0
	for wordIndex := 0; wordIndex < compressedLen && memIndex < uncompressedLength; wordIndex++ {
		var curWord = (*compressedDate)[wordIndex]
		meta := (curWord & 0xfe000000 >> 25) + 1
		color := curWord & 0x01ffffff

		for meta > 0 {
			memory[memIndex] = color
			memIndex++
			meta--
		}
	}

	bmp.constructor(pixelsPerLine, &memory)

	return &bmp
}
