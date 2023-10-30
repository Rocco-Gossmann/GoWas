package ui

import "github.com/rocco-gossmann/GoWas/core"

func CreateLabel(textDisplay *core.TextDisplay, length uint) *Label {
	label := Label{}
	label.maxLength = length
	label.textDisplay = textDisplay
	return &label
}
