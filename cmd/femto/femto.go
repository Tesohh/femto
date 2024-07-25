package main

import (
	"log"

	"github.com/Tesohh/femto/editor"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	editor := editor.Editor{}
	editor.Setup()

	if err := ebiten.RunGame(&editor); err != nil {
		log.Fatal(err)
	}
}
