// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gocui

import (
	"strings"

	"github.com/gdamore/tcell/v2"
)

// KeyName represents special keys or keys combinations.
type KeyName tcell.Key

// Modifier allows to define special keys combinations. They can be used
// in combination with Keys or Runes when a new keybinding is defined.
type Modifier tcell.ModMask

// Keybindings are used to link a given key-press event with a handler.
type keybinding struct {
	viewName string
	key      Key
	mod      Modifier
	handler  func(*Gui, *View) error
}

// Parse takes the input string and extracts the keybinding.
// Returns a Key / rune, a Modifier and an error.
func Parse(input string) (Key, Modifier, error) {
	if len(input) == 1 {
		r := rune(input[0])
		return Key{KeyName(tcell.KeyRune), r}, ModNone, nil
	}

	var modifier Modifier
	cleaned := make([]string, 0)

	tokens := strings.SplitSeq(input, "+")
	for t := range tokens {
		normalized := strings.Title(strings.ToLower(t))
		if t == "Alt" {
			modifier = ModAlt
			continue
		}
		cleaned = append(cleaned, normalized)
	}

	key, exist := translate[strings.Join(cleaned, "")]
	if !exist {
		return Key{0, 0}, ModNone, ErrNoSuchKeybind
	}

	return KeyWithName(key), modifier, nil
}

// ParseAll takes an array of strings and returns a map of all keybindings.
func ParseAll(input []string) (map[any]Modifier, error) {
	ret := make(map[any]Modifier)
	for _, i := range input {
		k, m, err := Parse(i)
		if err != nil {
			return ret, err
		}
		ret[k] = m
	}
	return ret, nil
}

// MustParse takes the input string and returns a Key / rune and a Modifier.
// It will panic if any error occured.
func MustParse(input string) (any, Modifier) {
	k, m, err := Parse(input)
	if err != nil {
		panic(err)
	}
	return k, m
}

// MustParseAll takes an array of strings and returns a map of all keybindings.
// It will panic if any error occured.
func MustParseAll(input []string) map[any]Modifier {
	result, err := ParseAll(input)
	if err != nil {
		panic(err)
	}
	return result
}

// newKeybinding returns a new Keybinding object.
func newKeybinding(viewname string, key Key, mod Modifier, handler func(*Gui, *View) error) (kb *keybinding) {
	kb = &keybinding{
		viewName: viewname,
		key:      key,
		mod:      mod,
		handler:  handler,
	}
	return kb
}

func eventMatchesKey(ev *GocuiEvent, key Key) bool {
	// assuming ModNone for now
	if ev.Mod != ModNone {
		return false
	}

	return key.KeyName() == ev.Key && key.Ch() == ev.Ch
}

// matchKeypress returns if the keybinding matches the keypress.
func (kb *keybinding) matchKeypress(key Key, mod Modifier) bool {
	return kb.key == key && kb.mod == mod
}

