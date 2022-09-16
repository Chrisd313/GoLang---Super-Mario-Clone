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
	playerAxis     = rl.NewColor(10, 255, 10, 255)
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

	xInvEntry float32
	yInvEntry float32
	xInvExit  float32
	yInvExit  float32

	musicPaused bool
	music       rl.Music

	cam rl.Camera2D
)

func drawScene() {
	rl.DrawTexture(bgSprite, 0, 58, rl.White)

	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(0, 0), 0, rl.White)
	rl.DrawRectangle(playerDest.ToInt32().X, playerDest.ToInt32().Y, playerDest.ToInt32().Width, playerDest.ToInt32().Height, colliderColor2)

	// rl.DrawLine(int32(playerDest.X), int32(playerDest.Y), playerDest.ToInt32().X+playerDest.ToInt32().Width, playerDest.ToInt32().Y, playerAxis)

	drawColliders()
}

func update() {
	running = !rl.WindowShouldClose()
	playerSrc.X = 0

	// fmt.Println("Velocity X: ", velocityX, " | Velocity Y: ", velocityY)

	// velocityY += gravity
	if !playerGrounded {
		velocityY += gravity
		playerDest.Y += velocityY
		// playerCollider.Y = float64(playerDest.Y)
	} else {
		velocityY = 0
	}

	// HORIZONTAL MOVEMENT
	if playerMoving {
		if playerLeft {
			playerDest.X -= velocityX
			playerCollider.X = float64(playerDest.X)
			// fmt.Println(playerDest.X, playerCollider.X)
		}
		if playerRight && canMoveRight {
			playerDest.X += velocityX
			playerCollider.X = float64(playerDest.X)
		}
		if frameCount%8 == 1 {
			playerFrame++
		}
		playerSrc.X = playerSrc.Width * float32(playerFrame)
	}

	// playerDest.Y += velocityY * rl.GetFrameTime()

	frameCount++
	if !playerJumping && playerFrame > 3 {
		playerFrame = 0
	} else if playerJumping {
		playerFrame = 5
	} else if !playerMoving {
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

	// //Collision detection
	// dx := float64(playerDest.X)
	// dy := float64(playerDest.Y)

	// if collision := playerCollider.Check(dx, 0, "pipeTag"); collision != nil {
	// 	fmt.Println("Pipe collision detected, dx: ", dx)

	// 	// if playerDir == 1 {
	// 	// 	dx = float64(collision.ContactWithObject(collision.Objects[0]).X())
	// 	// } else {
	// 	// 	dx = float64(collision.ContactWithObject(collision.Objects[0]).X()) * 2
	// 	// }
	// 	// playerDest.X += float32(dx64)
	// 	// playerMoving = false

	// 	// fmt.Println("dx64:", dx64)
	// 	// fmt.Println(collision.Objects)

	// 	// if playerDir == 1 {
	// 	// 	canMoveRight = true
	// 	// 	canMoveLeft = false
	// 	// } else if playerDir == 0 {
	// 	// 	canMoveRight = false
	// 	// 	canMoveLeft = true
	// 	// }
	// }

	// if collision := playerCollider.Check(dx, 32, "pipeTag"); collision != nil {
	// 	fmt.Println("Pipe collision detected, dy: ", dy)
	// 	objectX := collision.ContactWithObject(collision.Objects[0])
	// 	playerGrounded = true
	// 	playerJumping = false
	// 	fmt.Println("The objects X: ", objectX.X(), " | The objects Y: ", objectX.Y())
	// 	// fmt.Println(color.Colorize(color.Red, "JUMP END"))
	// 	// velocityY = 0
	// }

	// if collision := playerCollider.Check(dy, 0, "groundTag"); collision != nil {
	// 	// fmt.Println("Floor collision detected")
	// 	// playerDest.Y = float32(dy64)
	// 	playerGrounded = true
	// 	playerJumping = false
	// 	// fmt.Println(color.Colorize(color.Red, "JUMP END"))
	// 	// velocityY = 0

	// }

	// if collision := playerCollider.Check(dy, 0, "blockTag"); collision != nil {
	// 	fmt.Println("Block collision detected")
	// 	playerGrounded = true
	// 	playerJumping = false
	// 	// fmt.Println(color.Colorize(color.Red, "JUMP END"))
	// 	// velocityY = 0
	// }

	// playerCollider.X = float64(playerDest.X)
	// playerCollider.Y = float64(playerDest.Y)
	// playerCollider.Update()
	// // playerCollider.X += float32(playerCollider.X)

	// // fmt.Print("playerDest X/Y: ", playerDest.X, playerDest.Y, "| dx/dy: ", dx, dy, "| playerCollider X/Y: ", playerCollider.X, playerCollider.Y, "\n playerDest W/H: ", playerDest.Width, playerDest.Height, " | playerCollider W/H: ", playerCollider.W, playerCollider.H, "\n")

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
