package main

import (
	"fmt"
	"math"

	"github.com/TwiN/go-color"
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
	staticItems = []Collider{

		//PIPES
		{448, 234, 32, 32, colliderColor},
		{608, 218, 32, 48, colliderColor},
		{736, 202, 32, 64, colliderColor},
		{912, 202, 32, 64, colliderColor},
		{2608, 234, 32, 32, colliderColor},
		{2864, 234, 32, 32, colliderColor},

		{-20, 80, 20, 300, colliderColor},
		{3376, 80, 20, 300, colliderColor},

		//RAMPS
		{2144, 250, 16, 16, colliderColor},
		{2160, 234, 16, 32, colliderColor},
		{2176, 218, 16, 48, colliderColor},
		{2192, 202, 16, 64, colliderColor},

		{2240, 202, 16, 64, colliderColor},
		{2256, 218, 16, 48, colliderColor},
		{2272, 234, 16, 32, colliderColor},
		{2288, 250, 16, 16, colliderColor},

		{2368, 250, 16, 16, colliderColor},
		{2384, 234, 16, 32, colliderColor},
		{2400, 218, 16, 48, colliderColor},
		{2416, 202, 16, 64, colliderColor},
		{2432, 202, 16, 64, colliderColor},

		{2480, 202, 16, 64, colliderColor},
		{2496, 218, 16, 48, colliderColor},
		{2512, 234, 16, 32, colliderColor},
		{2528, 250, 16, 16, colliderColor},

		{3168, 250, 16, 16, colliderColor},
	}

	grounds = []Collider{
		//FLOOR 1
		{0, 266, 1104, 1000, colliderColor},

		//FLOOR 2
		{1136, 266, 240, 1000, colliderColor},

		//FLOOR 3
		{1424, 266, 1024, 1000, colliderColor},

		//FLOOR 4
		{2480, 266, 896, 1000, colliderColor},
	}

	blocks = []Collider{
		{256, 202, 16, 16, colliderColor2},
		{320, 202, 16, 16, colliderColor2},
		{336, 202, 16, 16, colliderColor2},
		{352, 202, 16, 16, colliderColor2},
		{368, 202, 16, 16, colliderColor2},
		{384, 202, 16, 16, colliderColor2},
		{352, 138, 16, 16, colliderColor2},

		{1232, 202, 16, 16, colliderColor2},
		{1248, 202, 16, 16, colliderColor2},
		{1264, 202, 16, 16, colliderColor2},

		{1280, 138, 16, 16, colliderColor2},
		{1296, 138, 16, 16, colliderColor2},
		{1312, 138, 16, 16, colliderColor2},
		{1328, 138, 16, 16, colliderColor2},
		{1344, 138, 16, 16, colliderColor2},
		{1360, 138, 16, 16, colliderColor2},
		{1376, 138, 16, 16, colliderColor2},
		{1392, 138, 16, 16, colliderColor2},

		{1456, 138, 16, 16, colliderColor2},
		{1472, 138, 16, 16, colliderColor2},
		{1488, 138, 16, 16, colliderColor2},
		{1504, 138, 16, 16, colliderColor2},

		{1504, 202, 16, 16, colliderColor2},

		{1600, 202, 16, 16, colliderColor2},
		{1616, 202, 16, 16, colliderColor2},

		{1696, 202, 16, 16, colliderColor2},

		{1744, 202, 16, 16, colliderColor2},
		{1744, 138, 16, 16, colliderColor2},
		{1792, 202, 16, 16, colliderColor2},

		{1888, 202, 16, 16, colliderColor2},
		{1936, 138, 16, 16, colliderColor2},
		{1952, 138, 16, 16, colliderColor2},
		{1968, 138, 16, 16, colliderColor2},

		{2048, 138, 16, 16, colliderColor2},
		{2064, 138, 16, 16, colliderColor2},
		{2080, 138, 16, 16, colliderColor2},
		{2096, 138, 16, 16, colliderColor2},

		{2064, 202, 16, 16, colliderColor2},
		{2080, 202, 16, 16, colliderColor2},

		{2688, 202, 16, 16, colliderColor2},
		{2704, 202, 16, 16, colliderColor2},
		{2720, 202, 16, 16, colliderColor2},
		{2736, 202, 16, 16, colliderColor2},
	}

	coinBlocks = []Collider{
		{256, 205, 16, 16, colliderColor3},
		{336, 205, 16, 16, colliderColor3},
		{368, 205, 16, 16, colliderColor3},
		{352, 141, 16, 16, colliderColor3},
		{1248, 205, 16, 16, colliderColor3},
		{1504, 141, 16, 16, colliderColor3},
		{1696, 205, 16, 16, colliderColor3},
		{1744, 205, 16, 16, colliderColor3},
		{1744, 141, 16, 16, colliderColor3},
		{1792, 205, 16, 16, colliderColor3},

		{2064, 141, 16, 16, colliderColor3},
		{2080, 141, 16, 16, colliderColor3},

		{2720, 205, 16, 16, colliderColor3},
	}

	coinBlocksHit = []Collider{
		{256, 204, 16, 16, colliderColor3},
		{336, 204, 16, 16, colliderColor3},
		{368, 204, 16, 16, colliderColor3},
		{352, 140, 16, 16, colliderColor3},
		{1248, 204, 16, 16, colliderColor3},
		{1504, 140, 16, 16, colliderColor3},
		{1696, 204, 16, 16, colliderColor3},
		{1744, 204, 16, 16, colliderColor3},
		{1744, 140, 16, 16, colliderColor3},
		{1792, 204, 16, 16, colliderColor3},

		{2064, 140, 16, 16, colliderColor3},
		{2080, 140, 16, 16, colliderColor3},

		{2720, 204, 16, 16, colliderColor3},
	}

	currentPlatformStart float32 = 0
	currentPlatformEnd   float32 = 0
	colliderHeight float32

	debugColorYellow = rl.NewColor(255, 255, 10, uint8(debugColliderAlpha))
	debugColorPurple = rl.NewColor(255, 50, 255, uint8(debugColliderAlpha))
	debugColorTeal = rl.NewColor(50, 255, 255, uint8(debugColliderAlpha))
	debugColor = rl.NewColor(100, 255, 10, uint8(debugColliderAlpha))
)

