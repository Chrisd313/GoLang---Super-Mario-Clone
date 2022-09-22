package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Enemy struct {
	X             float32
	Y             float32
	width         float32
	height        float32
	enemyGrounded bool
	// goombaDest    rl.Rectangle
	// goombaNum     float32
	// goombaSrc     rl.Rectangle
	// position      rl.Vector2
}

var (
	goomba = []Enemy{
		{250, 0, 16, 16, false},
		// {260, 235, 16, 16, false},

		// {240, 235, 16, 16, false, rl.NewRectangle(240, 235, 16, 16)},
		// {246, 220, 16, 16, false, rl.NewRectangle(246, 220, 16, 16)},
		// {220, 220, 16, 16, false, rl.NewRectangle(220, 220, 16, 16)},

		// {240, 235, 16, 16, false, rl.NewRectangle(240, 235, 16, 16), 1},
		// {246, 220, 16, 16, false, rl.NewRectangle(246, 220, 16, 16), 2},
		// {220, 220, 16, 16, false, rl.NewRectangle(235, 225, 16, 16), 3},
	}

	// goomba2 = []Enemy{
	// 	{235, 235, 16, 16, false, rl.NewRectangle(235, 235, 16, 16), 3},
	// }

	// enemyGravity   float32 = 0.5
	enemyVelocityY float32 = 2
	enemyVelocityX float32 = 3
)

// func drawEnemies() {
// 	for _, current_Goomba := range goomba {
// 		var goombaSrc = rl.NewRectangle(float32(current_Goomba.X), float32(current_Goomba.Y), float32(current_Goomba.width), float32(current_Goomba.height))
// 		goombaDest = rl.NewRectangle(float32(current_Goomba.X), float32(current_Goomba.Y), float32(current_Goomba.width), float32(current_Goomba.height))
// 		goombaSrc.X = 0
// 		goombaSrc.Y = 0
// 		rl.DrawTexturePro(goombaSprite, goombaSrc, goombaDest, rl.NewVector2(0, 0), 0, rl.White)

// 		fmt.Println("GOOMBA ", current_Goomba.goombaNum)

// 		current_Goomba.goombaDest.Y += enemyVelocityY

// 		if !current_Goomba.enemyGrounded {
// 			enemyVelocityY += enemyGravity
// 			current_Goomba.goombaDest.Y += enemyVelocityY
// 		} else {
// 			velocityY = 0
// 		}
// 	}
// }

func drawEnemies() {
	for _, current_Goomba := range goomba {

		// fmt.Println("CALLED")
		// fmt.Println(current_Goomba)
		rl.DrawRectangle(int32(current_Goomba.X), int32(current_Goomba.Y), int32(current_Goomba.width), int32(current_Goomba.height), rl.Blue)
		rl.DrawRectangle(goombaDest.ToInt32().X, goombaDest.ToInt32().Y, goombaDest.ToInt32().Width, goombaDest.ToInt32().Height, colliderColor3)

		var goombaSrc = rl.NewRectangle(float32(current_Goomba.X), float32(current_Goomba.Y), float32(current_Goomba.width), float32(current_Goomba.height))
		goombaDest = rl.NewRectangle(float32(current_Goomba.X), float32(current_Goomba.Y), float32(current_Goomba.width), float32(current_Goomba.height))
		goombaSrc.X = 0
		goombaSrc.Y = 0

		// if !current_Goomba.enemyGrounded {
		// 	enemyVelocityY += enemyGravity
		// 	goombaDest.Y += enemyVelocityY
		// }

		// for _, element := range grounds {

		// 	if rl.CheckCollisionRecs(goombaDest, rl.NewRectangle(float32(element.posX), float32(element.posY), float32(element.width), float32(element.height))) {
		// 		current_Goomba.enemyGrounded = true
		// 		goombaDest.Y = float32(element.posY) - current_Goomba.height
		// 	}
		// }

		// fmt.Println(current_Goomba.goombaDest.X, current_Goomba.goombaDest.Y, i)

		// fmt.Println("COLLISION | Grounded: ", current_Goomba.enemyGrounded, " | Object Y: ", current_Goomba.Y, " | Dest Y: ", goombaDest.Y)
		rl.DrawTexturePro(goombaSprite, goombaSrc, goombaDest, rl.NewVector2(0, 0), 0, rl.White)
	}
}

func updateEnemies() {

	for _, current_Goomba := range goomba {

		// if !current_Goomba.enemyGrounded {
		// 	enemyVelocityY += enemyGravity
		// 	goombaDest.Y += enemyVelocityY
		// }

		enemyVelocityX += enemyGravity
		goombaDest.X += enemyVelocityX

		for _, element := range staticItems {

			if rl.CheckCollisionRecs(goombaDest, rl.NewRectangle(float32(element.posX), float32(element.posY), float32(element.width), float32(element.height))) {
				fmt.Println("Collision", current_Goomba)
			}
		}

		// for _, element := range grounds {

		// 	if rl.CheckCollisionRecs(goombaDest, rl.NewRectangle(float32(element.posX), float32(element.posY), float32(element.width), float32(element.height))) && !current_Goomba.enemyGrounded {
		// 		current_Goomba.enemyGrounded = true
		// 		goombaDest.Y = float32(element.posY) - current_Goomba.height
		// 		// fmt.Println("Collision with floor")
		// 	}
		// }

		// fmt.Println("Dest: ", goombaDest.X, goombaDest.Y, " | Current: ", current_Goomba.X, current_Goomba.Y)

	}

}
