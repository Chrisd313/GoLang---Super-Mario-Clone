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
		{736, 202, 32, 64, colliderColor},
		{912, 202, 32, 64, colliderColor},

		// {320, 202, 16, 16, colliderColor2},
		// {352, 202, 16, 16, colliderColor2},
		// {384, 202, 16, 16, colliderColor2},

		// {0, 266, 1104, 32, colliderColor},
		// {1136, 266, 240, 32, colliderColor},
		// {1424, 266, 1024, 32, colliderColor},
		// {2480, 266, 896, 32, colliderColor},
	}

	grounds = []Collider{
		// {0, 266, 1104, 16, colliderColor2},

		{0, 266, 1104, 32, colliderColor},
		{1136, 266, 240, 32, colliderColor},
		{1424, 266, 1024, 32, colliderColor},
		{2480, 266, 896, 32, colliderColor},
	}

	blocks = []Collider{
		{320, 202, 80, 16, colliderColor2},
		{256, 202, 16, 16, colliderColor2},
		{1232, 202, 48, 16, colliderColor2},
		{1280, 138, 128, 16, colliderColor2},
		{1456, 138, 64, 16, colliderColor2},
		{1504, 202, 16, 16, colliderColor2},
		{1600, 202, 32, 16, colliderColor2},
		{1696, 202, 16, 16, colliderColor2},
		{1744, 202, 16, 16, colliderColor2},
		{1744, 138, 16, 16, colliderColor2},
		{1792, 202, 16, 16, colliderColor2},
		// {352, 138, 16, 16, colliderColor2},


		// {352, 202, 16, 16, colliderColor2},
		// {384, 202, 16, 16, colliderColor2},

		// {336, 202, 16, 16, colliderColor3},
		// {368, 202, 16, 16, colliderColor3},
	}

	coinBlocks = []Collider{
		{256, 204, 16, 16, colliderColor3},
		{336, 204, 16, 16, colliderColor3},
		{368, 204, 16, 16, colliderColor3},
		{352, 140, 16, 16, colliderColor3},
		{1248, 204, 16, 16, colliderColor3},
		{1504, 140, 16, 16, colliderColor3},
		{1696, 204, 16, 16, colliderColor2},
		{1744, 204, 16, 16, colliderColor2},
		{1744, 140, 16, 16, colliderColor2},
		{1792, 204, 16, 16, colliderColor2},
	}

	currentPlatformStart float32 = 0
	currentPlatformEnd   float32 = 0

	debugColorYellow = rl.NewColor(255, 255, 10, 255)

	debugColorPurple = rl.NewColor(255, 50, 255, 255)

	debugColorTeal = rl.NewColor(50, 255, 255, 255)

	debugColor = rl.NewColor(100, 255, 10, 255)

	colliderHeight float32
)