func deathCollider() {

	var deathColliderRec rl.Rectangle = rl.NewRectangle(-200, 400, 3000, 10)

	if rl.CheckCollisionRecs(playerDest, deathColliderRec) {
		rl.StopMusicStream(music)
		rl.PlaySound(deathSFX)
		playerDest.Y = float32(deathColliderRec.Y+50) - playerDest.Height
		playerGrounded = true
		playerJumping = false
		velocityY = 0
		canMoveLeft = false
		canMoveRight = false
		gameOver = true
	}
}

func drawColliders() {

	//DRAW STATIC ITEMS, CONTAINS STATIC COLLIDER LOGIC.
	for _, current_StaticItem := range staticItems {
		rl.DrawRectangle(current_StaticItem.posX, current_StaticItem.posY, current_StaticItem.width, current_StaticItem.height, current_StaticItem.Color)

		rl.DrawCircle(current_StaticItem.posX, current_StaticItem.posY, 2, debugColorYellow)
		rl.DrawCircle(current_StaticItem.posX+current_StaticItem.width, current_StaticItem.posY, 2, debugColorPurple)
		rl.DrawCircle(current_StaticItem.posX, current_StaticItem.posY+current_StaticItem.height, 2, debugColorTeal)
		rl.DrawCircle(current_StaticItem.posX+current_StaticItem.width, current_StaticItem.posY+current_StaticItem.height, 2, debugColor)

		if rl.CheckCollisionRecs(playerDest, rl.NewRectangle(float32(current_StaticItem.posX), float32(current_StaticItem.posY), float32(current_StaticItem.width), float32(current_StaticItem.height))) {

			currentPlatformStart = float32(current_StaticItem.posX)
			currentPlatformEnd = float32(current_StaticItem.posX) + float32(current_StaticItem.width)

			var xDistance float32
			var yDistance float32

			xDistance, yDistance = CalculateAABBDistanceTo(current_StaticItem)

			var xAxisTimeToCollide float32 = float32(math.Abs(float64(xDistance) / float64(velocityX)))
			var yAxisTimeToCollide float32 = float32(math.Abs(float64(yDistance) / float64(velocityY)))

			if xAxisTimeToCollide < yAxisTimeToCollide {

				fmt.Println(color.Colorize(color.Green, "X"))

				playerJumping = false

				if playerDest.X < float32(current_StaticItem.posX) {
					canMoveRight = false
					playerDest.X = float32(current_StaticItem.posX) - playerDest.Width
					colliderHeight = float32(current_StaticItem.posY)
					// fmt.Println("Colliding on LEFT side")
				} else if playerDest.X > float32(current_StaticItem.posX-current_StaticItem.width) {
					playerDest.X = float32(current_StaticItem.posX) + float32(current_StaticItem.width)
					canMoveLeft = false
					colliderHeight = float32(current_StaticItem.posY)
					// fmt.Println("Colliding on RIGHT side")
				}
			} else {

				fmt.Println(color.Colorize(color.Red, "Y"))

				playerDest.Y = float32(current_StaticItem.posY) - playerDest.Height
				playerGrounded = true
				playerJumping = false
				velocityY = 0
			}
		}
	}

	for _, current_Ground := range grounds {
		rl.DrawRectangle(current_Ground.posX, current_Ground.posY, current_Ground.width, current_Ground.height, current_Ground.Color)

		if rl.CheckCollisionRecs(playerDest, rl.NewRectangle(float32(current_Ground.posX), float32(current_Ground.posY),float32(current_Ground.width), float32(current_Ground.height))) {

			currentPlatformStart = float32(current_Ground.posX)
			currentPlatformEnd = float32(current_Ground.posX) + float32(current_Ground.width)

			var xDistance float32
			var yDistance float32

			xDistance, yDistance = CalculateAABBDistanceTo(current_Ground)

			var xAxisTimeToCollide float32 = float32(math.Abs(float64(xDistance) / float64(velocityX)))
			var yAxisTimeToCollide float32 = float32(math.Abs(float64(yDistance) / float64(velocityY)))

			if xAxisTimeToCollide < yAxisTimeToCollide {

				playerJumping = false

				if playerDest.X < float32(current_Ground.posX) {
					canMoveRight = false
					playerDest.X = float32(current_Ground.posX) - playerDest.Width
					colliderHeight = float32(current_Ground.posY)
				} else if playerDest.X > float32(current_Ground.posX-current_Ground.width) {
					playerDest.X = float32(current_Ground.posX) + float32(current_Ground.width)
					canMoveLeft = false
				}
			} else {
				playerDest.Y = float32(current_Ground.posY) - playerDest.Height
				playerGrounded = true
				playerJumping = false
				velocityY = 0
			}
		}
	}

	for _, current_Block := range blocks {

		var blockSrc = rl.NewRectangle(float32(current_Block.posX), float32(current_Block.posY), float32(current_Block.width), float32(current_Block.height))
		var blockDest = rl.NewRectangle(float32(current_Block.posX), float32(current_Block.posY), float32(current_Block.width), float32(current_Block.height))

		blockSrc.X = 0
		blockSrc.Y = 2

		rl.DrawRectangle(current_Block.posX, current_Block.posY, current_Block.width, current_Block.height, current_Block.Color)
		rl.DrawTexturePro(blockSprite, blockSrc, blockDest, rl.NewVector2(0, 0), 0, rl.White)

		if rl.CheckCollisionRecs(playerDest, rl.NewRectangle(float32(current_Block.posX), float32(current_Block.posY), float32(current_Block.width), float32(current_Block.height))) {

			var xDistance float32
			var yDistance float32

			currentPlatformStart = float32(current_Block.posX)
			currentPlatformEnd = float32(current_Block.posX) + float32(current_Block.width)

			xDistance, yDistance = CalculateAABBDistanceTo(current_Block)

			var xAxisTimeToCollide float32 = float32(math.Abs(float64(xDistance) / float64(velocityX)))
			var yAxisTimeToCollide float32 = float32(math.Abs(float64(yDistance) / float64(velocityY)))

			fmt.Println("BLOCK", xAxisTimeToCollide, yAxisTimeToCollide)

			if xAxisTimeToCollide < yAxisTimeToCollide {

				fmt.Println(color.Colorize(color.Green, "X"))

				playerJumping = false

				if playerDest.X < float32(current_Block.posX) {
					playerDest.X = float32(current_Block.posX) - playerDest.Width
					playerJumping = false
					fmt.Println("Triggered LEFT")
				} else if playerDest.X > float32(current_Block.posX-current_Block.width) {
					playerJumping = false
					fmt.Println("Triggered RIGHT")
				}
			} else {

				fmt.Println(color.Colorize(color.Red, "Y"))

				if playerDest.Y < float32(current_Block.posY) {
					playerDest.Y = float32(current_Block.posY) - playerDest.Height
					playerGrounded = true
					playerJumping = false
					velocityY = 0

				} else {
					playerDest.Y = float32(current_Block.posY) + float32(current_Block.height)
					playerJumping = false
					velocityY = 0
					rl.PlaySound(bumpSFX)
				}
			}

		}
	}

	for _, current_coinBlockHit := range coinBlocksHit {
		rl.DrawTexture(coinBlockHitSprite, current_coinBlockHit.posX, current_coinBlockHit.posY-2, rl.White)
	}

	for i, current_coinBlock := range coinBlocks {

		rl.DrawRectangle(current_coinBlock.posX, current_coinBlock.posY, 
			current_coinBlock.width, current_coinBlock.height, 
			current_coinBlock.Color)

		rl.DrawTexture(coinBlockSprite, current_coinBlock.posX, 
			current_coinBlock.posY-3, rl.White)

		if rl.CheckCollisionRecs(playerDest, rl.NewRectangle(float32(current_coinBlock.posX), float32(current_coinBlock.posY),
			float32(current_coinBlock.width), float32(current_coinBlock.height))) {

			currentPlatformStart = float32(current_coinBlock.posX)
			currentPlatformEnd = float32(current_coinBlock.posX) + float32(current_coinBlock.width)

			var xDistance float32
			var yDistance float32

			xDistance, yDistance = CalculateAABBDistanceTo(current_coinBlock)

			var xAxisTimeToCollide float32 = float32(math.Abs(float64(xDistance) / float64(velocityX)))
			var yAxisTimeToCollide float32 = float32(math.Abs(float64(yDistance) / float64(velocityY)))

			if xAxisTimeToCollide < yAxisTimeToCollide {
				if playerDest.X < float32(current_coinBlock.posX) {
					playerDest.X = float32(current_coinBlock.posX) - playerDest.Width
				} else if playerDest.X > float32(current_coinBlock.posX-current_coinBlock.width) {
					playerDest.X = float32(current_coinBlock.posX) + float32(current_coinBlock.width)
				}
			} else {

				if playerDest.Y < float32(current_coinBlock.posY) {
					playerDest.Y = float32(current_coinBlock.posY) - playerDest.Height
					playerGrounded = true
					velocityY = 0
				} else if playerDest.Y > float32(current_coinBlock.posY) && playerJumping {
					playerDest.Y = float32(current_coinBlock.posY) + float32(current_coinBlock.height)
					playerJumping = false
					velocityY = 0
					fmt.Println("HIT COIN BLOCK")
					coinCount++
					fmt.Println(coinCount)
					coinBlocks = append(coinBlocks[:i], coinBlocks[i+1:]...)
					rl.PlaySound(coinSFX)
				} else {
					playerDest.Y = float32(current_coinBlock.posY) + float32(current_coinBlock.height)
					velocityY = 0
				}
				playerJumping = false
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
