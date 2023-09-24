package gfx

type TileMap struct {
	ts     *TileSet
	mw, mh uint32
	tiles  []uint16
}
