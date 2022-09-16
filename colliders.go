package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Collider struct {
	posX   int32
	posY   int32
	width  int32
	height int32
	Color  rl.Color
}

var (
	pipes = []Collider{
		// {0, 266, 1104, 20, colliderColor},
		{450, 234, 28, 32, colliderColor},
		{608, 218, 32, 48, colliderColor},
	}

	grounds = []Collider{
		{0, 266, 1104, 16, colliderColor},
	}

	blocks = []Collider{
		{320, 202, 16, 16, colliderColor2},
		{352, 202, 16, 16, colliderColor2},
		{384, 202, 16, 16, colliderColor2},
	}
)

func drawColliders() {
	for _, current_Pipe := range pipes {
		rl.DrawRectangle(current_Pipe.posX, current_Pipe.posY, current_Pipe.width, current_Pipe.height, current_Pipe.Color)

		if rl.CheckCollisionRecs(playerDest, rl.NewRectangle(float32(current_Pipe.posX), float32(current_Pipe.posY), float32(current_Pipe.width), float32(current_Pipe.height))) {

			var xDistance float32
			var yDistance float32

			var dx float32
			var dy float32

			if playerDest.X < float32(current_Pipe.posX) {
				dx = float32(current_Pipe.posX) - playerDest.Width
			} else if playerDest.X > float32(current_Pipe.posX-current_Pipe.width) {
				dx = float32(current_Pipe.posX) + float32(current_Pipe.width)
			}

			if playerDest.Y < float32(current_Pipe.posY) {
				dy = float32(current_Pipe.posY) - (playerDest.Y + playerDest.Height)
				playerGrounded = true
				playerJumping = false
			} else if playerDest.Y > float32(current_Pipe.posY) {
				dy = float32(current_Pipe.posY) + (float32(current_Pipe.posY) + float32(current_Pipe.height))
			}

			xDistance = dx
			yDistance = dy

			fmt.Println(xDistance, yDistance)

			var xAxisTimeToCollide float32 = float32(math.Abs(float64(xDistance) / float64(velocityX)))
			var yAxisTimeToCollide float32 = float32(math.Abs(float64(yDistance) / float64(velocityY)))


			if xAxisTimeToCollide < yAxisTimeToCollide {
				fmt.Println("Collision on the X axis")
				if playerDest.X < float32(current_Pipe.posX) {
					playerDest.X = float32(current_Pipe.posX) - playerDest.Width;
				} else if playerDest.X > float32(current_Pipe.posX-current_Pipe.width) {
					playerDest.X = float32(current_Pipe.posX) + float32(current_Pipe.width);
				}
			} else {
				fmt.Println("Collsion on the Y axis")
				playerDest.Y = float32(current_Pipe.posY) - playerDest.Height
			}
		} else {
			playerGrounded = false
		}
	}

	for _, current_Ground := range grounds {
		rl.DrawRectangle(current_Ground.posX, current_Ground.posY, current_Ground.width, current_Ground.height, current_Ground.Color)

		if rl.CheckCollisionRecs(playerDest, rl.NewRectangle(float32(current_Ground.posX), float32(current_Ground.posY), float32(current_Ground.width), float32(current_Ground.height))) {
			// fmt.Println("Ground collision detected")
			playerGrounded = true
			playerJumping = false
			playerDest.Y = float32(current_Ground.posY - current_Ground.height)
		}
	}

	for _, current_Block := range blocks {
		rl.DrawRectangle(current_Block.posX, current_Block.posY, current_Block.width, current_Block.height, current_Block.Color)

		if rl.CheckCollisionRecs(playerDest, rl.NewRectangle(float32(current_Block.posX), float32(current_Block.posY), float32(current_Block.width), float32(current_Block.height))) {
			fmt.Println("Block collision detected")
		}
	}
}
