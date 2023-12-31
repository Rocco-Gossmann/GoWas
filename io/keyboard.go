package io

const keyboardHistoryLimit = 64

type KeyboardKey byte

const (
	KEY_BACKSPACE KeyboardKey = 8
	KEY_TAB       KeyboardKey = 9
	KEY_ENTER     KeyboardKey = 13
	KEY_LSHIFT    KeyboardKey = 16
	KEY_LCTRL     KeyboardKey = 17
	KEY_LALT      KeyboardKey = 18
	KEY_PAUSE     KeyboardKey = 19
	KEY_CAPSLOCK  KeyboardKey = 20
	KEY_ESC       KeyboardKey = 27
	KEY_SPACE     KeyboardKey = 32

	KEY_PGUP   KeyboardKey = 33
	KEY_PGDOWN KeyboardKey = 34
	KEY_END    KeyboardKey = 35
	KEY_HOME   KeyboardKey = 36

	KEY_LEFT  KeyboardKey = 37
	KEY_UP    KeyboardKey = 38
	KEY_RIGHT KeyboardKey = 39
	KEY_DOWN  KeyboardKey = 40

	KEY_PRINT  KeyboardKey = 44
	KEY_INSERT KeyboardKey = 45
	KEY_DEL    KeyboardKey = 46

	KEY_0 KeyboardKey = 48
	KEY_1 KeyboardKey = 49
	KEY_2 KeyboardKey = 50
	KEY_3 KeyboardKey = 51
	KEY_4 KeyboardKey = 52
	KEY_5 KeyboardKey = 53
	KEY_6 KeyboardKey = 54
	KEY_7 KeyboardKey = 55
	KEY_8 KeyboardKey = 56
	KEY_9 KeyboardKey = 57

	KEY_A KeyboardKey = 65
	KEY_B KeyboardKey = 66
	KEY_C KeyboardKey = 67
	KEY_D KeyboardKey = 68
	KEY_E KeyboardKey = 69
	KEY_F KeyboardKey = 70
	KEY_G KeyboardKey = 71
	KEY_H KeyboardKey = 72
	KEY_I KeyboardKey = 73
	KEY_J KeyboardKey = 74
	KEY_K KeyboardKey = 75
	KEY_L KeyboardKey = 76
	KEY_M KeyboardKey = 77
	KEY_N KeyboardKey = 78
	KEY_O KeyboardKey = 79
	KEY_P KeyboardKey = 80
	KEY_Q KeyboardKey = 81
	KEY_R KeyboardKey = 82
	KEY_S KeyboardKey = 83
	KEY_T KeyboardKey = 84
	KEY_U KeyboardKey = 85
	KEY_V KeyboardKey = 86
	KEY_W KeyboardKey = 87
	KEY_X KeyboardKey = 88
	KEY_Y KeyboardKey = 89
	KEY_Z KeyboardKey = 90

	KEY_SUPER KeyboardKey = 91
	KEY_APP   KeyboardKey = 93

	KEY_NUM0     KeyboardKey = 96
	KEY_NUM1     KeyboardKey = 97
	KEY_NUM2     KeyboardKey = 98
	KEY_NUM3     KeyboardKey = 99
	KEY_NUM4     KeyboardKey = 100
	KEY_NUM5     KeyboardKey = 101
	KEY_NUM6     KeyboardKey = 102
	KEY_NUM7     KeyboardKey = 103
	KEY_NUM8     KeyboardKey = 104
	KEY_NUM9     KeyboardKey = 105
	KEY_NUMMULT  KeyboardKey = 106
	KEY_NUMADD   KeyboardKey = 107
	KEY_NUMSUBST KeyboardKey = 109
	KEY_NUMDEC   KeyboardKey = 110
	KEY_NUMDIV   KeyboardKey = 111

	KEY_F1  KeyboardKey = 112
	KEY_F2  KeyboardKey = 113
	KEY_F3  KeyboardKey = 114
	KEY_F4  KeyboardKey = 115
	KEY_F5  KeyboardKey = 116
	KEY_F6  KeyboardKey = 117
	KEY_F7  KeyboardKey = 118
	KEY_F8  KeyboardKey = 119
	KEY_F9  KeyboardKey = 120
	KEY_F10 KeyboardKey = 121
	KEY_F11 KeyboardKey = 122
	KEY_F12 KeyboardKey = 123
	KEY_F13 KeyboardKey = 124
	KEY_F14 KeyboardKey = 125
	KEY_F15 KeyboardKey = 126
	KEY_F16 KeyboardKey = 127
	KEY_F17 KeyboardKey = 128
	KEY_F18 KeyboardKey = 129
	KEY_F19 KeyboardKey = 130
	KEY_F20 KeyboardKey = 131
	KEY_F21 KeyboardKey = 132
	KEY_F22 KeyboardKey = 133
	KEY_F23 KeyboardKey = 134
	KEY_F24 KeyboardKey = 135

	KEY_NUMLOCK KeyboardKey = 144
	KEY_SCRLOCK KeyboardKey = 145

	KEY_SEMICOLON KeyboardKey = 186
	KEY_EQUAL     KeyboardKey = 187

	KEY_COMMA     KeyboardKey = 188
	KEY_DASH      KeyboardKey = 189
	KEY_MINUS     KeyboardKey = 189
	KEY_PERIOD    KeyboardKey = 190
	KEY_SLASH     KeyboardKey = 191
	KEY_BACKQUOTE KeyboardKey = 192
	KEY_BRKOPEN   KeyboardKey = 219
	KEY_BACKSLASH KeyboardKey = 220
	KEY_BRKCLOSE  KeyboardKey = 221
	KEY_QUOTE     KeyboardKey = 222
)

