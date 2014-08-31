// Copyright 2009 The XGB Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ui defines a basic graphical user interface programming model.
package ui

import (
	"image"
	"image/draw"
	"time"
)

// A Window represents a single graphics window.
type Window interface {
	// Screen returns an editable Image for the window.
	Screen() draw.Image
	// FlushImage flushes changes made to Screen() back to screen.
	FlushImage()
	// EventChan returns a channel carrying UI events such as key presses,
	// mouse movements and window resizes.
	EventChan() <-chan interface{}
	// Close closes the window.
	Close() error
}

// A KeyEvent is sent for a key press or release.
type KeyEvent struct {
	// The value k represents key k being pressed.
	// The value -k represents key k being released.
	// The specific set of key values is not specified,
	// but ordinary characters represent themselves.
	Key int
}

// A MouseEvent is sent for a button press or release or for a mouse movement.
type MouseEvent struct {
	// Buttons is a bit mask of buttons: 1<<0 is left, 1<<1 middle, 1<<2 right.
	// It represents button state and not necessarily the state delta: bit 0
	// being on means that the left mouse button is down, but does not imply
	// that the same button was up in the previous MouseEvent.
	Buttons int
	// Loc is the location of the cursor.
	Loc image.Point
	// Time is the event's timestamp.
	Time time.Time
}

// A ConfigEvent is sent each time the window's color model or size changes.
// The client should respond by calling Window.Screen to obtain a new image.
type ConfigEvent struct {
	Config image.Config
}

// An ErrEvent is sent when an error occurs.
type ErrEvent struct {
	Err error
}
