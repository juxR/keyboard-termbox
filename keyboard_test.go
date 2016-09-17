// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package termbox

import (
	"fmt"
	"github.com/jteeuwen/keyboard"
	term "github.com/nsf/termbox-go"
	"testing"
)

func Test(t *testing.T) {
	err := term.Init()
	if err != nil {
		t.Fatal(err)
	}

	defer term.Close()

	running := true
	var arr [6]bool

	kb := New()
	kb.Bind(func() {
		arr[0] = !arr[0]
		drawList(kb, arr[:])
	}, "s")

	kb.Bind(func() {
		arr[1] = !arr[1]
		arr[2] = !arr[2]
		drawList(kb, arr[:])
	}, "ctrl+s", "command+s")

	kb.Bind(func() {
		arr[3] = !arr[3]
		drawList(kb, arr[:])
	}, "t e s t")

	kb.Bind(func() {
		arr[4] = !arr[4]
		drawList(kb, arr[:])
	}, "shift+a b c")

	kb.Bind(func() {
		running = false
	}, "escape")

	drawList(kb, arr[:])

	for running {
		kb.Poll(term.PollEvent())
	}
}

func drawList(kb keyboard.Keyboard, arr []bool) {
	term.Clear(term.ColorWhite, term.ColorBlack)
	printf(1, 1, term.ColorYellow, "Known Shortcuts:")

	for i, b := range kb.Bindings() {
		if arr[i] {
			printf(3, i+3, term.ColorBlue|term.AttrBold, "- %q", b)
		} else {
			printf(3, i+3, term.ColorWhite, "- %q", b)
		}
	}

	term.Flush()
}

func printf(x, y int, fg term.Attribute, f string, argv ...interface{}) {
	bg := term.ColorBlack
	f = fmt.Sprintf(f, argv...)

	for i, r := range f {
		term.SetCell(x+i, y, r, fg, bg)
	}
}