var charMap = make(map[KeyboardKey]rune, 256)

func init() {
	charMap[KEY_SPACE] = rune(' ')
	charMap[KEY_0] = rune('0')
	charMap[KEY_1] = rune('1')
	charMap[KEY_2] = rune('2')
	charMap[KEY_3] = rune('3')
	charMap[KEY_4] = rune('4')
	charMap[KEY_5] = rune('5')
	charMap[KEY_6] = rune('6')
	charMap[KEY_7] = rune('7')
	charMap[KEY_8] = rune('8')
	charMap[KEY_9] = rune('9')
	charMap[KEY_A] = rune('a')
	charMap[KEY_B] = rune('b')
	charMap[KEY_C] = rune('c')
	charMap[KEY_D] = rune('d')
	charMap[KEY_E] = rune('e')
	charMap[KEY_F] = rune('f')
	charMap[KEY_G] = rune('g')
	charMap[KEY_H] = rune('h')
	charMap[KEY_I] = rune('i')
	charMap[KEY_J] = rune('j')
	charMap[KEY_K] = rune('k')
	charMap[KEY_L] = rune('l')
	charMap[KEY_M] = rune('m')
	charMap[KEY_N] = rune('n')
	charMap[KEY_O] = rune('o')
	charMap[KEY_P] = rune('p')
	charMap[KEY_Q] = rune('q')
	charMap[KEY_R] = rune('r')
	charMap[KEY_S] = rune('s')
	charMap[KEY_T] = rune('t')
	charMap[KEY_U] = rune('u')
	charMap[KEY_V] = rune('v')
	charMap[KEY_W] = rune('w')
	charMap[KEY_X] = rune('x')
	charMap[KEY_Y] = rune('y')
	charMap[KEY_Z] = rune('z')
	charMap[KEY_MINUS] = rune('-')
	charMap[KEY_BRKOPEN] = rune('[')
	charMap[KEY_BRKCLOSE] = rune(']')
	charMap[KEY_EQUAL] = rune('+')
	charMap[KEY_COMMA] = rune(',')
	charMap[KEY_SLASH] = rune('?')
	charMap[KEY_PERIOD] = rune('.')
	charMap[KEY_BACKSLASH] = rune('\\')
	charMap[KEY_SEMICOLON] = rune(';')
	charMap[KEY_DASH] = rune('-')
	charMap[KEY_QUOTE] = rune('"')
	charMap[KEY_EQUAL] = rune('=')
	charMap[KEY_NUMDIV] = rune('/')
	charMap[KEY_NUMDEC] = rune('.')
	charMap[KEY_NUMSUBST] = rune('-')
	charMap[KEY_NUMADD] = rune('+')
	charMap[KEY_NUMMULT] = rune('*')

	charMap[KEY_NUM0] = rune('0')
	charMap[KEY_NUM1] = rune('1')
	charMap[KEY_NUM2] = rune('2')
	charMap[KEY_NUM3] = rune('3')
	charMap[KEY_NUM4] = rune('4')
	charMap[KEY_NUM5] = rune('5')
	charMap[KEY_NUM6] = rune('6')
	charMap[KEY_NUM7] = rune('7')
	charMap[KEY_NUM8] = rune('8')
	charMap[KEY_NUM9] = rune('9')
}

