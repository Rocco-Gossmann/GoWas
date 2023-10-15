package core

import (
	"syscall/js"
)

type MouseButton uint32

const (
	MOUSE_BTN1  MouseButton = 0x00000001
	MOUSE_BTN2  MouseButton = 0x00000002
	MOUSE_BTN3  MouseButton = 0x00000004
	MOUSE_BTN4  MouseButton = 0x00000008
	MOUSE_BTN5  MouseButton = 0x00000010
	MOUSE_BTN6  MouseButton = 0x00000020
	MOUSE_BTN7  MouseButton = 0x00000040
	MOUSE_BTN8  MouseButton = 0x00000080
	MOUSE_BTN9  MouseButton = 0x00000100
	MOUSE_BTN10 MouseButton = 0x00000200
	MOUSE_BTN11 MouseButton = 0x00000400
	MOUSE_BTN12 MouseButton = 0x00000800
	MOUSE_BTN13 MouseButton = 0x00001000
	MOUSE_BTN14 MouseButton = 0x00002000
	MOUSE_BTN15 MouseButton = 0x00004000
	MOUSE_BTN16 MouseButton = 0x00008000
	MOUSE_BTN17 MouseButton = 0x00010000
	MOUSE_BTN18 MouseButton = 0x00020000
	MOUSE_BTN19 MouseButton = 0x00040000
	MOUSE_BTN20 MouseButton = 0x00080000
	MOUSE_BTN21 MouseButton = 0x00100000
	MOUSE_BTN22 MouseButton = 0x00200000
	MOUSE_BTN23 MouseButton = 0x00400000
	MOUSE_BTN24 MouseButton = 0x00800000
	MOUSE_BTN25 MouseButton = 0x01000000
	MOUSE_BTN26 MouseButton = 0x02000000
	MOUSE_BTN27 MouseButton = 0x04000000
	MOUSE_BTN28 MouseButton = 0x08000000
	MOUSE_BTN29 MouseButton = 0x10000000
	MOUSE_BTN30 MouseButton = 0x20000000
	MOUSE_BTN31 MouseButton = 0x40000000
	MOUSE_BTN32 MouseButton = 0x80000000
)

type MouseState struct {
	X, Y          uint16
	Pressed       MouseButton
	Held          MouseButton
	Released      MouseButton
	PressedOrHeld MouseButton
}

const keyboardHistoryLimit = 64

type KeyboardState struct {
	Pressed       [128]bool
	Held          [128]bool
	Released      [128]bool
	PressedOrHeld [128]bool

	history    [keyboardHistoryLimit]byte
	historyPtr byte
}

func (st *KeyboardState) HistoryStr(limit byte) string {

	l := max(0, min(64, limit))
	var str = make([]rune, l, l)
	var idx byte
	var ptr byte = l - 1

	for idx < l {
		str[idx] = rune(st.history[(keyboardHistoryLimit+(st.historyPtr-ptr))%keyboardHistoryLimit])
		ptr--
		idx++
	}

	return string(str)
}

// Holding the MouseState itself
var mouseState struct {
	X, Y    uint16
	Buttons MouseButton // Held Buttons
}

var keyboardState [4]uint32

var lastMouseState MouseState
var lastKeyboardState KeyboardState

func onMouseMessage(this js.Value, args []js.Value) interface{} {
	ev := args[0].Get("data")
	if (ev.Type() != js.Undefined().Type()) && (ev.Get("0").String() == "vblankdone") {

		data := ev.Get("2")
		if data.Get("length").Int() != 7 {
			js.Global().Get("console").Call("warn", "mousestate has not the required length expect 7 got", data.Get("length"))
		}

		mouseState.X = uint16(data.Get("0").Int())
		mouseState.Y = uint16(data.Get("1").Int())
		mouseState.Buttons = MouseButton(data.Get("2").Int())

		keyboardState[0] = uint32(data.Get("3").Int())
		keyboardState[1] = uint32(data.Get("4").Int())
		keyboardState[2] = uint32(data.Get("5").Int())
		keyboardState[3] = uint32(data.Get("6").Int())

	}
	return nil
}

func init() {
	js.Global().Call("addEventListener", "message", js.FuncOf(onMouseMessage), false)
}

func UpdateKeys(st *KeyboardState) {
	for a := 0; a < 4; a++ {
		for b := 0; b < 32; b++ {
			offset := uint32(1 << b)
			idx := a*32 + b
			bit := keyboardState[a]&offset == offset

			lastKeyboardState.Held[idx] = lastKeyboardState.Held[idx] || lastKeyboardState.Pressed[idx]
			lastKeyboardState.Released[idx] = lastKeyboardState.Held[idx] && !bit
			lastKeyboardState.Pressed[idx] = !lastKeyboardState.Held[idx] && bit
			lastKeyboardState.Held[idx] = bit
			lastKeyboardState.PressedOrHeld[idx] = lastKeyboardState.Held[idx] || lastKeyboardState.Pressed[idx]

			st.Held[idx] = lastKeyboardState.Held[idx]
			st.Released[idx] = lastKeyboardState.Released[idx]
			st.Pressed[idx] = lastKeyboardState.Pressed[idx]
			st.PressedOrHeld[idx] = lastKeyboardState.PressedOrHeld[idx]
		}
	}

	for a := 0; a < 128; a++ {
		if lastKeyboardState.Pressed[a] {
			lastKeyboardState.history[lastKeyboardState.historyPtr] = byte(a)
			lastKeyboardState.historyPtr = (lastKeyboardState.historyPtr + 1) % keyboardHistoryLimit
		}
	}

	st.history = lastKeyboardState.history
	st.historyPtr = lastKeyboardState.historyPtr

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
