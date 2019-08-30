// Copyright 2019 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package event is a package that models events that occur during
// the execution of a program.
package event

// Event is an interface that custom events should implement.
// It is empty for now because there are no general methods
// required of events yet.
type Event interface {
}

// KeyCharacter is an event that occurs when a character is actually typed on
// the keyboard. This may be provided by an input method.
type KeyCharacter struct {
	// Code is the key code of the key typed.
	// TODO: this should change later from an int to an enumeration type.
	Code int
	// Modifiers are the modifiers pressed together with the key.
	// TODO: this should change later from an int to an enumeration type.
	Modifiers int
	// Character is the character that was typed.
	Character rune
}

// KeyDown is an event that occurs when a key is pressed on the keyboard.
type KeyDown struct {
	// Code is the key code of the key pressed or released.
	// TODO: this should change later from an int to an enumeration type.
	Code int
	// Modifiers are the modifiers pressed together with the key.
	// TODO: this should change later from an int to an enumeration type.
	Modifiers int
}

// KeyUp is an event that occurs when a key is released on the keyboard.
// The data is the same as for a KeyDown event.
type KeyUp KeyDown

// GamepadAxis is for event where an axis on a gamepad changes.
type GamepadAxis struct {
	// ID represents which gamepad caused the event.
	ID int
	// Axis is the axis of the game pad that changed position.
	Axis int
	// Position is the psoition of the axis after the change.
	// It varies between -1.0 and 1.0.
	Position float32
}

// GamepadButtonDown is a gamepad button press event.
type GamepadButtonDown struct {
	// ID represents which gamepad caused the event.
	ID int
	// Button is the button that was pressed on the game pad.
	Button int
	// Pressure is the pressure that is applied to the gamepad button.
	// It varies between 0.0 for not pressed, and 1.0 for completely pressed.
	Pressure float32
}

// GamepadButtonDown is a gamepad button release event.
// The data is identical to a GamePadButtonDown event.
type GamepadButtonUp GamepadButtonDown

// GamepadAttach happens when a new gamepad is attached.
type GamepadAttach struct {
	// ID represents which gamepad caused the event.
	ID int
	// Axes represents the amount of axes the gamepad has.
	Axes int
	// Buttons represents the amount of buttons the gamepad has.
	Buttons int
}

// GamepadDetach happens when a gamepad is detached.
type GamepadDetach struct {
	// ID represents which gamepad caused the event.
	ID int
}

// MouseMove is a mouse movement event.
type MouseMove struct {
	// X is the X position of the mouse pointer.
	// This value is expressed in device independent pixels.
	X float32
	// Y is the Y position of the mouse pointer.
	// This value is expressed in device independent pixels.
	Y float32
	// DeltaX is the change in X since the last MouseMove event.
	// This value is expressed in device independent pixels.
	DeltaX float32
	// DeltaY is the change in Y since the last MouseMove event.
	// This value is expressed in device independent pixels.
	DeltaY float32
}

// MouseWheel is a mouse wheel event.
type MouseWheel struct {
	// X is the X position of the mouse wheel.
	// This value is expressed in arbitrary units.
	// It increases when the mouse wheel is scrolled downwards,
	// and decreases when the mouse is scrolled upwards.
	X float32
	// Y is the Y position of the mouse wheel.
	// This value is expressed in arbitrary units.
	// It increases when the mouse wheel is scrolled to the right,
	// and decreases when the mouse is scrolled to the left.
	Y float32
	// DeltaX is the change in X since the last MouseWheel event.
	// This value is expressed in arbitrary units.
	// It is positive when the mouse wheel is scrolled downwards,
	// and negative when the mouse is scrolled upwards.
	DeltaX float32
	// DeltaY is the change in Y since the last MouseWheel event.
	// This value is expressed in arbitrary units.
	// It is positive when the mouse wheel is scrolled to the right,
	// and negative when the mouse is scrolled to the left.
	DeltaY float32
}

// MouseButtonDown is a mouse button press event.
type MouseButtonDown struct {
	// X is the X position of the mouse pointer.
	// This value is expressed in device independent pixels.
	X float32
	// Y is the Y position of the mouse pointer.
	// This value is expressed in device independent pixels.
	Y float32
	// Button is the button on the mouse that was pressed.
	// TODO: this should change later from an int to an enumeration type.
	Button int
	// Pressure is the pressure applied on the mouse button.
	// It varies between 0.0 for not pressed, and 1.0 for completely pressed.
	Pressure float32
}

// MouseButtonDown is a mouse button Release event.
// The data is identical to a MouseButtonDown event.
type MouseButtonUp MouseButtonDown

// MouseEnter occurs when the mouse enters the view window.
type MouseEnter struct {
	// X is the X position of the mouse pointer.
	// This value is expressed in device independent pixels.
	X float32
	// Y is the Y position of the mouse pointer.
	// This value is expressed in device independent pixels.
	Y float32
}

// MouseLeave occurs when the mouse leaves the view window.
// The data is identical to MouseEnter.
type MouseLeave MouseEnter

// ViewUpdate occurs when the application is ready to update
// the next frame on the view port.
type ViewUpdate struct {
	// No data neccesary, for now.
}

// ViewSize occurs when the size of the application's view port changes.
type ViewSize struct {
	// Width is the width of the view.
	// This value is expressed in device independent pixels.
	Width int
	// Height is the height of the view.
	// This value is expressed in device independent pixels.
	Height int
}

// TouchBegin occurs when a touch begins.
type TouchBegin struct {
	// ID identifies the touch that caused the touch event.
	ID int
	// X is the X position of the touch.
	// This value is expressed in device independent pixels.
	X float32
	// Y is the Y position of the touch.
	// This value is expressed in device independent pixels.
	Y float32
	// Pressure is the pressure applied to the touch.
	// It varies between 0.0 for not pressed, and 1.0 for completely pressed.
	Pressure float32
	// Primary represents whether the touch event is the primary touch or not.
	// If it is true, then it is a primary touch.
	// Otherwise it is false.
	Primary bool
}

// TouchMove occurs when a touch moved, or in other words, is dragged.
type TouchMove struct {
	// ID identifies the touch that caused the touch event.
	ID int
	// X is the X position of the touch.
	// This value is expressed in device independent pixels.
	X float32
	// Y is the Y position of the touch.
	// This value is expressed in device independent pixels.
	Y float32
	// DeltaX is the change in X since last touch event.
	// This value is expressed in device independent pixels.
	DeltaX float32
	// Deltay is the change in Y since last touch event.
	// This value is expressed in device independent pixels.
	DeltaY float32
	// Pressure of applied touch.
	// It varies between 0.0 for not pressed, and 1.0 for completely pressed.
	Pressure float32
	// Primary represents whether the touch event is the primary touch or not.
	// If it is true, then it is a primary touch.
	// If it is false then it is not.
	Primary bool
}

// TouchEnd occurs when a touch ends.
// The data is the same as for a TouchMove event.
type TouchEnd TouchMove

// TouchCancel occurs when a touch is canceled.
// This can happen in various situations, depending on the underlying platform,
// for example when the aplication loses focus.
type TouchCancel struct {
	// ID identifies the touch that caused the touch event.
	ID int
}
