package core

import (
	"fmt"

	"github.com/rocco-gossmann/GoWas/types"
)

type TileSetType uint8

const (
	TILESET_TYPE_UNINITIALIZED TileSetType = 0
	TILESET_TYPE_MAP           TileSetType = 1
	TILESET_TYPE_SPRITESHEET   TileSetType = 2
)

type TileSet struct {
	gfx    *Bitmap
	tiles  []types.Rect
	tw, th uint16
	tstype TileSetType
}

func (ts TileSet) GetTileWidth() uint16  { return ts.tw }
func (ts TileSet) GetTileHeight() uint16 { return ts.th }
func (ts TileSet) Type() TileSetType     { return ts.tstype }
func (ts TileSet) TileCount() int        { return len(ts.tiles) }

type TilesetBlitOptions struct {
	X, Y      int32
	Alpha     CanvasAlpha
	Alphazero bool
	Layers    CanvasCollisionLayers
}

var defaultOpts = TilesetBlitOptions{}

func (pTs *TileSet) BlitTo(canvas *Canvas, tileindex int, pOpts *TilesetBlitOptions) CanvasCollisionLayers {
	if pOpts == nil {
		pOpts = &defaultOpts
	}

	var opts = (*pOpts)

	if (*pTs).gfx == nil || len((*pTs).tiles) > tileindex {

		return canvas.Blit(&CanvasBlitOpts{
			Bmp:       (*pTs).gfx,
			X:         opts.X,
			Y:         opts.Y,
			Alpha:     opts.Alpha,
			Alphazero: opts.Alphazero,
			Layers:    opts.Layers,
			Clip:      &((*pTs).tiles[tileindex]),
		})

	} else {
		return CANV_CL_NONE
	}
}

func (ts *TileSet) InitFromMapSheet(bmp *Bitmap, tilepixelwidth, tilepixelheight uint16) {

	if tilepixelwidth == 0 {
		panic("'tilepixelwidth' must be bigger than 0")
	}
	if tilepixelheight == 0 {
		panic("'tilepixelheight' must be bigger than 0")
	}

	bw, bh := bmp.Width(), bmp.Height()

	if bw%uint16(tilepixelwidth) > 0 {
		panic(fmt.Sprintf("tileset missaligned: the given bitmaps width must be a multiple of %vpx (tilepixelwidth)", tilepixelwidth))
	}
	if bh%uint16(tilepixelheight) > 0 {
		panic(fmt.Sprintf("tileset missaligned: the given bitmaps height must be a multiple of %vpx (tilepixelheight)", tilepixelheight))
	}

	ts.tw, ts.th = tilepixelwidth, tilepixelheight
	mw, mh := uint16(bw/tilepixelwidth), uint16(bh/tilepixelheight)

	ts.gfx = bmp
	ts.tiles = make([]types.Rect, mw*mh)
	ts.tstype = TILESET_TYPE_MAP

	for y := uint16(0); y < mh; y++ {
		for x := uint16(0); x < mw; x++ {
			var i = y*mw + x
			ts.tiles[i].X = x * tilepixelwidth
			ts.tiles[i].Y = y * tilepixelwidth
			ts.tiles[i].W = tilepixelwidth
			ts.tiles[i].H = tilepixelheight
		}
	}
}
