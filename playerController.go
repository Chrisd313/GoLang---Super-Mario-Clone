package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	playerSpeed        float32 = 0
	playerMaxSpeed     float32 = 120
	playerAcceleration float32 = 200
	// jumpSpeed float32 = 2

	playerSrc    rl.Rectangle
	playerDest   rl.Rectangle
	playerMoving bool
	playerDir    int
	// playerRight, playerLeft bool
	playerFrame    int
	canMoveRight   bool = true
	canMoveLeft    bool = true
	playerGrounded bool = false
	playerJumping  bool = false
)

func input() {

	playerSrc.X = 0

	if rl.IsKeyDown(rl.KeyA) && canMoveLeft {
		playerMoving = true
		playerDir = 1
		canMoveRight = true
		playerDest.X -= velocityX
	} else if rl.IsKeyDown(rl.KeyD) && canMoveRight {
		playerMoving = true
		playerDir = 0
		canMoveLeft = true
		playerDest.X += velocityX
	} else if rl.IsKeyDown(rl.KeyLeft) && canMoveLeft {
		playerMoving = true
		playerDir = 1
		canMoveRight = true
		if playerSpeed > -playerMaxSpeed {
			playerSpeed = playerSpeed - (playerAcceleration * rl.GetFrameTime())
		}
		if playerSpeed > 0 && !playerJumping {
			playerFrame = 4
		}
	} else if rl.IsKeyDown(rl.KeyRight) && canMoveRight {
		playerMoving = true
		playerDir = 0
		canMoveLeft = true
		if playerSpeed < playerMaxSpeed {
			playerSpeed = playerSpeed + (playerAcceleration * rl.GetFrameTime())
		}
		if playerSpeed < 0 && !playerJumping {
			playerFrame = 4
		}
	} else {
		// fmt.Println(playerSpeed)
		if playerSpeed > 1.5 {
			playerSpeed = playerSpeed - (playerAcceleration * rl.GetFrameTime())
		} else if playerSpeed < -1.5 {
			playerSpeed = playerSpeed + (playerAcceleration * rl.GetFrameTime())
		} else {
			playerSpeed = 0
		}
		playerFrame = 0
	}

	playerSrc.X = playerSrc.Width * float32(playerFrame)

	playerDest.X = playerDest.X + playerSpeed*rl.GetFrameTime()

	if rl.IsKeyPressed(rl.KeyQ) {
		musicPaused = !musicPaused
	}

	if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyUp) {
		startJump()
	}

	if rl.IsKeyReleased(rl.KeySpace) || rl.IsKeyReleased(rl.KeyUp) {
		endJump()
	}

	if rl.IsKeyPressed(rl.KeyEnter) && gameOver {
		playerDest = rl.NewRectangle(200, 200, 16, 16)
		canMoveLeft = true
		canMoveRight = true
		gameOver = false
		rl.PlayMusicStream(music)
		drawScene()
		rl.StopAudioStream(deathSFX.Stream)
	}

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
	if playerGrounded {
		velocityY = -9
		playerGrounded = false
		playerJumping = true
		var jumpSFX = rl.LoadSound("assets/sfx/jump.wav")
		rl.PlaySound(jumpSFX)
	}
}

func endJump() {
	if velocityY < -5.0 {
		velocityY = -5.0
	}
}
