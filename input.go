package redpix

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type InputKey int
type InputAction int

const (
	IN_NONE InputKey = iota
	IN_PLAYER_FORWARD
	IN_PLAYER_BACKWARD
	IN_PLAYER_LEFT
	IN_PLAYER_RIGHT
	IN_PLAYER_STRAFE_LEFT
	IN_PLAYER_STRAFE_RIGHT
)

const (
	IN_ACT_NONE InputAction = iota	
	IN_ACT_PRESSED
	IN_ACT_RELEASED
	IN_ACT_REPEATED
)

type InputEvent struct {
	Key InputKey
	Action InputAction
}

func inputHandler(_ *glfw.Window, key glfw.Key, _ int, action glfw.Action, _ glfw.ModifierKey) {
	event := InputEvent{Key: IN_NONE, Action: IN_ACT_NONE}
	switch key {
	case glfw.KeyW:
		event.Key = IN_PLAYER_FORWARD
	case glfw.KeyK:
		event.Key = IN_PLAYER_FORWARD
	case glfw.KeyA:
		event.Key = IN_PLAYER_LEFT
	case glfw.KeyH:
		event.Key = IN_PLAYER_LEFT
	case glfw.KeyS:
		event.Key = IN_PLAYER_BACKWARD
	case glfw.KeyJ:
		event.Key = IN_PLAYER_BACKWARD
	case glfw.KeyD:
		event.Key = IN_PLAYER_RIGHT
	case glfw.KeyL:
		event.Key = IN_PLAYER_RIGHT
	}

	switch action {
	case glfw.Press:
		event.Action = IN_ACT_PRESSED
	case glfw.Release:
		event.Action = IN_ACT_RELEASED
	case glfw.Repeat:
		event.Action = IN_ACT_REPEATED
	}

	if event.Key != IN_NONE && event.Action != IN_ACT_NONE {
		instance.inputHandler(event)
	}
}

