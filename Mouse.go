package GoWas

import (
	"syscall/js"
)

type MouseState struct {
	X, Y    uint16
	Buttons uint32
}

// Holding the MouseState itself
var mouseState = MouseState{}

// Channel to responde to the other parts of the engine about the current MouseState
var mouseStateChannel chan MouseState = make(chan MouseState)

func onMouseMessage(this js.Value, args []js.Value) interface{} {
	ev := args[0].Get("data")
	if (ev.Type() != js.Undefined().Type()) && (ev.Get("0").String() == "mousestate") {

		data := ev.Get("1")
		if data.Get("length").Int() != 3 {
			js.Global().Get("console").Call("warn", "mousestate has not the required length expect 3 got", data.Get("length"))
		}

		mouseState.X = uint16(data.Get("0").Int())
		mouseState.Y = uint16(data.Get("1").Int())
		mouseState.Buttons = uint32(data.Get("2").Int())

		mouseStateChannel <- mouseState
	}
	return nil
}

func init() {
	js.Global().Call("addEventListener", "message", js.FuncOf(onMouseMessage), false)
}

func UpdateMouse() chan MouseState {
	js.Global().Call(
		"postMessage",                // command
		[]interface{}{"mouseupdate"}, // data
	)

	return mouseStateChannel
}
