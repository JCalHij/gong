package main

import (
	"math/rand"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

//TODO[javi]: In-game pause menu -> Continue, Main Menu, Quit
//TODO[javi]: sound effects

/* Game state updates */

func init_game() GameState {
	return GameState{
		Update:  menu_update,
		Render:  menu_render,
		Running: true,

		LeftPaddle:              InitialLeftPaddle,
		RightPaddle:             InitialRightPaddle,
		Ball:                    InitialBall,
		BallDirection:           vec2_from_angle(Random.Float64()),
		LeftScore:               0,
		RightScore:              0,
		LeftInput:               left_player_input,
		LeftPlayerHuman:         true,
		RightInput:              right_ai_input,
		RightPlayerHuman:        false,
		SelectedPauseMenuOption: 0,
		PauseMenuOptions: [2]MenuOptionData{
			{Name: "Continue", Callback: on_continue},
			{Name: "Quit", Callback: on_return_to_main_menu},
		},

		SelectedMainMenuOption: 0,
		MainMenuOptions: [3]MenuOptionData{
			{Name: "Player vs AI", Callback: on_player_vs_ai},
			{Name: "Player vs Player", Callback: on_player_vs_player},
			{Name: "Quit", Callback: on_quit},
		}}
}

/* Entry Point */

func main() {
	var RandSource = rand.NewSource(time.Now().UnixNano())
	Random = rand.New(RandSource)
	rl.InitWindow(WindowWidth, WindowHeight, "gong")
	rl.SetExitKey(0)

	//rl.SetTargetFPS(60)
	var GS GameState = init_game()

	for !rl.WindowShouldClose() && GS.Running {
		var DeltaTime float32 = rl.GetFrameTime() // [s] frame time

		/* Game Logic */
		GS.Update(&GS, DeltaTime)

		/* Rendering */
		{
			rl.BeginDrawing()
			rl.ClearBackground(rl.Black)

			GS.Render(&GS)

			// Debugging
			//rl.DrawLine(0, WindowHeight/2, WindowWidth, WindowHeight/2, rl.Gray)
			//rl.DrawLine(WindowWidth/2, 0, WindowWidth/2, WindowHeight, rl.Gray)
			//rl.DrawText(fmt.Sprintf("%.03f ms", DeltaTime*1000), 10, 10, 25, rl.Red)

			rl.EndDrawing()
		}
	}

	rl.CloseWindow()
}