// translations for strings to keys
var translate = map[string]KeyName{
	"F1":             KeyF1,
	"F2":             KeyF2,
	"F3":             KeyF3,
	"F4":             KeyF4,
	"F5":             KeyF5,
	"F6":             KeyF6,
	"F7":             KeyF7,
	"F8":             KeyF8,
	"F9":             KeyF9,
	"F10":            KeyF10,
	"F11":            KeyF11,
	"F12":            KeyF12,
	"Insert":         KeyInsert,
	"Delete":         KeyDelete,
	"Home":           KeyHome,
	"End":            KeyEnd,
	"Pgup":           KeyPgup,
	"Pgdn":           KeyPgdn,
	"ArrowUp":        KeyArrowUp,
	"ShiftArrowUp":   KeyShiftArrowUp,
	"ArrowDown":      KeyArrowDown,
	"ShiftArrowDown": KeyShiftArrowDown,
	"ArrowLeft":      KeyArrowLeft,
	"ArrowRight":     KeyArrowRight,
	"CtrlTilde":      KeyCtrlTilde,
	"Ctrl2":          KeyCtrl2,
	"CtrlSpace":      KeyCtrlSpace,
	"CtrlA":          KeyCtrlA,
	"CtrlB":          KeyCtrlB,
	"CtrlC":          KeyCtrlC,
	"CtrlD":          KeyCtrlD,
	"CtrlE":          KeyCtrlE,
	"CtrlF":          KeyCtrlF,
	"CtrlG":          KeyCtrlG,
	"Backspace":      KeyBackspace,
	"CtrlH":          KeyCtrlH,
	"Tab":            KeyTab,
	"BackTab":        KeyBacktab,
	"CtrlI":          KeyCtrlI,
	"CtrlJ":          KeyCtrlJ,
	"CtrlK":          KeyCtrlK,
	"CtrlL":          KeyCtrlL,
	"Enter":          KeyEnter,
	"CtrlM":          KeyCtrlM,
	"CtrlN":          KeyCtrlN,
	"CtrlO":          KeyCtrlO,
	"CtrlP":          KeyCtrlP,
	"CtrlQ":          KeyCtrlQ,
	"CtrlR":          KeyCtrlR,
	"CtrlS":          KeyCtrlS,
	"CtrlT":          KeyCtrlT,
	"CtrlU":          KeyCtrlU,
	"CtrlV":          KeyCtrlV,
	"CtrlW":          KeyCtrlW,
	"CtrlX":          KeyCtrlX,
	"CtrlY":          KeyCtrlY,
	"CtrlZ":          KeyCtrlZ,
	"Esc":            KeyEsc,
	"CtrlLsqBracket": KeyCtrlLsqBracket,
	"Ctrl3":          KeyCtrl3,
	"Ctrl4":          KeyCtrl4,
	"CtrlBackslash":  KeyCtrlBackslash,
	"Ctrl5":          KeyCtrl5,
	"CtrlRsqBracket": KeyCtrlRsqBracket,
	"Ctrl6":          KeyCtrl6,
	"Ctrl7":          KeyCtrl7,
	"CtrlSlash":      KeyCtrlSlash,
	"CtrlUnderscore": KeyCtrlUnderscore,
	"Space":          KeySpace,
	"Backspace2":     KeyBackspace2,
	"Ctrl8":          KeyCtrl8,
	"Mouseleft":      MouseLeft,
	"Mousemiddle":    MouseMiddle,
	"Mouseright":     MouseRight,
	"Mouserelease":   MouseRelease,
	"MousewheelUp":   MouseWheelUp,
	"MousewheelDown": MouseWheelDown,
}

// Special keys.
const (
	KeyF1             KeyName = KeyName(tcell.KeyF1)
	KeyF2                     = KeyName(tcell.KeyF2)
	KeyF3                     = KeyName(tcell.KeyF3)
	KeyF4                     = KeyName(tcell.KeyF4)
	KeyF5                     = KeyName(tcell.KeyF5)
	KeyF6                     = KeyName(tcell.KeyF6)
	KeyF7                     = KeyName(tcell.KeyF7)
	KeyF8                     = KeyName(tcell.KeyF8)
	KeyF9                     = KeyName(tcell.KeyF9)
	KeyF10                    = KeyName(tcell.KeyF10)
	KeyF11                    = KeyName(tcell.KeyF11)
	KeyF12                    = KeyName(tcell.KeyF12)
	KeyInsert                 = KeyName(tcell.KeyInsert)
	KeyDelete                 = KeyName(tcell.KeyDelete)
	KeyHome                   = KeyName(tcell.KeyHome)
	KeyEnd                    = KeyName(tcell.KeyEnd)
	KeyPgdn                   = KeyName(tcell.KeyPgDn)
	KeyPgup                   = KeyName(tcell.KeyPgUp)
	KeyArrowUp                = KeyName(tcell.KeyUp)
	KeyShiftArrowUp           = KeyName(tcell.KeyF62)
	KeyArrowDown              = KeyName(tcell.KeyDown)
	KeyShiftArrowDown         = KeyName(tcell.KeyF63)
	KeyArrowLeft              = KeyName(tcell.KeyLeft)
	KeyArrowRight             = KeyName(tcell.KeyRight)
)

