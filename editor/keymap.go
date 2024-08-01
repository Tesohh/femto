package editor

import (
	"github.com/Tesohh/femto/humankey"
)

var defaultKeymap = humankey.HumanKeymap{
	"normal": {
		"ctrl+c": "quit",
		"h":      "left",
		"left":   "left",
		"j":      "down",
		"down":   "down",
		"k":      "up",
		"up":     "up",
		"l":      "right",
		"right":  "right",
		"H":      "bigLeft",
		"L":      "bigRight",
	},
}
