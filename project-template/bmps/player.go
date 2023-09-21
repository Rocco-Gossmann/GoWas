package bmps

import "github.com/rocco-gossmann/GoWas/canvas"

var memory [64]uint32

func init() {

	for a := 0; a < len(memory); a++ {
		memory[a] = 0x00881111
	}
}

var memslice = memory[:]

var PlayerBMP = canvas.CreateBitmap(8, &memslice)