type KeyboardState struct {
	Pressed       [256]bool
	Held          [256]bool
	Released      [256]bool
	PressedOrHeld [256]bool

	HistoryIndex uint

	historyStr    [keyboardHistoryLimit]rune
	historyStrPtr byte

	history    [keyboardHistoryLimit]KeyboardKey
	historyPtr byte
}

func normalizeLimit(limit byte) byte {
	return max(0, min(keyboardHistoryLimit, limit))
}

func (st *KeyboardState) History(limit byte) []KeyboardKey {
	limit = normalizeLimit(limit)
	l := limit
	var str = make([]KeyboardKey, l, l)

	var idx, offs byte

	for idx < limit {
		k := st.history[(keyboardHistoryLimit+(st.historyPtr-l))%keyboardHistoryLimit]
		if k == 0 {
			offs++
		} else {
			str[idx-offs] = k
		}

		l--
		idx++
	}

	return str
}

func (st *KeyboardState) HistoryRunes(limit byte) []rune {
	limit = normalizeLimit(limit)
	l := limit
	var str = make([]rune, l, l)

	var idx, offs byte

	for idx < limit {
		k := st.historyStr[(keyboardHistoryLimit+(st.historyStrPtr-l))%keyboardHistoryLimit]
		if k == 0 {
			offs++
		} else {
			str[idx-offs] = k
		}
		l--
		idx++
	}

	return str
}

func (st *KeyboardState) HistoryClear() *KeyboardState {
	for a := 0; a < keyboardHistoryLimit; a++ {
		st.history[a] = 0
		st.historyStr[a] = 0
	}
	st.historyPtr = 0
	st.historyStrPtr = 0
	st.HistoryIndex = 0

	return st
}

var keyboardState [8]uint32

var lastKeyboardState KeyboardState

func UpdateKeys(st *KeyboardState) {
	for a := 0; a < 8; a++ {
		for b := 0; b < 32; b++ {
			offset := uint32(1 << b)
			idx := a*32 + b
			bit := keyboardState[a]&offset == offset

			// Same steps as in UpdateMouse
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

	for a := 0; a < 256; a++ {
		key := KeyboardKey(a)
		chr, ok := charMap[key]
		if ok && chr != 0 && lastKeyboardState.Released[key] {
			lastKeyboardState.historyStr[lastKeyboardState.historyStrPtr] = chr
			lastKeyboardState.historyStrPtr = (lastKeyboardState.historyStrPtr + 1) % keyboardHistoryLimit
		}

		if lastKeyboardState.Released[a] {
			lastKeyboardState.history[lastKeyboardState.historyPtr] = KeyboardKey(a)
			lastKeyboardState.historyPtr = (lastKeyboardState.historyPtr + 1) % keyboardHistoryLimit
			lastKeyboardState.HistoryIndex++
		}
	}

	st.history = lastKeyboardState.history
	st.historyPtr = lastKeyboardState.historyPtr
	st.historyStr = lastKeyboardState.historyStr
	st.historyStrPtr = lastKeyboardState.historyStrPtr

}
