package plugins

import (
	"github.com/Tesohh/femto/editor"
	"github.com/gdamore/tcell/v2"
)

var RosePine = editor.DumbPlugin{
	Info: editor.PluginInfo{
		Id:          "femto.themes.rosepine",
		Author:      "femto",
		Name:        "Rose Pine",
		Description: "All natural pine, faux fur and a bit of soho vibes for the classy minimalist.",
	},
	Themes: map[string]editor.Theme{
		"rosepine.main": {
			Name: "Rose Pine (Main)",

			Default: tcell.StyleDefault.Background(tcell.NewHexColor(0x191724)).Foreground(tcell.NewHexColor(0xe0def4)),
			Borders: tcell.NewHexColor(0x524f67),

			Error: tcell.StyleDefault.Foreground(tcell.NewHexColor(0xeb6f92)),

			Red:       tcell.NewHexColor(0xeb6f92),
			Yellow:    tcell.NewHexColor(0xf6c177),
			Pink:      tcell.NewHexColor(0xebbcba),
			Blue:      tcell.NewHexColor(0x31748f),
			LightBlue: tcell.NewHexColor(0x9ccfd8),
			Purple:    tcell.NewHexColor(0xc4a7e7),

			NormalModeAccent: tcell.NewHexColor(0xc4a7e7),
			InsertModeAccent: tcell.NewHexColor(0xeb6f92),
		},
	},
}
