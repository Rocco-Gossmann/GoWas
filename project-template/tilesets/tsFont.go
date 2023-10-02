package tilesets

import (
	"GoWasProject/bmps"

	"github.com/rocco-gossmann/GoWas/gfx"
)

var TsFont = gfx.TileSet{}

func init() {
	TsFont.InitFromMapSheet(&bmps.AsciiFont, 8, 8)
}
