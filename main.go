package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 512
	screenHeight = 460
)

//PLAYER JUMP
/*

add var canJump as a bool

*/

var (
	running        = true
	bkgColor       = rl.NewColor(147, 211, 196, 255)
	colliderColor  = rl.NewColor(15, 10, 222, 100)
	colliderColor2 = rl.NewColor(255, 10, 10, 100)
	colliderColor3 = rl.NewColor(0, 255, 255, 100)
	// fontColor = rl.NewColor(221, 89, 24, 255)

	bgSprite     rl.Texture2D
	playerSprite rl.Texture2D

	playerSrc               rl.Rectangle
	playerDest              rl.Rectangle
	playerMoving            bool
	playerDir               int
	playerRight, playerLeft bool
	playerFrame             int
	canMoveRight            bool = true
	canMoveLeft             bool = true
	playerGrounded          bool = false
	playerJumping           bool = false

	frameCount int

	velocityX float32 = 3
	velocityY float32 = 0
	gravity   float32 = 0.5

	musicPaused bool
	music       rl.Music

	coinCount int

	cam rl.Camera2D
)

func drawScene() {
	rl.DrawTexture(bgSprite, 0, 58, rl.White)

	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(0, 0), 0, rl.White)
	rl.DrawRectangle(playerDest.ToInt32().X, playerDest.ToInt32().Y, playerDest.ToInt32().Width, playerDest.ToInt32().Height, colliderColor2)

	rl.DrawCircle(playerDest.ToInt32().X, playerDest.ToInt32().Y, 2, debugColorYellow)
	rl.DrawCircle(playerDest.ToInt32().X+int32(playerDest.Width), playerDest.ToInt32().Y, 2, debugColorPurple)
	rl.DrawCircle(playerDest.ToInt32().X, playerDest.ToInt32().Y+playerDest.ToInt32().Height, 2, debugColorTeal)
	rl.DrawCircle(playerDest.ToInt32().X+playerDest.ToInt32().Width, playerDest.ToInt32().Y+int32(playerDest.Height), 2, debugColor)

	drawColliders()
}

func update() {

	running = !rl.WindowShouldClose()
	playerSrc.X = 0

	if playerDest.X > currentPlatformEnd || playerDest.X < currentPlatformStart {
		playerGrounded = false
		velocityX = 3
	}

	if !playerGrounded {
		velocityY += gravity
		playerDest.Y += velocityY
	} else {
		velocityY = 0
	}

	if playerDest.Y+playerDest.Height < colliderHeight {
		canMoveLeft = true
		canMoveRight = true
	}

	playerCollider.X = float64(playerDest.X)
	playerCollider.Y = float64(playerDest.Y)

	// HORIZONTAL MOVEMENT
	if playerMoving {
		if playerLeft {
			playerDest.X -= velocityX
			// playerCollider.X = float64(playerDest.X)
		}
		if playerRight && canMoveRight {
			playerDest.X += velocityX
			// playerCollider.X = float64(playerDest.X)
		}
		if !playerJumping && frameCount%8 == 1 {
			playerFrame++
		} else if playerJumping {
			playerFrame = 5
		}
		playerSrc.X = playerSrc.Width * float32(playerFrame)
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

	// if !playerJumping && playerFrame > 3 {
	// 	playerFrame = 0
	// } else if playerJumping {
	// 	playerFrame = 5
	// }

	playerSrc.Y = playerSrc.Height * float32(playerDir)

	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	cam.Target = rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(258))

	playerMoving = false
	playerLeft, playerRight = false, false
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgColor)
	rl.BeginMode2D(cam)
	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	bgSprite = rl.LoadTexture("assets/images/bg1.png")
	playerSprite = rl.LoadTexture("assets/images/mario.png")

	playerSrc = rl.NewRectangle(0, 0, 16, 16)
	playerDest = rl.NewRectangle(200, 200, 16, 16)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("assets/music/01_Running_About.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)

	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(380)), rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(200)), 0.0, 2.0)
}

func quit() {
	rl.UnloadTexture(bgSprite)
	rl.UnloadTexture(playerSprite)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func main() {

	for running {
		input()
		update()
		render()
	}

	quit()
}
