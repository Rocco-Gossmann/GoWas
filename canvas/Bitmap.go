package canvas

import "fmt"

type Bitmap struct {
	width, height uint16
	MemoryBuffer  Buffer
}

type BitmapFlag uint32

const (
	BMP_OPAQUE = 0x01000000 // Set on the Bitmap itself. A Pixel can either be drawn or not, their is no true Alpha channel

	// ColorChannels of a BMP
	BMP_CHANNEL_RED   = 0x00ff0000
	BMP_CHANNEL_GREEN = 0x0000ff00
	BMP_CHANNEL_BLUE  = 0x000000ff
)

func (b *Bitmap) Width() uint16  { return b.width }
func (b *Bitmap) Height() uint16 { return b.height }
func (b *Bitmap) PPL() uint16    { return b.MemoryBuffer.PixelPerLine }
func (b *Bitmap) Pixels() int    { return len((*(*b).MemoryBuffer.Memory)) }

func CreateBitmap(pixelsPerLine uint16, pixelMemory *[]uint32) Bitmap {

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

	bmp := Bitmap{
		width:  pixelsPerLine,
		height: uint16(memoryLen / int(pixelsPerLine)),
	}
	bmp.MemoryBuffer.Memory = pixelMemory
	bmp.MemoryBuffer.PixelPerLine = pixelsPerLine

	return bmp
}