// Keys combinations.
const (
	KeyCtrlTilde      = KeyName(tcell.KeyF64) // arbitrary assignment
	KeyCtrlSpace      = KeyName(tcell.KeyCtrlSpace)
	KeyCtrlA          = KeyName(tcell.KeyCtrlA)
	KeyCtrlB          = KeyName(tcell.KeyCtrlB)
	KeyCtrlC          = KeyName(tcell.KeyCtrlC)
	KeyCtrlD          = KeyName(tcell.KeyCtrlD)
	KeyCtrlE          = KeyName(tcell.KeyCtrlE)
	KeyCtrlF          = KeyName(tcell.KeyCtrlF)
	KeyCtrlG          = KeyName(tcell.KeyCtrlG)
	KeyBackspace      = KeyName(tcell.KeyBackspace)
	KeyCtrlH          = KeyName(tcell.KeyCtrlH)
	KeyTab            = KeyName(tcell.KeyTab)
	KeyBacktab        = KeyName(tcell.KeyBacktab)
	KeyCtrlI          = KeyName(tcell.KeyCtrlI)
	KeyCtrlJ          = KeyName(tcell.KeyCtrlJ)
	KeyCtrlK          = KeyName(tcell.KeyCtrlK)
	KeyCtrlL          = KeyName(tcell.KeyCtrlL)
	KeyEnter          = KeyName(tcell.KeyEnter)
	KeyCtrlM          = KeyName(tcell.KeyCtrlM)
	KeyCtrlN          = KeyName(tcell.KeyCtrlN)
	KeyCtrlO          = KeyName(tcell.KeyCtrlO)
	KeyCtrlP          = KeyName(tcell.KeyCtrlP)
	KeyCtrlQ          = KeyName(tcell.KeyCtrlQ)
	KeyCtrlR          = KeyName(tcell.KeyCtrlR)
	KeyCtrlS          = KeyName(tcell.KeyCtrlS)
	KeyCtrlT          = KeyName(tcell.KeyCtrlT)
	KeyCtrlU          = KeyName(tcell.KeyCtrlU)
	KeyCtrlV          = KeyName(tcell.KeyCtrlV)
	KeyCtrlW          = KeyName(tcell.KeyCtrlW)
	KeyCtrlX          = KeyName(tcell.KeyCtrlX)
	KeyCtrlY          = KeyName(tcell.KeyCtrlY)
	KeyCtrlZ          = KeyName(tcell.KeyCtrlZ)
	KeyEsc            = KeyName(tcell.KeyEscape)
	KeyCtrlUnderscore = KeyName(tcell.KeyCtrlUnderscore)
	KeySpace          = KeyName(32)
	KeyBackspace2     = KeyName(tcell.KeyBackspace2)
	KeyCtrl8          = KeyName(tcell.KeyBackspace2) // same key as in termbox-go

	// The following assignments were used in termbox implementation.
	// In tcell, these are not keys per se. But in gocui we have them
	// mapped to the keys so we have to use placeholder keys.

	KeyAltEnter       = KeyName(tcell.KeyF64) // arbitrary assignments
	MouseLeft         = KeyName(tcell.KeyF63)
	MouseRight        = KeyName(tcell.KeyF62)
	MouseMiddle       = KeyName(tcell.KeyF61)
	MouseRelease      = KeyName(tcell.KeyF60)
	MouseWheelUp      = KeyName(tcell.KeyF59)
	MouseWheelDown    = KeyName(tcell.KeyF58)
	MouseWheelLeft    = KeyName(tcell.KeyF57)
	MouseWheelRight   = KeyName(tcell.KeyF56)
	KeyCtrl2          = KeyName(tcell.KeyNUL) // termbox defines theses
	KeyCtrl3          = KeyName(tcell.KeyEscape)
	KeyCtrl4          = KeyName(tcell.KeyCtrlBackslash)
	KeyCtrl5          = KeyName(tcell.KeyCtrlRightSq)
	KeyCtrl6          = KeyName(tcell.KeyCtrlCarat)
	KeyCtrl7          = KeyName(tcell.KeyCtrlUnderscore)
	KeyCtrlSlash      = KeyName(tcell.KeyCtrlUnderscore)
	KeyCtrlRsqBracket = KeyName(tcell.KeyCtrlRightSq)
	KeyCtrlBackslash  = KeyName(tcell.KeyCtrlBackslash)
	KeyCtrlLsqBracket = KeyName(tcell.KeyCtrlLeftSq)
)

// Modifiers.
const (
	ModNone   Modifier = Modifier(0)
	ModAlt             = Modifier(tcell.ModAlt)
	ModMotion          = Modifier(2) // just picking an arbitrary number here that doesn't clash with tcell.ModAlt
	// ModCtrl doesn't work with keyboard keys. Use CtrlKey in Key and ModNone. This is was for mouse clicks only (tcell.v1)
	// ModCtrl = Modifier(tcell.ModCtrl)
)
