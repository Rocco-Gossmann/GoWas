package bmps

import "github.com/rocco-gossmann/GoWas/core"

var memory [64]uint32

func init() {
	for a := 0; a < len(memory); a++ {
		memory[a] = 0x01881111
	}
}

var memslice = memory[:]

var PlayerBMP = core.CreateBitmap(8, &memslice)
