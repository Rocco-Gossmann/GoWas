package tilesets

import (
	"GoWasProject/bmps"

	"github.com/rocco-gossmann/GoWas/gfx"
)

var TsFont = gfx.TileSet{}

func init() {
	TsFont.InitFromMapSheet(&bmps.DebugTiles, 8, 8)
}
