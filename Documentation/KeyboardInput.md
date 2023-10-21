# Keyboard-Input

The Keyboard-Input is bound to the [Engine-State](./reference/EngineState.md) as
the `Keyboard` property. Thus it is provided to your [Scene](./Scenes.md)
through the [Engine-Lifecycle](./EngineLifecycle.md).

Here is a list of things, it can do:

<!-- vim-markdown-toc GFM -->

* [Checking specific Keys](#checking-specific-keys)
    * [Let's give an example.](#lets-give-an-example)
* [Checking key sequences](#checking-key-sequences)
    * [Key-History](#key-history)
    * [Rune-History](#rune-history)
        * [CLI-Like example](#cli-like-example)
    * [Clearing the History](#clearing-the-history)
* [Key - List](#key---list)

<!-- vim-markdown-toc -->

## Checking specific Keys

[Engine-State's](./reference/EngineState.md) `Keyboard` property comes with the
following properties:

```go
type KeyboardKey byte

type KeyboardState struct {
    Pressed       [KeyboardKey]bool
    Held          [KeyboardKey]bool
    Released      [KeyboardKey]bool
    PressedOrHeld [KeyboardKey]bool
    // ...
}
```

Each of these properties holds a Mapping between `KeyboardKey` and a boolean
value defining if a Key has the checked State.

These States are:

| State           | Description                                                                    |
| --------------- | ------------------------------------------------------------------------------ |
| `Pressed`       | `true` If the key has been freshly pressed this cycle but was not last cycle   |
| `Held`          | `true` If the key was pressed last cycle and is still pressed this cycle       |
| `Released`      | `true` If the Key was pressed/held last cycle but is no longer held this cycle |
| `PressedOrHeld` | `true` If the key was just pressed this cycle or is already held               |

what keys you can check for their States, you can see in the
[Key-List](#keylist) further down.

to check if a specific key has a specific state, you can use the following
method/code:

```go
func (me *demoScene) Tick(e *core.EngineState) bool {


    if e.Keyboard.Pressed[io.KEY_ESC] {
        // Do something if the Escape key was just pressed
        // but don't execute this as long is escape is held or when it is released
        // ...
    }

    if e.Keyboard.PressedOrHeld[io.KEY_ESC] {
        // Do something if the Escape key was just Pressed and while it is held down
        // ...
    }

    if e.Keyboard.Held[io.KEY_ESC] {
        // Do something if the Escape key was held for at least 2 cycles / frames
        // ...
    }

    if e.Keyboard.Released[io.KEY_ESC] {
        // Do something if the Escape key was released
        // ...
    }

}
```

### Let's give an example.

Many 3rd and 1st person, keyboard based games implement a "Sneak" key. A key
that makes them move slower while it is held. This is how you can implement that
in GoWas

```go
func (me *demoScene) Tick(e *core.EngineState) bool {


}
```

A much easier way would be to just check each Cycle/Tick if the Key is still
Held

```go
import GoWasIO 'github.com/rocco-gossmann/GoWas/io'

var sneakKey = GoWasIO.KEY_LSHIFT
var sneakModeActive = false

func (me *demoScene) Tick(e *core.EngineState) bool {

    sneakModeActive =  e.Keyboard.PressedOrHeld[sneakKey];

    if e.Keyboard.Pressed[sneakKey] {
        // Do something here, that only needs to be executed,
        // when entering "sneak mode"
    }

    if e.Keyboard.Released[sneakKey] {
        // Do something here, that only needs to be executed,
        // when leaving "sneak mode"
    }
}
```

## Checking key sequences

In addition to the current KeyStates the
[Engine-State's](./reference/EngineState.md) `Keyboard` property also provides
access to a history of the last 64 Keys that have been entered.

> [!notice]\
> "entered" meaning pressed and released

you can check the history in the following via the following functions.

```go
func (*KeyboardState) History(limit byte) []KeyboardKey
// and
func (*KeyboardState) HistoryRunes(limit byte) []rune
```

The `limit` parameter in both of these means `the last X history entries` where
`X` means how many entries.

> [!warning]\
> These histories should **NOT** be used if you need to do Processing of
> **complex string**, due to the histories lack of Shift or Alt/AltGr Layer
> recognition. (Modifier Keys are simply just another key in the History list)
>
> These histories are more intended for a simple CLI like usage for now.

### Key-History

The last 64 released keys are keept in an list, managed by the Engine.

```go
func (me *demoScene) Tick(e *core.EngineState) bool {
    var last3PressedKeys []rune := e.Keyboard.History(3)
    if len(last3PressedKeys) >= 3 {
        // Do something if at least 3 keys have been pressed and released
        // ...
    }
}
```

here is a more complex example:

```go
import 'github.com/rocco-gossmann/GoWas/io'

func (me *demoScene) Tick(e *core.EngineState) bool {
    // Get the last 2 keys
    var keys []rune = e.Keyboard.History(2)

    // If there have been 2 keys pressed and released
    if listLen == 2 {

        // If the last key was Enter/Return
        if(keys[1] == io.KEY_ENTER) {
            switch key[0] {
                // Do something based on what key was pressed before Enter/Return
                // ...
            }

            // Reset the history to wait until a key followed by Enter/Return is pressed again
            e.Keyboard.HistoryClear();
        }
    }
}
```

### Rune-History

Some keys have a Human-readable representation. In Go, these are called `rune`s.
If a key is released, that has a `Rune` assigned, it will be put into the
Engines `Rune-History`

The last 64 runes are keept in an list, managed by the Engine.

(See the [Key - List](#key---list) below, to learn, what key has what rune
attached.)

it works exactly as the [Key-History](#key-history), but uses runes, which opens
this possiblity of a cli like interface for text adventure games for example.

#### CLI-Like example
```go

import "github.com/rocco-gossmann/GoWas/io"
import "fmt"

func (me *demoScene) Tick(e *core.EngineState) bool {

    // If the enter key was released
    if io.Keyboard.Released[io.KEY_ENTER] {

        command := string(io.KeyboardHistoryRunes(20))
        switch command {
            case "turn left": fmt.PrintLn("you turned left")
            case "turn right": fmt.PrintLn("you turned right")
            case "look around": fmt.PrintLn("you see nothing")
            default: fmt.PrintLn("Could not read your command")
        }

        // Clear the History to wait for the next command
        e.Keyboard.HistoryClear();
    }

}

```

### Clearing the History



## Key - List

| Key                | Value | Rune  | Description             |
| ------------------ | ----- | ----- | ----------------------- |
| `io.KEY_BACKSPACE` | 8     |       |                         |
| `io.KEY_TAB`       | 9     |       |                         |
| `io.KEY_ENTER`     | 13    |       |                         |
| `io.KEY_LSHIFT`    | 16    |       |                         |
| `io.KEY_LCTRL`     | 17    |       |                         |
| `io.KEY_LALT`      | 18    |       |                         |
| `io.KEY_PAUSE`     | 19    |       |                         |
| `io.KEY_CAPSLOCK`  | 20    |       |                         |
| `io.KEY_ESC`       | 27    |       |                         |
| `io.KEY_SPACE`     | 32    | `" "` |                         |
|                    |       |       |                         |
| `io.KEY_PGUP`      | 33    |       | Page Up                 |
| `io.KEY_PGDOWN`    | 34    |       | Page DOwn               |
| `io.KEY_END`       | 35    |       |                         |
| `io.KEY_HOME`      | 36    |       |                         |
|                    |       |       |                         |
| `io.KEY_LEFT`      | 37    |       | Cursor Left             |
| `io.KEY_UP`        | 38    |       | Cursor Up               |
| `io.KEY_RIGHT`     | 39    |       |                         |
| `io.KEY_DOWN`      | 40    |       |                         |
|                    |       |       |                         |
| `io.KEY_PRINT`     | 44    |       |                         |
| `io.KEY_INSERT`    | 45    |       |                         |
| `io.KEY_DEL`       | 46    |       |                         |
|                    |       |       |                         |
| `io.KEY_0`         | 48    | `0`   | Number Keys             |
| `io.KEY_1`         | 49    | `1`   |                         |
| `io.KEY_2`         | 50    | `2`   |                         |
| `io.KEY_3`         | 51    | `3`   |                         |
| `io.KEY_4`         | 52    | `4`   |                         |
| `io.KEY_5`         | 53    | `5`   |                         |
| `io.KEY_6`         | 54    | `6`   |                         |
| `io.KEY_7`         | 55    | `7`   |                         |
| `io.KEY_8`         | 56    | `8`   |                         |
| `io.KEY_9`         | 57    | `9`   |                         |
|                    |       |       |                         |
| `io.KEY_A`         | 65    | `a`   | Alphabetic Keys         |
| `io.KEY_B`         | 66    | `b`   |                         |
| `io.KEY_C`         | 67    | `c`   |                         |
| `io.KEY_D`         | 68    | `d`   |                         |
| `io.KEY_E`         | 69    | `e`   |                         |
| `io.KEY_F`         | 70    | `f`   |                         |
| `io.KEY_G`         | 71    | `g`   |                         |
| `io.KEY_H`         | 72    | `h`   |                         |
| `io.KEY_I`         | 73    | `i`   |                         |
| `io.KEY_J`         | 74    | `j`   |                         |
| `io.KEY_K`         | 75    | `k`   |                         |
| `io.KEY_L`         | 76    | `l`   |                         |
| `io.KEY_M`         | 77    | `m`   |                         |
| `io.KEY_N`         | 78    | `n`   |                         |
| `io.KEY_O`         | 79    | `o`   |                         |
| `io.KEY_P`         | 80    | `p`   |                         |
| `io.KEY_Q`         | 81    | `q`   |                         |
| `io.KEY_R`         | 82    | `r`   |                         |
| `io.KEY_S`         | 83    | `s`   |                         |
| `io.KEY_T`         | 84    | `t`   |                         |
| `io.KEY_U`         | 85    | `u`   |                         |
| `io.KEY_V`         | 86    | `v`   |                         |
| `io.KEY_W`         | 87    | `w`   |                         |
| `io.KEY_X`         | 88    | `x`   |                         |
| `io.KEY_Y`         | 89    | `y`   |                         |
| `io.KEY_Z`         | 90    | `z`   |                         |
|                    |       |       |                         |
| `io.KEY_SUPER`     | 91    |       | Windows \| Cmd \| Super |
| `io.KEY_APP`       | 93    |       | Application / Menu      |
|                    |       |       |                         |
| `io.KEY_NUM0`      | 96    | `0`   | Numpad                  |
| `io.KEY_NUM1`      | 97    | `1`   |                         |
| `io.KEY_NUM2`      | 98    | `2`   |                         |
| `io.KEY_NUM3`      | 99    | `3`   |                         |
| `io.KEY_NUM4`      | 100   | `4`   |                         |
| `io.KEY_NUM5`      | 101   | `5`   |                         |
| `io.KEY_NUM6`      | 102   | `6`   |                         |
| `io.KEY_NUM7`      | 103   | `7`   |                         |
| `io.KEY_NUM8`      | 104   | `8`   |                         |
| `io.KEY_NUM9`      | 105   | `9`   |                         |
| `io.KEY_NUMMULT`   | 106   | `*`   |                         |
| `io.KEY_NUMADD`    | 107   | `+`   |                         |
| `io.KEY_NUMSUBST`  | 109   | `-`   |                         |
| `io.KEY_NUMDEC`    | 110   | `.`   |                         |
| `io.KEY_NUMDIV`    | 111   | `/`   |                         |
|                    |       |       |                         |
| `io.KEY_F1`        | 112   |       | Function                |
| `io.KEY_F2`        | 113   |       |                         |
| `io.KEY_F3`        | 114   |       |                         |
| `io.KEY_F4`        | 115   |       |                         |
| `io.KEY_F5`        | 116   |       |                         |
| `io.KEY_F6`        | 117   |       |                         |
| `io.KEY_F7`        | 118   |       |                         |
| `io.KEY_F8`        | 119   |       |                         |
| `io.KEY_F9`        | 120   |       |                         |
| `io.KEY_F10`       | 121   |       |                         |
| `io.KEY_F11`       | 122   |       |                         |
| `io.KEY_F12`       | 123   |       |                         |
| `io.KEY_F13`       | 124   |       |                         |
| `io.KEY_F14`       | 125   |       |                         |
| `io.KEY_F15`       | 126   |       |                         |
| `io.KEY_F16`       | 127   |       |                         |
| `io.KEY_F17`       | 128   |       |                         |
| `io.KEY_F18`       | 129   |       |                         |
| `io.KEY_F19`       | 130   |       |                         |
| `io.KEY_F20`       | 131   |       |                         |
| `io.KEY_F21`       | 132   |       |                         |
| `io.KEY_F22`       | 133   |       |                         |
| `io.KEY_F23`       | 134   |       |                         |
| `io.KEY_F24`       | 135   |       |                         |
|                    |       |       |                         |
| `io.KEY_NUMLOCK`   | 144   |       |                         |
| `io.KEY_SCRLOCK`   | 145   |       |                         |
|                    |       |       | Symbols                 |
| `io.KEY_SEMICOLON` | 186   | `;`   | ;                       |
| `io.KEY_EQUAL`     | 187   | `=`   | =                       |
|                    |       |       |                         |
| `io.KEY_COMMA`     | 188   | `,`   | ,                       |
| `io.KEY_DASH`      | 189   | `-`   | -                       |
| `io.KEY_MINUS`     | 189   | `-`   |                         |
| `io.KEY_PERIOD`    | 190   | `.`   | .                       |
| `io.KEY_SLASH`     | 191   | `/`   | /                       |
| `io.KEY_BACKQUOTE` | 192   |       |                         |
| `io.KEY_BRKOPEN`   | 219   | `[`   | [                       |
| `io.KEY_BACKSLASH` | 220   | `\`   | \                       |
| `io.KEY_BRKCLOSE`  | 221   | `]`   | ]                       |
| `io.KEY_QUOTE`     | 222   | `"`   | "                       |
