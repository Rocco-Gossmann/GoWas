package gfx

import "fmt"

type TileMap struct {
	init   bool
	ts     *TileSet
	memory []byte
	mw, mh uint32
}

func (pTm *TileMap) validate() {
	if !(*pTm).init {
		panic("tilemap was not initialized")
	}
}

func (tm *TileMap) Init(pTs *TileSet, width, height uint32) *TileMap {

	if tm == nil {
		panic("tilemap is nil")
	}

	var ts = (*pTs)

	if tm.init {
		panic("cant initialize a map twice")
	}

	if ts.TileCount() > 255 {
		panic("Maps can't use Tilesets that have more than 255 tiles")
	}

	if width == 0 {
		panic("map width must be bigger than 0")
	}
	if height == 0 {
		panic("map height must be bigger than 0")
	}

	if ts.Type() != TILESET_TYPE_MAP {
		panic(" can't create TileMap from a Tileset, that is not of Type TILESET_TYPE_MAP")
	}

	tm.mh = height
	tm.mw = width
	tm.ts = pTs

	tm.memory = make([]byte, tm.mw*tm.mh)

	tm.init = true
	return tm
}

func (me *TileMap) SetMap(mapData []byte) *TileMap {
	me.validate()

	if len(mapData) != len((*me).memory) {
		panic(fmt.Sprintf("SetMap length missmatch. expected: %v bytes got: %v bytes", len(me.memory), len(mapData)))
	}

	copy(me.memory, mapData)
	return me
}

func (me *TileMap) SetTile(x, y uint32, tileIndex byte) *TileMap {
	me.validate()

	var mapIndex = y*me.mw + x
	if len((*me).memory) <= int(mapIndex) {
		panic("can't set a tile that is not on the Map, check your x and y coordinates and make sure, they are within the maps with and height")
	}

	// The Tilecount was checked during Init, so it is not bigger than 255.
	//                                this byte-cast should be fine here
	(*me).memory[mapIndex] = tileIndex % byte((*((*me).ts)).TileCount())

	return me
}
