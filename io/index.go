package io

import "syscall/js"

func onIOMessage(this js.Value, args []js.Value) interface{} {
	ev := args[0].Get("data")
	if (ev.Type() != js.Undefined().Type()) && (ev.Get("0").String() == "vblankdone") {

		data := ev.Get("2")
		if data.Get("length").Int() != 11 {
			js.Global().Get("console").Call("warn", "mousestate has not the required length expect 11 got", data.Get("length"))
		}

		mouseState.X = uint16(data.Get("0").Int())
		mouseState.Y = uint16(data.Get("1").Int())
		mouseState.Buttons = MouseButton(data.Get("2").Int())

		keyboardState[0] = uint32(data.Get("3").Int())
		keyboardState[1] = uint32(data.Get("4").Int())
		keyboardState[2] = uint32(data.Get("5").Int())
		keyboardState[3] = uint32(data.Get("6").Int())
		keyboardState[4] = uint32(data.Get("7").Int())
		keyboardState[5] = uint32(data.Get("8").Int())
		keyboardState[6] = uint32(data.Get("9").Int())
		keyboardState[7] = uint32(data.Get("10").Int())

	}
	return nil
}

func init() {
	js.Global().Call("addEventListener", "message", js.FuncOf(onIOMessage), false)
}
