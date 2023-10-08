package core

import (
	"syscall/js"
)

type MouseButton uint32

const (
	MOUSE_BTN1 MouseButton = 0x00000001
	MOUSE_BTN2 MouseButton = 0x00000010
	MOUSE_BTN3 MouseButton = 0x00000100
	MOUSE_BTN4 MouseButton = 0x00001000
	MOUSE_BTN5 MouseButton = 0x00010000
	MOUSE_BTN6 MouseButton = 0x00100000
	MOUSE_BTN7 MouseButton = 0x01000000
	MOUSE_BTN8 MouseButton = 0x10000000
)

type MouseState struct {
	X, Y          uint16
	Pressed       uint32
	Held          uint32
	Released      uint32
	PressedOrHeld uint32
}

// Holding the MouseState itself
var mouseState struct {
	X, Y    uint16
	Buttons uint32 // Held Buttons
}

var lastMouseState MouseState

func onMouseMessage(this js.Value, args []js.Value) interface{} {
	ev := args[0].Get("data")
	if (ev.Type() != js.Undefined().Type()) && (ev.Get("0").String() == "vblankdone") {

		data := ev.Get("2")
		if data.Get("length").Int() != 3 {
			js.Global().Get("console").Call("warn", "mousestate has not the required length expect 3 got", data.Get("length"))
		}

		mouseState.X = uint16(data.Get("0").Int())
		mouseState.Y = uint16(data.Get("1").Int())
		mouseState.Buttons = uint32(data.Get("2").Int())

	}
	return nil
}

func init() {
	js.Global().Call("addEventListener", "message", js.FuncOf(onMouseMessage), false)
}

func UpdateMouse(st *MouseState) {
	var btns = mouseState.Buttons
	var ntbtns = ^mouseState.Buttons

	lastMouseState.X = mouseState.X
	lastMouseState.Y = mouseState.Y

	// move all last pressed keys, that are still pressed to held
	lastMouseState.Held |= lastMouseState.Pressed

	// Released buttons are the once no longer pressed or held
	lastMouseState.Released = lastMouseState.Held & ntbtns

	// All keys in the new state, but not in Held are the new Pressed
	lastMouseState.Pressed = (^lastMouseState.Held) & btns

	// Remove all no longer registered buttons from held
	lastMouseState.Held &= btns

	// Combine Pressed and held into one conviniently readable value too
	lastMouseState.PressedOrHeld = lastMouseState.Held | lastMouseState.Pressed

	st.Held = lastMouseState.Held
	st.Pressed = lastMouseState.Pressed
	st.Released = lastMouseState.Released
	st.PressedOrHeld = lastMouseState.PressedOrHeld
	st.X = lastMouseState.X
	st.Y = lastMouseState.Y
}
