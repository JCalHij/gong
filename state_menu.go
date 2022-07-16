package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func menu_update(GS *GameState, DeltaTime float32) {
	if rl.IsKeyPressed(rl.KeyS) || rl.IsKeyPressed(rl.KeyDown) {
		GS.SelectedOption += 1
		if GS.SelectedOption >= len(GS.MenuOptions) {
			GS.SelectedOption = 0
		}
	}
	if rl.IsKeyPressed(rl.KeyW) || rl.IsKeyPressed(rl.KeyUp) {
		GS.SelectedOption -= 1
		if GS.SelectedOption < 0 {
			GS.SelectedOption = len(GS.MenuOptions) - 1
		}
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		GS.MenuOptions[GS.SelectedOption].Callback(GS)
	}
}

func menu_render(GS *GameState) {
	// Title
	{
		const MenuTitle string = "gong"
		const MenuTitleFontSize int32 = 115
		var TitleWidth int32 = rl.MeasureText(MenuTitle, MenuTitleFontSize)
		rl.DrawText(MenuTitle, (WindowWidth-TitleWidth)/2, 10, MenuTitleFontSize, rl.White)
	}

	// Menu options
	{
		const OptionFontSize int32 = 45
		var YPosition int32 = WindowHeight / 2
		const DeltaY int32 = OptionFontSize * 2
		for i := 0; i < len(GS.MenuOptions); i++ {
			var TitleWidth int32 = rl.MeasureText(GS.MenuOptions[i].Name, OptionFontSize)
			var OptionColor rl.Color = rl.Gray
			if GS.SelectedOption == i {
				OptionColor = rl.White
			}
			rl.DrawText(GS.MenuOptions[i].Name, (WindowWidth-TitleWidth)/2, YPosition, OptionFontSize, OptionColor)
			YPosition += DeltaY
		}
	}
}

func on_player_vs_ai(GS *GameState) {
	GS.LeftInput = left_player_input
	GS.RightInput = right_ai_input
	change_to_idle(GS)
}

func on_player_vs_player(GS *GameState) {
	GS.LeftInput = left_player_input
	GS.RightInput = right_player_input
	change_to_idle(GS)
}

func on_quit(GS *GameState) {
	GS.Running = false
}