func drawColliders() {
	for _, current_Pipe := range pipes {
		rl.DrawRectangle(current_Pipe.posX, current_Pipe.posY, current_Pipe.width, current_Pipe.height, current_Pipe.Color)

		rl.DrawCircle(current_Pipe.posX, current_Pipe.posY, 2, debugColorYellow)
		rl.DrawCircle(current_Pipe.posX+current_Pipe.width, current_Pipe.posY, 2, debugColorPurple)
		rl.DrawCircle(current_Pipe.posX, current_Pipe.posY+current_Pipe.height, 2, debugColorTeal)
		rl.DrawCircle(current_Pipe.posX+current_Pipe.width, current_Pipe.posY+current_Pipe.height, 2, debugColor)

		if rl.CheckCollisionRecs(playerDest, rl.NewRectangle(float32(current_Pipe.posX), float32(current_Pipe.posY), float32(current_Pipe.width), float32(current_Pipe.height))) {

			currentPlatformStart = float32(current_Pipe.posX)
			currentPlatformEnd = float32(current_Pipe.posX) + float32(current_Pipe.width)

			var xDistance float32
			var yDistance float32

			xDistance, yDistance = CalculateAABBDistanceTo(current_Pipe)

			// fmt.Println(xDistance, yDistance)

			var xAxisTimeToCollide float32 = float32(math.Abs(float64(xDistance) / float64(velocityX)))
			var yAxisTimeToCollide float32 = float32(math.Abs(float64(yDistance) / float64(velocityY)))

			// fmt.Println("X Time: ", xAxisTimeToCollide, " | Y Time: ", yAxisTimeToCollide)

			if xAxisTimeToCollide < yAxisTimeToCollide {

				playerJumping = false

				// fmt.Println("Collision on the X axis")
				if playerDest.X < float32(current_Pipe.posX) {
					canMoveRight = false
					playerDest.X = float32(current_Pipe.posX) - playerDest.Width
					colliderHeight = float32(current_Pipe.posY)
				} else if playerDest.X > float32(current_Pipe.posX-current_Pipe.width) {
					playerDest.X = float32(current_Pipe.posX) + float32(current_Pipe.width)
					canMoveLeft = false
				}
			} else {
				// fmt.Println("Collsion on the Y axis")
				playerDest.Y = float32(current_Pipe.posY) - playerDest.Height
				playerGrounded = true
				playerJumping = false
				velocityY = 0
			}
		}
	}

	for _, current_Ground := range grounds {
		rl.DrawRectangle(current_Ground.posX, current_Ground.posY, current_Ground.width, current_Ground.height, current_Ground.Color)

		if rl.CheckCollisionRecs(playerDest, rl.NewRectangle(float32(current_Ground.posX), float32(current_Ground.posY), float32(current_Ground.width), float32(current_Ground.height))) {

			currentPlatformStart = float32(current_Ground.posX)
			currentPlatformEnd = float32(current_Ground.posX) + float32(current_Ground.width)

			var xDistance float32
			var yDistance float32

			xDistance, yDistance = CalculateAABBDistanceTo(current_Ground)

			// fmt.Println(xDistance, yDistance)

			var xAxisTimeToCollide float32 = float32(math.Abs(float64(xDistance) / float64(velocityX)))
			var yAxisTimeToCollide float32 = float32(math.Abs(float64(yDistance) / float64(velocityY)))

			// fmt.Println("X Time: ", xAxisTimeToCollide, " | Y Time: ", yAxisTimeToCollide)

			if xAxisTimeToCollide < yAxisTimeToCollide {

				playerJumping = false

				// fmt.Println("Collision on the X axis")
				if playerDest.X < float32(current_Ground.posX) {
					canMoveRight = false
					playerDest.X = float32(current_Ground.posX) - playerDest.Width
					colliderHeight = float32(current_Ground.posY)
				} else if playerDest.X > float32(current_Ground.posX-current_Ground.width) {
					playerDest.X = float32(current_Ground.posX) + float32(current_Ground.width)
					canMoveLeft = false
				}
			} else {
				// fmt.Println("Collsion on the Y axis")

				// fmt.Println("Player Y: ", playerDest.Y, "| Collider Y: ", current_Pipe.posY, "\n Player X: ", playerDest.X, "| Collider X: ", current_Pipe.posX)

				playerDest.Y = float32(current_Ground.posY) - playerDest.Height
				playerGrounded = true
				playerJumping = false
				velocityY = 0
			}
		}
	}

	for _, current_Block := range blocks {
		rl.DrawRectangle(current_Block.posX, current_Block.posY, current_Block.width, current_Block.height, current_Block.Color)

		if rl.CheckCollisionRecs(playerDest, rl.NewRectangle(float32(current_Block.posX), float32(current_Block.posY), float32(current_Block.width), float32(current_Block.height))) {

			var xDistance float32
			var yDistance float32

			currentPlatformStart = float32(current_Block.posX)
			currentPlatformEnd = float32(current_Block.posX) + float32(current_Block.width)

			xDistance, yDistance = CalculateAABBDistanceTo(current_Block)
			// fmt.Println(xDistance, yDistance)

			var xAxisTimeToCollide float32 = float32(math.Abs(float64(xDistance) / float64(velocityX)))
			var yAxisTimeToCollide float32 = float32(math.Abs(float64(yDistance) / float64(velocityY)))

			// fmt.Println("X Time: ", xAxisTimeToCollide, " | Y Time: ", yAxisTimeToCollide)

			if xAxisTimeToCollide < yAxisTimeToCollide {

				playerJumping = false

				// fmt.Println("Collision on the X axis")
				if playerDest.X < float32(current_Block.posX) {
					playerDest.X = float32(current_Block.posX) - playerDest.Width
					playerJumping = false
				} else if playerDest.X > float32(current_Block.posX-current_Block.width) {
					playerDest.X = float32(current_Block.posX) + float32(current_Block.width)
					playerJumping = false
				}
			} else {
				// fmt.Println("Collsion on the Y axis")

				// fmt.Println("Player Y: ", playerDest.Y, "| Collider Y: ", current_Block.posY, "\n Player X: ", playerDest.X, "| Collider X: ", current_Block.posX, "\n Block Y+H: ", current_Block.posY+current_Block.height, "| Block Y + Player H: ", current_Block.posY-playerDest.ToInt32().Height)

				if playerDest.Y < float32(current_Block.posY) {
					playerDest.Y = float32(current_Block.posY) - playerDest.Height
					playerGrounded = true
					playerJumping = false
					velocityY = 0
				} else {
					playerDest.Y = float32(current_Block.posY) + float32(current_Block.height)
					playerJumping = false
					velocityY = 0
				}
			}

		}
	}

	// i := 0

	for i, current_coinBlock := range coinBlocks {
		rl.DrawRectangle(current_coinBlock.posX, current_coinBlock.posY, current_coinBlock.width, current_coinBlock.height, current_coinBlock.Color)

		// var beenHit bool = false

		if rl.CheckCollisionRecs(playerDest, rl.NewRectangle(float32(current_coinBlock.posX), float32(current_coinBlock.posY), float32(current_coinBlock.width), float32(current_coinBlock.height))) {

			currentPlatformStart = float32(current_coinBlock.posX)
			currentPlatformEnd = float32(current_coinBlock.posX) + float32(current_coinBlock.width)

			var xDistance float32
			var yDistance float32

			xDistance, yDistance = CalculateAABBDistanceTo(current_coinBlock)

			// fmt.Println(xDistance, yDistance)

			var xAxisTimeToCollide float32 = float32(math.Abs(float64(xDistance) / float64(velocityX)))
			var yAxisTimeToCollide float32 = float32(math.Abs(float64(yDistance) / float64(velocityY)))

			// fmt.Println("X Time: ", xAxisTimeToCollide, " | Y Time: ", yAxisTimeToCollide)

			if xAxisTimeToCollide < yAxisTimeToCollide {

				playerJumping = false

				// fmt.Println("Collision on the X axis")
				if playerDest.X < float32(current_coinBlock.posX) {
					playerDest.X = float32(current_coinBlock.posX) - playerDest.Width
					playerJumping = false
				} else if playerDest.X > float32(current_coinBlock.posX-current_coinBlock.width) {
					playerDest.X = float32(current_coinBlock.posX) + float32(current_coinBlock.width)
					playerJumping = false
				}
			} else {
				// fmt.Println("Collsion on the Y axis")

				// fmt.Println("Player Y: ", playerDest.Y, "| Collider Y: ", current_Block.posY, "\n Player X: ", playerDest.X, "| Collider X: ", current_Block.posX, "\n Block Y+H: ", current_Block.posY+current_Block.height, "| Block Y + Player H: ", current_Block.posY-playerDest.ToInt32().Height)

				if playerDest.Y < float32(current_coinBlock.posY) {
					playerDest.Y = float32(current_coinBlock.posY) - playerDest.Height
					playerGrounded = true
					playerJumping = false
					velocityY = 0
				} else {
					playerDest.Y = float32(current_coinBlock.posY) + float32(current_coinBlock.height)
					playerJumping = false
					velocityY = 0
					fmt.Println("HIT COIN BLOCK")
					coinCount++
					fmt.Println(coinCount)

					// coinBlocks[i] = current_coinBlock
					coinBlocks = append(coinBlocks[:i], coinBlocks[i+1:]...)

					// coinBlocks = append(coinBlocks, current_coinBlock)
				}
				// i++
				// coinBlocks = coinBlocks[:i]
				// i--
				fmt.Println(coinBlocks)
			}
		}


	}
}

func CalculateAABBDistanceTo(e2 Collider) (float32, float32) {
	var dx1 float32
	var dy1 float32

	if playerDest.X < float32(e2.posX) {
		dx1 = float32(e2.posX) - (playerDest.X + playerDest.Width)
	} else if playerDest.X > float32(e2.posX) {
		dx1 = float32(playerDest.X) - (float32(e2.posX) + float32(e2.width))
	}

	if playerDest.Y < float32(e2.posY) {
		dy1 = float32(e2.posY) - (playerDest.Y + playerDest.Height)
	} else if playerDest.Y > float32(e2.posY) {
		dy1 = float32(playerDest.Y) - (float32(e2.posY) + float32(e2.height))
	}

	return dx1, dy1
}

