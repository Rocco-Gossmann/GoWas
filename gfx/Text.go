package gfx

import (
	"math"
	"strings"

	"github.com/rocco-gossmann/GoWas/core"
	"github.com/rocco-gossmann/GoWas/ressource"
)

type TextDisplay struct {
	// Internals
	canvasWidth  uint16
	canvasHeight uint16
	charsPerLine uint16
	lines        uint16
	ts           *TileSet

	// Internals
	mp *TileMap

	// Formating
	cursorx, cursory uint16
	wrap             bool
}

// ==============================================================================
// Constructors
// ==============================================================================
func InitTextDisplay(ca *core.Canvas) *TextDisplay {

	if ca == nil {
		panic("cant initiate TextDisplay without a Canvas to initialize it for.")
	}

	text := TextDisplay{}
	text.canvasWidth = ca.GetWidth()
	text.canvasHeight = ca.GetHeight()

	text.ts = &TileSet{}
	text.ts.InitFromMapSheet(ressource.AsciiFontBMP, 8, 8)

	text.setFont(text.ts)
	text.wrap = true

	text.mp.SetTileSetOffset(64)

	return &text
}

// ==============================================================================
// Setters
// ==============================================================================
func (me *TextDisplay) SetCursor(x, y int32) *TextDisplay {
	x = x % int32(me.charsPerLine)
	y = y % int32(me.lines)

	if x < 0 {
		me.cursorx = uint16(int32(me.charsPerLine) + x)
	} else {
		me.cursorx = uint16(x)
	}

	if y < 0 {
		me.cursory = uint16(int32(me.lines) + y)

	} else {
		me.cursory = uint16(y)

	}

	return me
}

func (me *TextDisplay) Clear(numOfCharacters uint) *TextDisplay {
	cx, cy := me.cursorx, me.cursory
	return me.
		Echo(strings.Repeat(" ", int(numOfCharacters))).
		SetCursor(int32(cx), int32(cy))
}

func (me *TextDisplay) Echo(text string) *TextDisplay {

	seq := make([]rune, me.lines*me.charsPerLine)

	seqIndex := 0
	cursorx, cursory := uint16(me.cursorx), uint16(me.cursory)
	printValid := true
	for _, chr := range []rune(text) {

		if cursorx >= me.charsPerLine {
			if me.wrap {
				cursorx = 0
				cursory++
			} else {
				printValid = false
			}
		}

		if chr == '\n' {
			var dst = me.charsPerLine - cursorx
			for dst > 0 {
				seq[seqIndex] = 0
				seqIndex++
				dst--
			}

			if me.wrap {
				cursorx = me.charsPerLine
			} else {
				cursory++
				cursorx = 0
				printValid = true
			}

		} else if printValid {
			seq[seqIndex] = chr
			seqIndex++
			cursorx++
		}
	}

	seq = seq[0:seqIndex]
	me.mp.SetSequence(string(seq), me.cursorx, me.cursory, true)

	me.cursorx = cursorx
	me.cursory = cursory

	return me
}

func (me *TextDisplay) SetWrap(wrap bool) *TextDisplay {
	me.wrap = wrap
	return me
}

func (me *TextDisplay) setFont(ts *TileSet) *TextDisplay {
	tw, th := ts.GetTileWidth(), ts.GetTileWidth()

	me.charsPerLine = uint16(math.Floor(float64(me.canvasWidth) / float64(tw)))
	me.lines = uint16(math.Floor(float64(me.canvasHeight) / float64(th)))

	me.mp = &TileMap{}
	me.mp = me.mp.Init(me.ts, uint32(me.charsPerLine), uint32(me.lines))

	return me
}

func (me *TextDisplay) MoveTo(x, y int32) *TextDisplay {
	me.mp.MoveTo(x*-1, y*-1)
	return me
}

func (me *TextDisplay) MoveBy(x, y int32) *TextDisplay {
	me.mp.MoveBy(x, y)
	return me
}

func (me *TextDisplay) SetAlpha(a byte) *TextDisplay {
	me.mp.SetAlpha(a)
	return me
}

// ==============================================================================
// Actions
// ==============================================================================
func (me *TextDisplay) ToCanvas(ca *core.Canvas) {
	me.mp.ToCanvas(ca)
}
