package core

import (
	"fmt"
)

type tileMapOpts struct {
	X, Y      int32 // Where to blit it on the screen
	Alpha     CanvasAlpha
	AlphaZero bool
}

type TileMap struct {
	init     bool
	opts     tileMapOpts
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

func (tm *TileMap) Init(ts *TileSet, width, height uint32) *TileMap {

	if tm == nil {
		panic("tilemap is nil")
	}

	if tm.init {
		panic("cant initialize a map twice")
	}

	if width == 0 {
		panic("map width must be bigger than 0")
	}
	if height == 0 {
		panic("map height must be bigger than 0")
	}

	tm.SetTileSet(ts)

	tm.mh = height
	tm.mw = width
	tm.memory = make([]byte, tm.mw*tm.mh)

	tm.init = true
	return tm
}

// -----------------------------------------------------------------------------
// Setters
// -----------------------------------------------------------------------------
func (me *TileMap) SetTileSet(ts *TileSet) *TileMap {
	if ts != nil {

		tileCount := ts.TileCount()

		if tileCount > 255 {
			panic("Maps can't use Tilesets that have more than 255 tiles")
		}

		if ts.Type() != TILESET_TYPE_MAP {
			panic(" can't create TileMap from a Tileset, that is not of Type TILESET_TYPE_MAP")
		}

		me.ts = ts
		me.accessTc = byte(min(255, tileCount))

	} else {
		me.ts = nil
	}

	return me
}

// Tiles
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

// Display
func (me *TileMap) AlphaSet(a CanvasAlpha) *TileMap {
	me.opts.Alpha = a
	me.opts.AlphaZero = a == CANV_ALPHA_NONE || me.opts.AlphaZero
	return me
}

func (me *TileMap) AlphaReset() *TileMap {
	me.opts.Alpha = CANV_ALPHA_NONE
	me.opts.AlphaZero = false
	return me
}

func (me *TileMap) MoveTo(x, y int32) *TileMap {

	me.opts.X = x
	me.opts.Y = y

	return me
}
func (me *TileMap) MoveBy(x, y int32) *TileMap {

	me.opts.X += x
	me.opts.Y += y

	return me
}

// -----------------------------------------------------------------------------
// Getters
// -----------------------------------------------------------------------------

func (me *TileMap) HasTileSet() bool { return me.ts != nil }

func (me *TileMap) X() int32           { return me.opts.X }
func (me *TileMap) Y() int32           { return me.opts.Y }
func (me *TileMap) XY() (int32, int32) { return me.opts.X, me.opts.Y }
func (me *TileMap) Alpha() CanvasAlpha { return me.opts.Alpha }

// -----------------------------------------------------------------------------
// Actions
// -----------------------------------------------------------------------------

func (me *TileMap) ToCanvas(ca *Canvas) {

	if me.ts == nil {
		return
	}

	caw, cah := ca.GetWidth(), ca.GetHeight()
	tc := me.ts.TileCount()
	tw, th := uint32(me.ts.GetTileWidth()), uint32(me.ts.GetTileHeight())
	offsetX, offsetY := int32(0), int32(0)
	startX, startY := uint32(0), uint32(0)

	offsetX = int32(me.opts.X) * -1
	offsetY = int32(me.opts.Y) * -1

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

	bopts := TilesetBlitOptions{
		Alpha:     me.opts.Alpha,
		Alphazero: me.opts.AlphaZero,
	}

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
