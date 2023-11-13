package core

import (
	"fmt"
	"math"
	"strings"
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
func InitTextDisplay(ca *Canvas) *TextDisplay {

	if ca == nil {
		panic("cant initiate TextDisplay without a Canvas to initialize it for.")
	}

	text := TextDisplay{}
	text.canvasWidth = ca.GetWidth()
	text.canvasHeight = ca.GetHeight()

	text.ts = &TileSet{}
	text.ts.InitFromMapSheet(AsciiFontBMP, 8, 8)

	text.SetFont(text.ts)
	text.mp.SetTileSetOffset(64)
	text.wrap = true

	return &text
}

// ==============================================================================
// Getters
// ==============================================================================
func (me *TextDisplay) Wrap() bool {
	return me.wrap
}

func (me *TextDisplay) Cursor() (uint16, uint16) {
	return me.cursorx, me.cursory
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

	var displayTextArrLength = me.lines * me.charsPerLine
	seq := make([]rune, displayTextArrLength)

	seqIndex := 0
	cursorx, cursory := uint16(me.cursorx), uint16(me.cursory)
	printValid := true
	for _, chr := range []rune(text) {
		if cursory*me.charsPerLine+cursorx >= displayTextArrLength {
			break
		}

		// If line overflow is tetected
		if cursorx >= me.charsPerLine {
			// IF line wrap is turned on
			if me.wrap {
				// Set cursor to start of next line
				cursorx = 0
				cursory++
			} else {
				// stop printing, until you encounter a "\n"
				printValid = false
			}
		}

		// When encountering a Line-Break "\n"
		if chr == '\n' {
			// Characters left to render
			var dst = me.charsPerLine - cursorx

			// Fill remaining character of line with 0
			for dst > 0 {
				seq[seqIndex] = 0
				seqIndex++
				dst--
			}

			// If line wrap is on
			if me.wrap {
				// Set Cursor X after then end of the line and let the code above
				// handle the rest
				cursorx = me.charsPerLine

			} else {
				// else set cursor to start of the next line
				// and enable printing again
				cursory++
				cursorx = 0
				printValid = true
			}

		} else if printValid {

			defer (func() {
				if r := recover(); r != nil {
					fmt.Printf(
						"hit panic(%v) invalid index: %v on cursor: %vx%v for char(%c) of Text with length: %v\n",
						r,
						seqIndex,
						cursorx,
						cursory,
						chr,
						len(seq),
					)
				}
				seqIndex++
				cursorx++
			})()

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

func (me *TextDisplay) SetFont(ts *TileSet) *TextDisplay {
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

func (me *TextDisplay) SetAlpha(a CanvasAlpha) *TextDisplay {
	me.mp.SetAlpha(a)
	return me
}

// ==============================================================================
// Actions
// ==============================================================================
func (me *TextDisplay) ToCanvas(ca *Canvas) {
	me.mp.ToCanvas(ca)
}
