package main

import (
	"fmt"

	"github.com/solarlune/resolv"
)

var space *resolv.Space
var playerCollider *resolv.Object

func init() {
	space = resolv.NewSpace(3376, 480, 16, 16)

	// space.Add(
	// 	resolv.NewObject(896, 234, 64, 64),
	// )

	playerCollider = resolv.NewObject(float64(playerDest.X), float64(playerDest.Y), 8, 8)
	space.Add(playerCollider)

	for _, current_Pipe := range pipes {

		x := float64(current_Pipe.posX * 2)
		y := float64(current_Pipe.posY)
		width := float64(current_Pipe.width)
		height := float64(current_Pipe.height * 2)

		space.Add(
			resolv.NewObject(x, y, width, height, "pipeTag"),
		)
	}

	for _, current_Ground := range grounds {

		x := float64(current_Ground.posX)
		y := float64(current_Ground.posY + 6)
		width := float64(current_Ground.width)
		height := float64(current_Ground.height)

		space.Add(
			resolv.NewObject(x, y, width, height, "groundTag"),
		)
	}

	for _, current_Block := range blocks {

		x := float64(current_Block.posX)
		y := float64(current_Block.posY + 6)
		width := float64(current_Block.width)
		height := float64(current_Block.height)

		space.Add(
			resolv.NewObject(x, y, width, height, "blockTag"),
		)
	}

	fmt.Println("Pipe details: ", pipes[0])
}
