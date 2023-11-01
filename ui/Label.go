package ui

import (
	"github.com/rocco-gossmann/GoWas/core"
	"github.com/rocco-gossmann/GoWas/types"
)

type Label struct {
	display  *core.TextDisplay
	length   uint
	text     string
	position types.PointI32
}

func (me *Label) Text(txt string) {
	wrap := me.display.Wrap()
	x, y := me.display.Cursor()
	me.text = string([]rune(txt)[:me.length])

	me.display.
		SetCursor(me.position.X, me.position.Y).
		Clear(me.length).
		Echo(me.text).
		SetCursor(int32(x), int32(y)).
		SetWrap(wrap)
}

func (me *Label) Clear() *Label {
	wrap := me.display.Wrap()
	x, y := me.display.Cursor()

	me.display.
		SetCursor(int32(me.position.X), int32(me.position.Y)).
		Clear(me.length).
		SetCursor(int32(x), int32(y)).
		SetWrap(wrap)

	return me
}

func (me *Label) Position(x, y int32) *Label {

	wrap := me.display.Wrap()
	origx, origy := me.display.Cursor()

	me.display.
		SetWrap(false).
		SetCursor(me.position.X, me.position.Y).
		Clear(me.length).
		SetCursor(x, y).
		Clear(me.length).
		Echo(me.text).
		SetWrap(wrap).
		SetCursor(int32(origx), int32(origy))

	return me
}
