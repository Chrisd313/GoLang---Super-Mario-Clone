package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func input() {
	if rl.IsKeyDown(rl.KeyA) && canMoveLeft || rl.IsKeyDown(rl.KeyLeft) && canMoveLeft {
		playerMoving = true
		playerDir = 1
		playerLeft = true
		canMoveRight = true
	}
	if rl.IsKeyDown(rl.KeyD) && canMoveRight || rl.IsKeyDown(rl.KeyRight) && canMoveRight {
		playerMoving = true
		playerDir = 0
		playerRight = true
		canMoveLeft = true
	}

	if rl.IsKeyPressed(rl.KeyQ) {
		musicPaused = !musicPaused
	}

	if rl.IsKeyDown(rl.KeyP) && playerGrounded || rl.IsKeyDown(rl.KeySpace) && !playerJumping {
		startJump()
	}

	if rl.IsKeyReleased(rl.KeyP) || rl.IsKeyReleased(rl.KeySpace) {
		endJump()
	}
}

func startJump() {
	if playerGrounded {
		// fmt.Println(color.Colorize(color.Green, "JUMP START"))
		velocityY = -9.0
		playerGrounded = false
		playerJumping = true

		playerFrame = 5
	}
}

func endJump() {
	if velocityY < -5.0 {
		velocityY = -5.0
	}
}
