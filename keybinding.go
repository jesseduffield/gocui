// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gocui

import (
	"github.com/gdamore/tcell/v2"
)

// Key represents special keys or keys combinations.
type Key tcell.Key

// Modifier allows to define special keys combinations. They can be used
// in combination with Keys or Runes when a new keybinding is defined.
type Modifier tcell.ModMask

// Keybindings are used to link a given key-press event with a handler.
type keybinding struct {
	viewName string
	key      Key
	ch       rune
	mod      Modifier
	handler  func(*Gui, *View) error
}

// newKeybinding returns a new Keybinding object.
func newKeybinding(viewname string, key Key, ch rune, mod Modifier, handler func(*Gui, *View) error) (kb *keybinding) {
	kb = &keybinding{
		viewName: viewname,
		key:      key,
		ch:       ch,
		mod:      mod,
		handler:  handler,
	}
	return kb
}

func eventMatchesKey(ev *GocuiEvent, key any) bool {
	// assuming ModNone for now
	if ev.Mod != ModNone {
		return false
	}

	k, ch, err := getKey(key)
	if err != nil {
		return false
	}

	return k == ev.Key && ch == ev.Ch
}

// matchKeypress returns if the keybinding matches the keypress.
func (kb *keybinding) matchKeypress(key Key, ch rune, mod Modifier) bool {
	return kb.key == key && kb.ch == ch && kb.mod == mod
}

// Special keys.
const (
	KeyF1             Key = Key(tcell.KeyF1)
	KeyF2                 = Key(tcell.KeyF2)
	KeyF3                 = Key(tcell.KeyF3)
	KeyF4                 = Key(tcell.KeyF4)
	KeyF5                 = Key(tcell.KeyF5)
	KeyF6                 = Key(tcell.KeyF6)
	KeyF7                 = Key(tcell.KeyF7)
	KeyF8                 = Key(tcell.KeyF8)
	KeyF9                 = Key(tcell.KeyF9)
	KeyF10                = Key(tcell.KeyF10)
	KeyF11                = Key(tcell.KeyF11)
	KeyF12                = Key(tcell.KeyF12)
	KeyInsert             = Key(tcell.KeyInsert)
	KeyDelete             = Key(tcell.KeyDelete)
	KeyHome               = Key(tcell.KeyHome)
	KeyEnd                = Key(tcell.KeyEnd)
	KeyPgdn               = Key(tcell.KeyPgDn)
	KeyPgup               = Key(tcell.KeyPgUp)
	KeyArrowUp            = Key(tcell.KeyUp)
	KeyShiftArrowUp       = Key(tcell.KeyF62)
	KeyArrowDown          = Key(tcell.KeyDown)
	KeyShiftArrowDown     = Key(tcell.KeyF63)
	KeyArrowLeft          = Key(tcell.KeyLeft)
	KeyArrowRight         = Key(tcell.KeyRight)
)

// Keys combinations.
const (
	KeyCtrlTilde      = Key(tcell.KeyF64) // arbitrary assignment
	KeyCtrlSpace      = Key(tcell.KeyCtrlSpace)
	KeyCtrlA          = Key(tcell.KeyCtrlA)
	KeyCtrlB          = Key(tcell.KeyCtrlB)
	KeyCtrlC          = Key(tcell.KeyCtrlC)
	KeyCtrlD          = Key(tcell.KeyCtrlD)
	KeyCtrlE          = Key(tcell.KeyCtrlE)
	KeyCtrlF          = Key(tcell.KeyCtrlF)
	KeyCtrlG          = Key(tcell.KeyCtrlG)
	KeyBackspace      = Key(tcell.KeyBackspace)
	KeyCtrlH          = Key(tcell.KeyCtrlH)
	KeyTab            = Key(tcell.KeyTab)
	KeyBacktab        = Key(tcell.KeyBacktab)
	KeyCtrlI          = Key(tcell.KeyCtrlI)
	KeyCtrlJ          = Key(tcell.KeyCtrlJ)
	KeyCtrlK          = Key(tcell.KeyCtrlK)
	KeyCtrlL          = Key(tcell.KeyCtrlL)
	KeyEnter          = Key(tcell.KeyEnter)
	KeyCtrlM          = Key(tcell.KeyCtrlM)
	KeyCtrlN          = Key(tcell.KeyCtrlN)
	KeyCtrlO          = Key(tcell.KeyCtrlO)
	KeyCtrlP          = Key(tcell.KeyCtrlP)
	KeyCtrlQ          = Key(tcell.KeyCtrlQ)
	KeyCtrlR          = Key(tcell.KeyCtrlR)
	KeyCtrlS          = Key(tcell.KeyCtrlS)
	KeyCtrlT          = Key(tcell.KeyCtrlT)
	KeyCtrlU          = Key(tcell.KeyCtrlU)
	KeyCtrlV          = Key(tcell.KeyCtrlV)
	KeyCtrlW          = Key(tcell.KeyCtrlW)
	KeyCtrlX          = Key(tcell.KeyCtrlX)
	KeyCtrlY          = Key(tcell.KeyCtrlY)
	KeyCtrlZ          = Key(tcell.KeyCtrlZ)
	KeyEsc            = Key(tcell.KeyEscape)
	KeyCtrlUnderscore = Key(tcell.KeyCtrlUnderscore)
	KeySpace          = Key(32)
	KeyBackspace2     = Key(tcell.KeyBackspace2)
	KeyCtrl8          = Key(tcell.KeyBackspace2) // same key as in termbox-go

	// The following assignments were used in termbox implementation.
	// In tcell, these are not keys per se. But in gocui we have them
	// mapped to the keys so we have to use placeholder keys.

	KeyAltEnter       = Key(tcell.KeyF64) // arbitrary assignments
	MouseLeft         = Key(tcell.KeyF63)
	MouseRight        = Key(tcell.KeyF62)
	MouseMiddle       = Key(tcell.KeyF61)
	MouseRelease      = Key(tcell.KeyF60)
	MouseWheelUp      = Key(tcell.KeyF59)
	MouseWheelDown    = Key(tcell.KeyF58)
	MouseWheelLeft    = Key(tcell.KeyF57)
	MouseWheelRight   = Key(tcell.KeyF56)
	KeyCtrl2          = Key(tcell.KeyNUL) // termbox defines theses
	KeyCtrl3          = Key(tcell.KeyEscape)
	KeyCtrl4          = Key(tcell.KeyCtrlBackslash)
	KeyCtrl5          = Key(tcell.KeyCtrlRightSq)
	KeyCtrl6          = Key(tcell.KeyCtrlCarat)
	KeyCtrl7          = Key(tcell.KeyCtrlUnderscore)
	KeyCtrlSlash      = Key(tcell.KeyCtrlUnderscore)
	KeyCtrlRsqBracket = Key(tcell.KeyCtrlRightSq)
	KeyCtrlBackslash  = Key(tcell.KeyCtrlBackslash)
	KeyCtrlLsqBracket = Key(tcell.KeyCtrlLeftSq)
)

// Modifiers.
const (
	ModNone   Modifier = Modifier(0)
	ModAlt             = Modifier(tcell.ModAlt)
	ModMotion          = Modifier(2) // just picking an arbitrary number here that doesn't clash with tcell.ModAlt
	// ModCtrl doesn't work with keyboard keys. Use CtrlKey in Key and ModNone. This is was for mouse clicks only (tcell.v1)
	// ModCtrl = Modifier(tcell.ModCtrl)
)
