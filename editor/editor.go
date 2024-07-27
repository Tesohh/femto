package editor

import (
	"image"
	"log"
	"os"
	"runtime"

	"github.com/Tesohh/femto/buffer"
	"github.com/asaskevich/EventBus"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/etxt"
)

type Editor struct {
	Tabs  []Tab
	TabId int

	Bus          EventBus.Bus
	TextRenderer *etxt.Renderer

	fonts *etxt.FontLibrary
}

func (e *Editor) tab() *Tab {
	return &e.Tabs[e.TabId]
}

func (e *Editor) Setup() {
	ebiten.SetWindowSize(640, 320)
	ebiten.SetTPS(60)
	ebiten.SetWindowTitle("femto")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowDecorated(false)

	if runtime.GOOS != "darwin" {
		file, err := os.Open("assets/appicon.png")
		if err != nil {
			log.Fatal(err)
		}
		icon, _, err := image.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		ebiten.SetWindowIcon([]image.Image{icon})
	}

	e.Bus = EventBus.New()

	e.Tabs = []Tab{{
		Buffer:   &buffer.SliceBuffer{},
		FilePath: "",
	}} // TEMP:
	e.tab().Buffer.Write([][]rune{
		[]rune("hello world"),
	})

	e.TextRenderer = etxt.NewStdRenderer()
	e.fonts = etxt.NewFontLibrary()
	e.fonts.ParseDirFonts("assets/fonts")
	e.TextRenderer.SetFont(e.fonts.GetFont("Iosevka Extended"))
	e.TextRenderer.SetSizePx(24)
}

func (e *Editor) Layout(w int, h int) (int, int) {
	return ebiten.Monitor().Size()
}
