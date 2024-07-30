package editor

import (
	"github.com/Tesohh/femto/humankey"
)

var defaultKeymap = humankey.HumanKeymap{
	"normal": {
		"ctrl+c":  "quit",
		"space q": "quit",
	},
}
