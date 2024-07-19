package main

import (
	"log"

	"github.com/Djosar/kro-ecs/app/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	game, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
		return
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
		return
	}
}
