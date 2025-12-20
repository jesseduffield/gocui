// Copyright 2025 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gocui

import "github.com/gdamore/tcell/v3"

type Key struct {
	keyName KeyName
	ch      string
}

func KeyWithName(keyName KeyName) Key {
	return Key{
		keyName: keyName,
		ch:      "",
	}
}

func KeyWithRune(ch rune) Key {
	return Key{
		keyName: KeyName(tcell.KeyRune),
		ch:      string(ch),
	}
}

func (k Key) KeyName() KeyName {
	return k.keyName
}

func (k Key) Ch() string {
	return k.ch
}

func (k Key) IsSet() bool {
	return k.keyName != 0
}
