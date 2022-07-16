package main

import rl "github.com/gen2brain/raylib-go/raylib"

func pause_update(GS *GameState, DeltaTime float32) {
	if rl.IsKeyPressed(rl.KeyS) || rl.IsKeyPressed(rl.KeyDown) {
		GS.SelectedPauseMenuOption += 1
		if GS.SelectedPauseMenuOption >= len(GS.PauseMenuOptions) {
			GS.SelectedPauseMenuOption = 0
		}
	}
	if rl.IsKeyPressed(rl.KeyW) || rl.IsKeyPressed(rl.KeyUp) {
		GS.SelectedPauseMenuOption -= 1
		if GS.SelectedPauseMenuOption < 0 {
			GS.SelectedPauseMenuOption = len(GS.PauseMenuOptions) - 1
		}
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		GS.PauseMenuOptions[GS.SelectedPauseMenuOption].Callback(GS)
	}
}

func pause_render(GS *GameState) {
	draw_paddles_ball_and_score(GS)

	// Menu options
	{
		const OptionFontSize int32 = 45
		var YPosition int32 = WindowHeight / 2
		const DeltaY int32 = OptionFontSize * 2
		for i := 0; i < len(GS.PauseMenuOptions); i++ {
			var TitleWidth int32 = rl.MeasureText(GS.PauseMenuOptions[i].Name, OptionFontSize)
			var OptionColor rl.Color = rl.Gray
			if GS.SelectedPauseMenuOption == i {
				OptionColor = rl.White
			}
			rl.DrawText(GS.PauseMenuOptions[i].Name, (WindowWidth-TitleWidth)/2, YPosition, OptionFontSize, OptionColor)
			YPosition += DeltaY
		}
	}
}

func on_continue(GS *GameState) {
	change_to_playing(GS)
}

func on_return_to_main_menu(GS *GameState) {
	restart_game(GS)
	change_to_menu(GS)
}
