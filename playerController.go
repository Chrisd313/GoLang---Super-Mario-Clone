package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	playerSpeed        float32 = 0
	playerMaxSpeed     float32 = 120
	playerAcceleration float32 = 100
	movementH          float32 = 0
	movementV          float32 = 0
)

func input() {

	playerSrc.X = 0

	if rl.IsKeyDown(rl.KeyA) && canMoveLeft || rl.IsKeyDown(rl.KeyLeft) && canMoveLeft {
		playerMoving = true
		playerDir = 1
		// playerLeft = true
		// canMoveRight = true
		// playerDest.X -= velocityX
		// fmt.Println(velocityX + (playerMaxSpeed * playerAcceleration * rl.GetFrameTime()))
		if playerSpeed > -playerMaxSpeed {
			playerSpeed = playerSpeed - (playerAcceleration * rl.GetFrameTime())
		}

	} else if rl.IsKeyDown(rl.KeyD) && canMoveRight || rl.IsKeyDown(rl.KeyRight) && canMoveRight {
		playerMoving = true
		playerDir = 0
		// playerRight = true
		// canMoveLeft = true

		// // velocityX += playerAcceleration * rl.GetFrameTime()
		// fmt.Println(velocityX + (playerMoveSpeed*playerAcceleration*rl.GetFrameTime()))
		// playerDest.X += velocityX
		if playerSpeed < playerMaxSpeed {
			playerSpeed = playerSpeed + (playerAcceleration * rl.GetFrameTime())
		}
	} else {
		playerSpeed = 0
		playerMoving = false
		playerFrame = 0

	}

	if rl.IsKeyReleased(rl.KeyA) || rl.IsKeyReleased(rl.KeyLeft) || rl.IsKeyReleased(rl.KeyD) || rl.IsKeyReleased(rl.KeyRight) {
		playerSpeed = 0
		playerMoving = false
		playerFrame = 0
	}

	playerDest.X = playerDest.X + playerSpeed*rl.GetFrameTime()

	if rl.IsKeyPressed(rl.KeyQ) {
		musicPaused = !musicPaused
	}

	if rl.IsKeyPressed(rl.KeyP) && playerGrounded || rl.IsKeyPressed(rl.KeySpace) && playerGrounded || rl.IsKeyPressed(rl.KeyUp) && playerGrounded {

		playerSpeed = playerSpeed - (playerAcceleration * rl.GetFrameTime())

		startJump()
	}

	if rl.IsKeyReleased(rl.KeyP) || rl.IsKeyReleased(rl.KeySpace) || rl.IsKeyReleased(rl.KeyUp) {
		endJump()
	}

	if rl.IsKeyPressed(rl.KeyM) {
		debugColliderAlpha = 0
	}

	playerSrc.X = playerSrc.Width * float32(playerFrame)

	if playerMoving {
		if !playerJumping && frameCount%8 == 1 {
			playerFrame++
		} else if playerJumping {
			playerFrame = 5
		}
	} else if !playerMoving {
		if playerJumping {
			playerFrame = 5
			playerSrc.X = playerSrc.Width * float32(playerFrame)
		}
	}

	frameCount++
	if playerFrame > 3 {
		playerFrame = 0
	}
	if !playerJumping && playerFrame > 3 {
		playerFrame = 0
	} else if playerJumping {
		playerFrame = 5
	}
}

func startJump() {
	playerDest.Y = playerDest.Y - playerSpeed*rl.GetFrameTime()

	if playerGrounded {
		// fmt.Println(color.Colorize(color.Green, "JUMP START"))
		velocityY = -9
		playerGrounded = false
		playerJumping = true
		// playerFrame = 5
		var jumpSFX = rl.LoadSound("assets/sfx/jump.wav")
		rl.PlaySound(jumpSFX)
	}
}

func endJump() {
	if velocityY < -5.0 {
		velocityY = -5.0
	}
}
