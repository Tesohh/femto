package editor

import (
	"image"
	"log"
	"os"
	"runtime"

	"github.com/asaskevich/EventBus"
	"github.com/hajimehoshi/ebiten/v2"
)

type Editor struct {
	Tabs  []Tab
	TabId int
	Bus   EventBus.EventBus
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
}

func (e *Editor) Layout(w int, h int) (int, int) {
	return ebiten.Monitor().Size()
}
