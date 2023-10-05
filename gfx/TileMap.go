package gfx

import (
	"fmt"

	"github.com/rocco-gossmann/GoWas/core"
	"github.com/rocco-gossmann/GoWas/types"
)

type TileMap struct {
	init     bool
	ts       *TileSet
	memory   []byte
	mw, mh   uint32
	tsOffset int
	accessTc byte
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
	tm.accessTc = byte(min(255, ts.TileCount()))

	tm.init = true
	return tm
}

// -----------------------------------------------------------------------------
// Setters
// -----------------------------------------------------------------------------
func (me *TileMap) SetTileSetOffset(o int) *TileMap {
	var tc = me.ts.TileCount()

	if o < 0 {
		me.tsOffset = tc + (o % tc)
	} else {
		me.tsOffset = o % tc
	}

	return me
}

func (me *TileMap) Clear(tileIndex byte) *TileMap {
	for i, _ := range me.memory {
		me.memory[i] = tileIndex
	}

	return me
}

func (me *TileMap) SetSequence(sequence string, xoffset uint16, yoffset uint16, ignoreZero bool) *TileMap {

	var offset = int(uint32(yoffset)*me.mw + uint32(xoffset))
	var ml = len(me.memory)

	ra := []rune(sequence)
	if ignoreZero {
		for i, chr := range ra {
			var m = byte(min(1, int32(chr)))
			me.memory[(i+offset)%ml] = (byte(chr&0x000000ff) * m) +
				(me.memory[(i+offset)%ml] * (1 - m))
		}
	} else {
		for i, chr := range ra {
			me.memory[(i+offset)%ml] = byte(chr & 0x000000ff)
		}
	}

	return me
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
	if len(me.memory) <= int(mapIndex) {
		panic("can't set a tile that is not on the Map, check your x and y coordinates and make sure, they are within the maps with and height")
	}

	// The Tilecount was checked during Init, so it is not bigger than 255.
	//                                this byte-cast should be fine here
	me.memory[mapIndex] = tileIndex % me.accessTc

	return me
}

// -----------------------------------------------------------------------------
// Getters
// -----------------------------------------------------------------------------
type ToCanvasOpts struct {
	Scroll types.Point
}

func (me *TileMap) ToCanvas(ca *core.Canvas, opts *ToCanvasOpts) {

	me.validate()

	caw, cah := ca.GetWidth(), ca.GetHeight()
	tc := me.ts.TileCount()
	tw, th := uint32(me.ts.GetTileWidth()), uint32(me.ts.GetTileHeight())
	offsetX, offsetY := int32(0), int32(0)
	startX, startY := uint32(0), uint32(0)

	if opts != nil {
		offsetX = int32(opts.Scroll.X) * -1
		offsetY = int32(opts.Scroll.Y) * -1
	}

	// Check for overshoots Horizontal
	if offsetX < 0 {
		var swOX = (offsetX * -1)
		var ox = swOX % int32(tw)
		var tox = uint32(swOX-ox) / tw
		startX = tox
	}

	// Check for overshoots Vertical
	if offsetY < 0 {
		var swOY = (offsetY * -1)
		var oy = swOY % int32(th)
		var toy = uint32(swOY-oy) / th
		startY = toy
	}

	mw := min(me.mw*tw, uint32(caw)+tw) / tw
	mh := min(me.mh*th, uint32(cah)+th) / th

	//	dstX, dstY := uint16((me.mw+overShootX)*tw), uint16((me.mh+overShootY)*th)

	bopts := TilesetBlitOptions{}

	ti, mi := 0, byte(0)
	for y := startY; y < startY+mh; y++ {
		bopts.Y = int32(y*th) + offsetY
		for x := startX; x < startX+mw; x++ {
			bopts.X = int32(x*tw) + offsetX

			mti := (int((y%me.mh)*me.mw + (x % me.mw)))
			mi = me.memory[mti] % me.accessTc
			om := 1
			if mi == 0 {
				om = 0
			}
			ti = (int(mi) + (me.tsOffset * om)) % tc
			if ti > 0 {
				me.ts.BlitTo(ca, ti, &bopts)
			}
		}
	}
}
