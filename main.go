package main

import (
	"fmt"

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

type Collider struct {
	posX   int32
	posY   int32
	width  int32
	height int32
	Color  rl.Color
}

var (
	running       = true
	bkgColor      = rl.NewColor(147, 211, 196, 255)
	colliderColor = rl.NewColor(15, 10, 222, 100)
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
	playerJumping			bool = false
	jumpVelocity	int = 0
	maxJumpHeight	int = 10

	frameCount int

	playerSpeed float32 = 3

	musicPaused bool
	music       rl.Music

	cam   rl.Camera2D
	pipes = []Collider{
		// {0, 266, 1104, 20, colliderColor},
		{456, 234, 16, 32, colliderColor},
		{608, 218, 32, 48, colliderColor},
	}

	grounds = []Collider{
		{0, 266, 1104, 20, colliderColor},
	}
)

func drawScene() {
	rl.DrawTexture(bgSprite, 0, 58, rl.White)

	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)

	for _, current_Pipe := range pipes {
		rl.DrawRectangle(current_Pipe.posX, current_Pipe.posY, current_Pipe.width, current_Pipe.height, current_Pipe.Color)
	}

	for _, current_Ground := range grounds {
		rl.DrawRectangle(current_Ground.posX, current_Ground.posY, current_Ground.width, current_Ground.height, current_Ground.Color)
	}
}

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

	if rl.IsKeyDown(rl.KeyP) && playerGrounded || rl.IsKeyDown(rl.KeySpace) && playerGrounded{
		playerJumping = true
		playerGrounded = false
		// fmt.Println("Jump!")
		for jumpVelocity < maxJumpHeight {
			if jumpVelocity == maxJumpHeight{
				jumpVelocity = 0
			} else {
				jumpVelocity += 1
				playerDest.Y -= float32(jumpVelocity)
				fmt.Println(jumpVelocity)
			}
		}

	}

	if rl.IsKeyDown(rl.KeyR){
		main()
	}
}

func update() {
	running = !rl.WindowShouldClose()

	if !playerGrounded {
		playerDest.Y += 2
	}

	// fmt.Println(playerDest.X, playerObj.X)
	// fmt.Println("coll width:", playerObj, "player:", playerDest)

	playerSrc.X = 0

	if playerMoving {
		if playerLeft {
			playerDest.X -= playerSpeed
			playerObj.X = float64(playerDest.X)
			// fmt.Println(playerDest.X, playerObj.X)
		}
		if playerRight && canMoveRight {
			playerDest.X += playerSpeed
			playerObj.X = float64(playerDest.X)
		}
		if frameCount%8 == 1 {
			playerFrame++
		}
		playerSrc.X = playerSrc.Width * float32(playerFrame)
	}

	frameCount++
	if playerFrame > 3 {
		playerFrame = 0
	}

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

	//Collision detection
	dx := float64(playerDest.X)
	dy := 2.0
	// dx := 1.0

	if collision := playerObj.Check(dx, 0, "pipeTag"); collision != nil {
		fmt.Println("Pipe collision detected, dx: ", dx)

		dx64 := float64(collision.ContactWithObject(collision.Objects[0]).X())
		playerDest.X = float32(dx64)
		playerMoving = false

		if playerDir == 1 {
			canMoveRight = true
			canMoveLeft = false
		} else if playerDir == 0 {
			canMoveRight = false
			canMoveLeft = true
		}
	}

	if collision := playerObj.Check(dy, 0, "groundTag"); collision != nil {
		// fmt.Println("Floor collision detected")
		// playerDest.Y = float32(dy64)
		playerGrounded = true
		playerJumping = false
		jumpVelocity = 0
	}

	playerObj.X = float64(playerDest.X)
	playerObj.Y = float64(playerDest.Y)
	playerObj.Update()
	// playerCollider.X += float32(playerObj.X)
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

	// pipes := []Pipe{}

	// pipes = append(pipes, Pipe)
	// pipes = append(pipes, pipe1)

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
