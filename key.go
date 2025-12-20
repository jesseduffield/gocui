// Copyright 2026 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gocui

import "github.com/gdamore/tcell/v3"

type Key struct {
	keyName KeyName
	str     string
}

func NewKey(keyName KeyName, str string) Key {
	return Key{
		keyName: keyName,
		str:     str,
	}
}

func NewKeyName(keyName KeyName) Key {
	return Key{
		keyName: keyName,
		str:     "",
	}
}

func NewKeyRune(ch rune) Key {
	return Key{
		keyName: KeyName(tcell.KeyRune),
		str:     string(ch),
	}
}

func (k Key) KeyName() KeyName {
	return k.keyName
}

func (k Key) Str() string {
	return k.str
}

func (k Key) IsSet() bool {
	return k.keyName != 0
}

func (k Key) Equals(otherKey Key) bool {
	return k.keyName == otherKey.keyName && k.str == otherKey.str
}
