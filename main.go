package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 512
	screenHeight = 460
)

var (
	debugColliderAlpha int = 0 // Alpha color control for debug colliders, switch to zero if want to turn off debugging.
	running                = true
	bkgColor               = rl.NewColor(147, 211, 196, 255)
	colliderColor          = rl.NewColor(15, 10, 222, uint8(debugColliderAlpha))
	colliderColor2         = rl.NewColor(255, 10, 10, uint8(debugColliderAlpha))
	colliderColor3         = rl.NewColor(0, 255, 255, uint8(debugColliderAlpha))

	bgSprite           rl.Texture2D
	playerSprite       rl.Texture2D
	blockSprite        rl.Texture2D
	coinBlockSprite    rl.Texture2D
	coinBlockHitSprite rl.Texture2D

	frameCount int

	velocityX float32 = 3
	velocityY float32 = 0
	gravity   float32 = 0.5

	musicPaused bool
	music       rl.Music

	deathSFX rl.Sound
	bumpSFX  rl.Sound
	coinSFX  rl.Sound

	coinCount int

	cam      rl.Camera2D
	gameOver bool = false
	camera   rl.Camera2D
)

func drawScene() {
	rl.DrawTexture(bgSprite, 0, 58, rl.White)

	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(0, 0), 0, rl.White)
	rl.DrawRectangle(playerDest.ToInt32().X, playerDest.ToInt32().Y, playerDest.ToInt32().Width, playerDest.ToInt32().Height, colliderColor2)

	/* #region Collider visualisation */
	rl.DrawCircle(playerDest.ToInt32().X, playerDest.ToInt32().Y, 2, debugColorYellow)
	rl.DrawCircle(playerDest.ToInt32().X+int32(playerDest.Width), playerDest.ToInt32().Y, 2, debugColorPurple)
	rl.DrawCircle(playerDest.ToInt32().X, playerDest.ToInt32().Y+playerDest.ToInt32().Height, 2, debugColorTeal)
	rl.DrawCircle(playerDest.ToInt32().X+playerDest.ToInt32().Width, playerDest.ToInt32().Y+int32(playerDest.Height), 2, debugColor)
	/* #endregion */

	drawColliders()
	deathCollider()
}

func update() {
	// fmt.Println(velocityY)

	running = !rl.WindowShouldClose()

	if playerDest.X > currentPlatformEnd || playerDest.X+playerDest.Width < currentPlatformStart {
		playerGrounded = false
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

	playerSrc.Y = playerSrc.Height * float32(playerDir)

	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	//CAMERA SETTINGS AND POSITIONING
	cam.Target = rl.NewVector2(float32(playerDest.X+(playerDest.Width*2)), float32(258))
	camera.Target = rl.NewVector2(float32(playerDest.X+(playerDest.Width*2)), float32(258))
	camera.Offset = rl.NewVector2(screenWidth/2, 380)

	var minX float32 = 0
	var maxX float32 = 3376
	var minY float32 = 266
	var maxY float32 = 480

	var max rl.Vector2 = rl.GetWorldToScreen2D(rl.Vector2{maxX, maxY}, camera)
	var min rl.Vector2 = rl.GetWorldToScreen2D(rl.Vector2{minX, minY}, camera)

	if max.X < screenWidth {
		camera.Offset.X = (screenWidth - (max.X - screenWidth/2))
	}
	// if max.Y < screenHeight {
	// 	camera.Offset.Y = screenHeight - (max.Y - screenHeight/2)
	// }
	if min.X > 0 {
		camera.Offset.X = screenHeight/2 - min.X
	}
	// if min.Y > 0 {
	// 	camera.Offset.Y = screenHeight - (screenHeight/2 - min.Y)
	// }

	playerMoving = false
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgColor)
	rl.BeginMode2D(camera)
	drawScene()

	rl.EndMode2D()

	rl.EndDrawing()

}

func initFunc() {
	rl.InitWindow(screenWidth, screenHeight, "Super Mario Bros. - Go")
	rl.SetExitKey(rl.KeyEscape)
	rl.SetTargetFPS(60)

	bgSprite = rl.LoadTexture("assets/images/bg1.png")
	playerSprite = rl.LoadTexture("assets/images/mario.png")
	blockSprite = rl.LoadTexture("assets/images/block.png")
	coinBlockSprite = rl.LoadTexture("assets/images/coinBlock.png")
	coinBlockHitSprite = rl.LoadTexture("assets/images/coinBlockHit.png")

	playerSrc = rl.NewRectangle(0, 0, 16, 16)
	playerDest = rl.NewRectangle(200, 200, 16, 16)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("assets/music/01_Running_About.mp3")
	deathSFX = rl.LoadSound("assets/sfx/death.wav")
	bumpSFX = rl.LoadSound("assets/sfx/bump.wav")
	coinSFX = rl.LoadSound("assets/sfx/coin.wav")
	musicPaused = false
	rl.PlayMusicStream(music)

	camera = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(258)), rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(258)), 0.0, 2.0)
}

func quit() {
	rl.UnloadTexture(bgSprite)
	rl.UnloadTexture(playerSprite)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func main() {

	initFunc()

	for running {
		input()
		update()
		render()
	}

	quit()
}
