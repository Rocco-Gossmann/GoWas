package ui

import (
	"github.com/rocco-gossmann/GoWas/core"
)

func CreateLabel(textDisplay *core.TextDisplay, x, y int32, maxLength uint) *Label {
	label := Label{}
	label.length = maxLength
	if textDisplay == nil {
		panic("can't create label without a TextDisplay")
	}
	label.display = textDisplay
	label.position.X = x
	label.position.Y = y

	return &label
}
