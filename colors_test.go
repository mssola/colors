// Copyright (C) 2014-2015 Miquel Sabaté Solà <mikisabate@gmail.com>
// This file is licensed under the MIT license.
// See the LICENSE file.

package colors

import "testing"

var colors = []struct {
	color    Colors
	bg       bool
	mode     Mode
	expected string
}{
	// Saved
	{Saved, false, Regular, "\x1b[0;49;39mHello\x1b[0;m"},
	{Saved, false, Bold, "\x1b[1;49;39mHello\x1b[0;m"},
	{Saved, true, Bold, "\x1b[1;49;39mHello\x1b[0;m"},
	{Saved, true, Underlined, "\x1b[4;49;39mHello\x1b[0;m"},
	{Saved, true, Reversed, "\x1b[7;49;39mHello\x1b[0;m"},

	// Black
	{Black, false, Regular, "\x1b[0;49;30mHello\x1b[0;m"},
	{Black, false, Bold, "\x1b[1;49;30mHello\x1b[0;m"},
	{Black, true, Bold, "\x1b[1;40;30mHello\x1b[0;m"},
	{Black, true, Underlined, "\x1b[4;40;30mHello\x1b[0;m"},
	{Black, true, Reversed, "\x1b[7;40;30mHello\x1b[0;m"},

	// Red
	{Red, false, Regular, "\x1b[0;40;31mHello\x1b[0;m"},
	{Red, false, Bold, "\x1b[1;40;31mHello\x1b[0;m"},
	{Red, true, Bold, "\x1b[1;41;31mHello\x1b[0;m"},
	{Red, true, Underlined, "\x1b[4;41;31mHello\x1b[0;m"},
	{Red, true, Reversed, "\x1b[7;41;31mHello\x1b[0;m"},

	// Green
	{Green, false, Regular, "\x1b[0;41;32mHello\x1b[0;m"},
	{Green, false, Bold, "\x1b[1;41;32mHello\x1b[0;m"},
	{Green, true, Bold, "\x1b[1;42;32mHello\x1b[0;m"},
	{Green, true, Underlined, "\x1b[4;42;32mHello\x1b[0;m"},
	{Green, true, Reversed, "\x1b[7;42;32mHello\x1b[0;m"},

	// Yellow
	{Yellow, false, Regular, "\x1b[0;42;33mHello\x1b[0;m"},
	{Yellow, false, Bold, "\x1b[1;42;33mHello\x1b[0;m"},
	{Yellow, true, Bold, "\x1b[1;43;33mHello\x1b[0;m"},
	{Yellow, true, Underlined, "\x1b[4;43;33mHello\x1b[0;m"},
	{Yellow, true, Reversed, "\x1b[7;43;33mHello\x1b[0;m"},

	// Blue
	{Blue, false, Regular, "\x1b[0;43;34mHello\x1b[0;m"},
	{Blue, false, Bold, "\x1b[1;43;34mHello\x1b[0;m"},
	{Blue, true, Bold, "\x1b[1;44;34mHello\x1b[0;m"},
	{Blue, true, Underlined, "\x1b[4;44;34mHello\x1b[0;m"},
	{Blue, true, Reversed, "\x1b[7;44;34mHello\x1b[0;m"},

	// Magenta
	{Magenta, false, Regular, "\x1b[0;44;35mHello\x1b[0;m"},
	{Magenta, false, Bold, "\x1b[1;44;35mHello\x1b[0;m"},
	{Magenta, true, Bold, "\x1b[1;45;35mHello\x1b[0;m"},
	{Magenta, true, Underlined, "\x1b[4;45;35mHello\x1b[0;m"},
	{Magenta, true, Reversed, "\x1b[7;45;35mHello\x1b[0;m"},

	// Cyan
	{Cyan, false, Regular, "\x1b[0;45;36mHello\x1b[0;m"},
	{Cyan, false, Bold, "\x1b[1;45;36mHello\x1b[0;m"},
	{Cyan, true, Bold, "\x1b[1;46;36mHello\x1b[0;m"},
	{Cyan, true, Underlined, "\x1b[4;46;36mHello\x1b[0;m"},
	{Cyan, true, Reversed, "\x1b[7;46;36mHello\x1b[0;m"},

	// White
	{White, false, Regular, "\x1b[0;46;37mHello\x1b[0;m"},
	{White, false, Bold, "\x1b[1;46;37mHello\x1b[0;m"},
	{White, true, Bold, "\x1b[1;47;37mHello\x1b[0;m"},
	{White, true, Underlined, "\x1b[4;47;37mHello\x1b[0;m"},
	{White, true, Reversed, "\x1b[7;47;37mHello\x1b[0;m"},
}

func TestColors(t *testing.T) {
	// Default
	color := Default()
	if color.Foreground != Saved {
		t.Errorf("Foreground should be Saved!")
	}
	if color.Background != Saved {
		t.Errorf("Background should be Saved!")
	}
	if color.Mode != Regular {
		t.Errorf("Mode should be Regular!")
	}

	// Different options
	for _, c := range colors {
		color.Change(c.color, c.bg)
		color.SetMode(c.mode)
		got := color.Get("Hello")
		if c.expected != got {
			t.Errorf("Expected: %v, Got: %v\n", c.expected, got)
		}
	}
}
