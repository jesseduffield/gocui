// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gocui

// Editor interface must be satisfied by gocui editors.
type Editor interface {
	Edit(v *View, key Key, mod Modifier) bool
}

// The EditorFunc type is an adapter to allow the use of ordinary functions as
// Editors. If f is a function with the appropriate signature, EditorFunc(f)
// is an Editor object that calls f.
type EditorFunc func(v *View, key Key, mod Modifier) bool

// Edit calls f(v, key, ch, mod)
func (f EditorFunc) Edit(v *View, key Key, mod Modifier) bool {
	return f(v, key, mod)
}

// DefaultEditor is the default editor.
var DefaultEditor Editor = EditorFunc(SimpleEditor)

// SimpleEditor is used as the default gocui editor.
func SimpleEditor(v *View, key Key, mod Modifier) bool {
	switch {
	case (key.KeyName() == KeyBackspace || key.KeyName() == KeyBackspace2) && (mod&ModAlt) != 0,
		key.KeyName() == KeyCtrlW:
		v.TextArea.BackSpaceWord()
	case key.KeyName() == KeyBackspace || key.KeyName() == KeyBackspace2:
		v.TextArea.BackSpaceChar()
	case key.KeyName() == KeyCtrlD || key.KeyName() == KeyDelete:
		v.TextArea.DeleteChar()
	case key.KeyName() == KeyArrowDown:
		v.TextArea.MoveCursorDown()
	case key.KeyName() == KeyArrowUp:
		v.TextArea.MoveCursorUp()
	case (key.KeyName() == KeyArrowLeft || key.Ch() == 'b') && (mod&ModAlt) != 0:
		v.TextArea.MoveLeftWord()
	case key.KeyName() == KeyArrowLeft:
		v.TextArea.MoveCursorLeft()
	case (key.KeyName() == KeyArrowRight || key.Ch() == 'f') && (mod&ModAlt) != 0:
		v.TextArea.MoveRightWord()
	case key.KeyName() == KeyArrowRight:
		v.TextArea.MoveCursorRight()
	case key.KeyName() == KeyEnter:
		v.TextArea.TypeCharacter("\n")
	case key.KeyName() == KeySpace:
		v.TextArea.TypeCharacter(" ")
	case key.KeyName() == KeyInsert:
		v.TextArea.ToggleOverwrite()
	case key.KeyName() == KeyCtrlU:
		v.TextArea.DeleteToStartOfLine()
	case key.KeyName() == KeyCtrlK:
		v.TextArea.DeleteToEndOfLine()
	case key.KeyName() == KeyCtrlA || key.KeyName() == KeyHome:
		v.TextArea.GoToStartOfLine()
	case key.KeyName() == KeyCtrlE || key.KeyName() == KeyEnd:
		v.TextArea.GoToEndOfLine()
	case key.KeyName() == KeyCtrlW:
		v.TextArea.BackSpaceWord()
	case key.KeyName() == KeyCtrlY:
		v.TextArea.Yank()
	case key.Ch() != 0:
		v.TextArea.TypeCharacter(string(key.Ch()))
	default:
		return false
	}

	v.RenderTextArea()

	return true
}
