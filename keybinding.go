// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gocui

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/jesseduffield/termbox-go"
)

// Keybidings are used to link a given key-press event with a handler.
type keybinding struct {
	viewName string
	key      Key
	ch       rune
	mod      Modifier
	handler  func(*Gui, *View) error
}

// Key contains all relevant information about the key
type ParsedKey struct {
	Value    Key
	Modifier Modifier
	Tokens   []string
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

// ParseKeybinding turns the input string into an actual Key.
// Returns an error when something goes wrong
func ParseKeybinding(input string) (ParsedKey, error) {

	if len(input) == 1 {
		K, _, err := getKey(rune(input[0]))
		if err != nil {
			return ParsedKey{}, err
		}
		return ParsedKey{K, ModNone, []string{input}}, nil
	}

	f := func(c rune) bool { return unicode.IsSpace(c) || c == '+' }
	tokens := strings.FieldsFunc(input, f)
	var normalizedTokens []string
	var modifier = ModNone

	for _, token := range tokens {
		normalized := token

		if value, exists := translate[normalized]; exists {
			normalized = value
		} else {
			normalized = strings.Title(normalized)
		}

		if normalized == "Alt" {
			modifier = ModAlt
			continue
		}

		if len(normalized) == 1 {
			normalizedTokens = append(normalizedTokens, strings.ToUpper(normalized))
			continue
		}

		normalizedTokens = append(normalizedTokens, normalized)
	}

	lookup := strings.Join(normalizedTokens, "")
	if !strings.Contains(lookup, "Mouse") {
		lookup = "Key" + strings.Join(normalizedTokens, "")
	}

	if key, exists := supportedKeybindings[lookup]; exists {
		return ParsedKey{key, modifier, normalizedTokens}, nil
	}

	if modifier != ModNone {
		return ParsedKey{0, modifier, normalizedTokens}, fmt.Errorf("unsupported keybinding: %s (+%+v)", lookup, modifier)
	}

	return ParsedKey{0, modifier, normalizedTokens}, fmt.Errorf("unsupported keybinding: %s", lookup)
}

// ParseAllKeybindings parses all strings to a Key.
// Returns an error when something goes wrong.
func ParseAllKeybindings(input string) ([]ParsedKey, error) {
	ret := make([]ParsedKey, 0)
	for _, value := range strings.Split(input, ",") {
		key, err := ParseKeybinding(value)
		if err != nil {
			return nil, fmt.Errorf("could not parse keybinding '%s' from request '%s': %+v", value, input, err)
		}
		ret = append(ret, key)
	}
	if len(ret) == 0 {
		return nil, fmt.Errorf("must have at least one keybinding")
	}
	return ret, nil
}

// MustParseKeybinding parses the input string to a key but instead an
// error, it panics when things go wrong.
// This forces the caller to react to an error
func MustParseKeybinding(input string) ParsedKey {
	if key, err := ParseKeybinding(input); err != nil {
		panic(err)
	} else {
		return key
	}
}

// MustParseAllKeybindings parses all the input strings to a key but instead an
// error, it panics when things go wrong.
// This forces the caller to react to an error
func MustParseAllKeybindings(input string) []ParsedKey {
	if key, err := ParseAllKeybindings(input); err != nil {
		panic(err)
	} else {
		return key
	}
}

// matchKeypress returns if the keybinding matches the keypress.
func (kb *keybinding) matchKeypress(key Key, ch rune, mod Modifier) bool {
	return kb.key == key && kb.ch == ch && kb.mod == mod
}

// matchView returns if the keybinding matches the current view.
func (kb *keybinding) matchView(v *View) bool {
	// if the user is typing in a field, ignore char keys
	if v == nil {
		return false
	}
	if v.Editable == true && kb.ch != 0 {
		return false
	}
	return kb.viewName == v.name
}

// Key represents special keys or keys combinations.
type Key termbox.Key

// Special keys.
const (
	KeyF1         Key = Key(termbox.KeyF1)
	KeyF2             = Key(termbox.KeyF2)
	KeyF3             = Key(termbox.KeyF3)
	KeyF4             = Key(termbox.KeyF4)
	KeyF5             = Key(termbox.KeyF5)
	KeyF6             = Key(termbox.KeyF6)
	KeyF7             = Key(termbox.KeyF7)
	KeyF8             = Key(termbox.KeyF8)
	KeyF9             = Key(termbox.KeyF9)
	KeyF10            = Key(termbox.KeyF10)
	KeyF11            = Key(termbox.KeyF11)
	KeyF12            = Key(termbox.KeyF12)
	KeyInsert         = Key(termbox.KeyInsert)
	KeyDelete         = Key(termbox.KeyDelete)
	KeyHome           = Key(termbox.KeyHome)
	KeyEnd            = Key(termbox.KeyEnd)
	KeyPgup           = Key(termbox.KeyPgup)
	KeyPgdn           = Key(termbox.KeyPgdn)
	KeyArrowUp        = Key(termbox.KeyArrowUp)
	KeyArrowDown      = Key(termbox.KeyArrowDown)
	KeyArrowLeft      = Key(termbox.KeyArrowLeft)
	KeyArrowRight     = Key(termbox.KeyArrowRight)

	MouseLeft      = Key(termbox.MouseLeft)
	MouseMiddle    = Key(termbox.MouseMiddle)
	MouseRight     = Key(termbox.MouseRight)
	MouseRelease   = Key(termbox.MouseRelease)
	MouseWheelUp   = Key(termbox.MouseWheelUp)
	MouseWheelDown = Key(termbox.MouseWheelDown)
)

// Keys combinations.
const (
	KeyCtrlTilde      Key = Key(termbox.KeyCtrlTilde)
	KeyCtrl2              = Key(termbox.KeyCtrl2)
	KeyCtrlSpace          = Key(termbox.KeyCtrlSpace)
	KeyCtrlA              = Key(termbox.KeyCtrlA)
	KeyCtrlB              = Key(termbox.KeyCtrlB)
	KeyCtrlC              = Key(termbox.KeyCtrlC)
	KeyCtrlD              = Key(termbox.KeyCtrlD)
	KeyCtrlE              = Key(termbox.KeyCtrlE)
	KeyCtrlF              = Key(termbox.KeyCtrlF)
	KeyCtrlG              = Key(termbox.KeyCtrlG)
	KeyBackspace          = Key(termbox.KeyBackspace)
	KeyCtrlH              = Key(termbox.KeyCtrlH)
	KeyTab                = Key(termbox.KeyTab)
	KeyCtrlI              = Key(termbox.KeyCtrlI)
	KeyCtrlJ              = Key(termbox.KeyCtrlJ)
	KeyCtrlK              = Key(termbox.KeyCtrlK)
	KeyCtrlL              = Key(termbox.KeyCtrlL)
	KeyEnter              = Key(termbox.KeyEnter)
	KeyCtrlM              = Key(termbox.KeyCtrlM)
	KeyCtrlN              = Key(termbox.KeyCtrlN)
	KeyCtrlO              = Key(termbox.KeyCtrlO)
	KeyCtrlP              = Key(termbox.KeyCtrlP)
	KeyCtrlQ              = Key(termbox.KeyCtrlQ)
	KeyCtrlR              = Key(termbox.KeyCtrlR)
	KeyCtrlS              = Key(termbox.KeyCtrlS)
	KeyCtrlT              = Key(termbox.KeyCtrlT)
	KeyCtrlU              = Key(termbox.KeyCtrlU)
	KeyCtrlV              = Key(termbox.KeyCtrlV)
	KeyCtrlW              = Key(termbox.KeyCtrlW)
	KeyCtrlX              = Key(termbox.KeyCtrlX)
	KeyCtrlY              = Key(termbox.KeyCtrlY)
	KeyCtrlZ              = Key(termbox.KeyCtrlZ)
	KeyEsc                = Key(termbox.KeyEsc)
	KeyCtrlLsqBracket     = Key(termbox.KeyCtrlLsqBracket)
	KeyCtrl3              = Key(termbox.KeyCtrl3)
	KeyCtrl4              = Key(termbox.KeyCtrl4)
	KeyCtrlBackslash      = Key(termbox.KeyCtrlBackslash)
	KeyCtrl5              = Key(termbox.KeyCtrl5)
	KeyCtrlRsqBracket     = Key(termbox.KeyCtrlRsqBracket)
	KeyCtrl6              = Key(termbox.KeyCtrl6)
	KeyCtrl7              = Key(termbox.KeyCtrl7)
	KeyCtrlSlash          = Key(termbox.KeyCtrlSlash)
	KeyCtrlUnderscore     = Key(termbox.KeyCtrlUnderscore)
	KeySpace              = Key(termbox.KeySpace)
	KeyBackspace2         = Key(termbox.KeyBackspace2)
	KeyCtrl8              = Key(termbox.KeyCtrl8)
)

// All the indirect translations
var translate = map[string]string{
	"/":        "Slash",
	"\\":       "Backslash",
	"[":        "LsqBracket",
	"]":        "RsqBracket",
	"_":        "Underscore",
	"escape":   "Esc",
	"~":        "Tilde",
	"pageup":   "Pgup",
	"pagedown": "Pgdn",
	"pgup":     "Pgup",
	"pgdown":   "Pgdn",
	"up":       "ArrowUp",
	"down":     "ArrowDown",
	"right":    "ArrowRight",
	"left":     "ArrowLeft",
	"ctl":      "Ctrl",
}

// All the direct translations
var supportedKeybindings = map[string]Key{
	"KeyF1":             KeyF1,
	"KeyF2":             KeyF2,
	"KeyF3":             KeyF3,
	"KeyF4":             KeyF4,
	"KeyF5":             KeyF5,
	"KeyF6":             KeyF6,
	"KeyF7":             KeyF7,
	"KeyF8":             KeyF8,
	"KeyF9":             KeyF9,
	"KeyF10":            KeyF10,
	"KeyF11":            KeyF11,
	"KeyF12":            KeyF12,
	"KeyInsert":         KeyInsert,
	"KeyDelete":         KeyDelete,
	"KeyHome":           KeyHome,
	"KeyEnd":            KeyEnd,
	"KeyPgup":           KeyPgup,
	"KeyPgdn":           KeyPgdn,
	"KeyArrowUp":        KeyArrowUp,
	"KeyArrowDown":      KeyArrowDown,
	"KeyArrowLeft":      KeyArrowLeft,
	"KeyArrowRight":     KeyArrowRight,
	"KeyCtrlTilde":      KeyCtrlTilde,
	"KeyCtrl2":          KeyCtrl2,
	"KeyCtrlSpace":      KeyCtrlSpace,
	"KeyCtrlA":          KeyCtrlA,
	"KeyCtrlB":          KeyCtrlB,
	"KeyCtrlC":          KeyCtrlC,
	"KeyCtrlD":          KeyCtrlD,
	"KeyCtrlE":          KeyCtrlE,
	"KeyCtrlF":          KeyCtrlF,
	"KeyCtrlG":          KeyCtrlG,
	"KeyBackspace":      KeyBackspace,
	"KeyCtrlH":          KeyCtrlH,
	"KeyTab":            KeyTab,
	"KeyCtrlI":          KeyCtrlI,
	"KeyCtrlJ":          KeyCtrlJ,
	"KeyCtrlK":          KeyCtrlK,
	"KeyCtrlL":          KeyCtrlL,
	"KeyEnter":          KeyEnter,
	"KeyCtrlM":          KeyCtrlM,
	"KeyCtrlN":          KeyCtrlN,
	"KeyCtrlO":          KeyCtrlO,
	"KeyCtrlP":          KeyCtrlP,
	"KeyCtrlQ":          KeyCtrlQ,
	"KeyCtrlR":          KeyCtrlR,
	"KeyCtrlS":          KeyCtrlS,
	"KeyCtrlT":          KeyCtrlT,
	"KeyCtrlU":          KeyCtrlU,
	"KeyCtrlV":          KeyCtrlV,
	"KeyCtrlW":          KeyCtrlW,
	"KeyCtrlX":          KeyCtrlX,
	"KeyCtrlY":          KeyCtrlY,
	"KeyCtrlZ":          KeyCtrlZ,
	"KeyEsc":            KeyEsc,
	"KeyCtrlLsqBracket": KeyCtrlLsqBracket,
	"KeyCtrl3":          KeyCtrl3,
	"KeyCtrl4":          KeyCtrl4,
	"KeyCtrlBackslash":  KeyCtrlBackslash,
	"KeyCtrl5":          KeyCtrl5,
	"KeyCtrlRsqBracket": KeyCtrlRsqBracket,
	"KeyCtrl6":          KeyCtrl6,
	"KeyCtrl7":          KeyCtrl7,
	"KeyCtrlSlash":      KeyCtrlSlash,
	"KeyCtrlUnderscore": KeyCtrlUnderscore,
	"KeySpace":          KeySpace,
	"KeyBackspace2":     KeyBackspace2,
	"KeyCtrl8":          KeyCtrl8,
	"MouseLeft":         MouseLeft,
	"MouseMiddle":       MouseMiddle,
	"MouseRight":        MouseRight,
	"MouseRelease":      MouseRelease,
	"MouseWheelUp":      MouseWheelUp,
	"MouseWheelDown":    MouseWheelDown,
}

// Modifier allows to define special keys combinations. They can be used
// in combination with Keys or Runes when a new keybinding is defined.
type Modifier termbox.Modifier

// Modifiers.
const (
	ModNone Modifier = Modifier(0)
	ModAlt           = Modifier(termbox.ModAlt)
)
