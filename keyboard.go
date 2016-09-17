// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package termbox

import (
	"github.com/jteeuwen/keyboard"
	term "github.com/nsf/termbox-go"
)

// Keyboard is a keyboard backend for termbox-go.
type Keyboard struct {
	*keyboard.Base
}

// New constructs a new keyboard instance.
func New() keyboard.Keyboard {
	kb := new(Keyboard)
	kb.Base = keyboard.NewBase()
	return kb
}

// Poll should be called in the termbox event loop.
// The argument should be a valid termbox event object.
func (kb *Keyboard) Poll(data interface{}) {
	evt, ok := data.(term.Event)

	if !ok || evt.Type != term.EventKey {
		return // Not meant for us.
	}

	var key keyboard.Key
	var mod keyboard.Modifier

	if evt.Mod == term.ModAlt {
		mod = keyboard.ModAlt
	}

	if evt.Ch > 0 {
		if evt.Ch >= 'A' && evt.Ch <= 'Z' {
			mod |= keyboard.ModShift
		}

		switch evt.Ch {
		case 'a', 'A':
			key = keyboard.KeyA
		case 'b', 'B':
			key = keyboard.KeyB
		case 'c', 'C':
			key = keyboard.KeyC
		case 'd', 'D':
			key = keyboard.KeyD
		case 'e', 'E':
			key = keyboard.KeyE
		case 'f', 'F':
			key = keyboard.KeyF
		case 'g', 'G':
			key = keyboard.KeyG
		case 'h', 'H':
			key = keyboard.KeyH
		case 'i', 'I':
			key = keyboard.KeyI
		case 'j', 'J':
			key = keyboard.KeyJ
		case 'k', 'K':
			key = keyboard.KeyK
		case 'l', 'L':
			key = keyboard.KeyL
		case 'm', 'M':
			key = keyboard.KeyM
		case 'n', 'N':
			key = keyboard.KeyN
		case 'o', 'O':
			key = keyboard.KeyO
		case 'p', 'P':
			key = keyboard.KeyP
		case 'q', 'Q':
			key = keyboard.KeyQ
		case 'r', 'R':
			key = keyboard.KeyR
		case 's', 'S':
			key = keyboard.KeyS
		case 't', 'T':
			key = keyboard.KeyT
		case 'u', 'U':
			key = keyboard.KeyU
		case 'v', 'V':
			key = keyboard.KeyV
		case 'w', 'W':
			key = keyboard.KeyW
		case 'x', 'X':
			key = keyboard.KeyX
		case 'y', 'Y':
			key = keyboard.KeyY
		case 'z', 'Z':
			key = keyboard.KeyZ
		case '0':
			key = keyboard.Key0
		case ')':
			mod |= keyboard.ModShift
			key = keyboard.Key0
		case '1':
			key = keyboard.Key1
		case '!':
			mod |= keyboard.ModShift
			key = keyboard.Key1
		case '2':
			key = keyboard.Key2
		case '@':
			mod |= keyboard.ModShift
			key = keyboard.Key2
		case '3':
			key = keyboard.Key3
		case '#':
			mod |= keyboard.ModShift
			key = keyboard.Key3
		case '4':
			key = keyboard.Key4
		case '$':
			mod |= keyboard.ModShift
			key = keyboard.Key4
		case '5':
			key = keyboard.Key5
		case '%':
			mod |= keyboard.ModShift
			key = keyboard.Key5
		case '6':
			key = keyboard.Key6
		case '^':
			mod |= keyboard.ModShift
			key = keyboard.Key6
		case '7':
			key = keyboard.Key7
		case '&':
			mod |= keyboard.ModShift
			key = keyboard.Key7
		case '8':
		case '*':
			mod |= keyboard.ModShift
			key = keyboard.Key8
		case '9':
		case '(':
			mod |= keyboard.ModShift
			key = keyboard.Key9
		case '`':
			key = keyboard.KeyGraveAccent
		case '~':
			mod |= keyboard.ModShift
			key = keyboard.KeyGraveAccent
		case ',':
			key = keyboard.KeyComma
		case '<':
			mod |= keyboard.ModShift
			key = keyboard.KeyComma
		case '.':
			key = keyboard.KeyPeriod
		case '>':
			mod |= keyboard.ModShift
			key = keyboard.KeyPeriod
		case '/':
			key = keyboard.KeySlash
		case '?':
			mod |= keyboard.ModShift
			key = keyboard.KeySlash
		case ';':
			key = keyboard.KeySemicolon
		case ':':
			mod |= keyboard.ModShift
			key = keyboard.KeySemicolon
		case '\'':
			key = keyboard.KeyApostrophe
		case '"':
			mod |= keyboard.ModShift
			key = keyboard.KeyApostrophe
		case '-':
			key = keyboard.KeyMinus
		case '_':
			mod |= keyboard.ModShift
			key = keyboard.KeyMinus
		case '=':
			key = keyboard.KeyEqual
		case '+':
			mod |= keyboard.ModShift
			key = keyboard.KeyEqual
		case '[':
			key = keyboard.KeyLeftBracket
		case '{':
			mod |= keyboard.ModShift
			key = keyboard.KeyLeftBracket
		case ']':
			key = keyboard.KeyRightBracket
		case '}':
			mod |= keyboard.ModShift
			key = keyboard.KeyRightBracket
		case '\\':
			key = keyboard.KeyBackslash
		case '|':
			mod |= keyboard.ModShift
			key = keyboard.KeyBackslash
		}

		kb.RecordKey(key, mod)
		return
	}
	switch evt.Key {
	case term.KeyBackspace:
		key = keyboard.KeyBackspace
	case term.KeyTab:
		key = keyboard.KeyTab
	case term.KeyEsc:
		key = keyboard.KeyEscape
	case term.KeySpace:
		key = keyboard.KeySpace
	case term.KeyF1:
		key = keyboard.KeyF1
	case term.KeyF2:
		key = keyboard.KeyF2
	case term.KeyF3:
		key = keyboard.KeyF3
	case term.KeyF4:
		key = keyboard.KeyF4
	case term.KeyF5:
		key = keyboard.KeyF5
	case term.KeyF6:
		key = keyboard.KeyF6
	case term.KeyF7:
		key = keyboard.KeyF7
	case term.KeyF8:
		key = keyboard.KeyF8
	case term.KeyF9:
		key = keyboard.KeyF9
	case term.KeyF10:
		key = keyboard.KeyF10
	case term.KeyF11:
		key = keyboard.KeyF11
	case term.KeyF12:
		key = keyboard.KeyF12
	case term.KeyInsert:
		key = keyboard.KeyInsert
	case term.KeyDelete:
		key = keyboard.KeyDelete
	case term.KeyHome:
		key = keyboard.KeyHome
	case term.KeyEnd:
		key = keyboard.KeyEnd
	case term.KeyPgup:
		key = keyboard.KeyPageUp
	case term.KeyPgdn:
		key = keyboard.KeyPageDown
	case term.KeyArrowUp:
		key = keyboard.KeyUp
	case term.KeyArrowDown:
		key = keyboard.KeyDown
	case term.KeyArrowLeft:
		key = keyboard.KeyLeft
	case term.KeyArrowRight:
		key = keyboard.KeyRight
	case term.KeyCtrlTilde:
		key = keyboard.KeyGraveAccent
		mod = keyboard.ModCtrl
	case term.KeyCtrlA:
		key = keyboard.KeyA
		mod = keyboard.ModCtrl
	case term.KeyCtrlB:
		key = keyboard.KeyB
		mod = keyboard.ModCtrl
	case term.KeyCtrlC:
		key = keyboard.KeyC
		mod = keyboard.ModCtrl
	case term.KeyCtrlD:
		key = keyboard.KeyD
		mod = keyboard.ModCtrl
	case term.KeyCtrlE:
		key = keyboard.KeyE
		mod = keyboard.ModCtrl
	case term.KeyCtrlF:
		key = keyboard.KeyF
		mod = keyboard.ModCtrl
	case term.KeyCtrlG:
		key = keyboard.KeyG
		mod = keyboard.ModCtrl
	case term.KeyCtrlJ:
		key = keyboard.KeyJ
		mod = keyboard.ModCtrl
	case term.KeyCtrlK:
		key = keyboard.KeyK
		mod = keyboard.ModCtrl
	case term.KeyCtrlL:
		key = keyboard.KeyL
		mod = keyboard.ModCtrl
	case term.KeyCtrlM:
		kb.RecordKey(keyboard.KeyEnter, mod)
		key = keyboard.KeyM
		mod = keyboard.ModCtrl
	case term.KeyCtrlN:
		key = keyboard.KeyN
		mod = keyboard.ModCtrl
	case term.KeyCtrlO:
		key = keyboard.KeyO
		mod = keyboard.ModCtrl
	case term.KeyCtrlP:
		key = keyboard.KeyP
		mod = keyboard.ModCtrl
	case term.KeyCtrlQ:
		key = keyboard.KeyQ
		mod = keyboard.ModCtrl
	case term.KeyCtrlR:
		key = keyboard.KeyR
		mod = keyboard.ModCtrl
	case term.KeyCtrlS:
		key = keyboard.KeyS
		mod = keyboard.ModCtrl
	case term.KeyCtrlT:
		key = keyboard.KeyT
		mod = keyboard.ModCtrl
	case term.KeyCtrlU:
		key = keyboard.KeyU
		mod = keyboard.ModCtrl
	case term.KeyCtrlV:
		key = keyboard.KeyV
		mod = keyboard.ModCtrl
	case term.KeyCtrlW:
		key = keyboard.KeyW
		mod = keyboard.ModCtrl
	case term.KeyCtrlX:
		key = keyboard.KeyX
		mod = keyboard.ModCtrl
	case term.KeyCtrlY:
		key = keyboard.KeyY
		mod = keyboard.ModCtrl
	case term.KeyCtrlZ:
		key = keyboard.KeyZ
		mod = keyboard.ModCtrl
	case term.KeyCtrlBackslash:
		key = keyboard.KeyBackslash
		mod = keyboard.ModCtrl
	case term.KeyCtrlRsqBracket:
		key = keyboard.KeyRightBracket
		mod = keyboard.ModCtrl
	case term.KeyCtrl6:
		key = keyboard.Key6
		mod = keyboard.ModCtrl
	case term.KeyCtrlUnderscore:
		key = keyboard.KeyMinus
		mod = keyboard.ModCtrl
	case term.KeyCtrl8:
		key = keyboard.Key8
		mod = keyboard.ModCtrl
	default:
		key = keyboard.KeyUnknown
	}

	kb.RecordKey(key, mod)
}
