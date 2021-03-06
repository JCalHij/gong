package main

import rl "github.com/gen2brain/raylib-go/raylib"

func pause_update(GS *GameState, DeltaTime float32) {
	if rl.IsKeyPressed(rl.KeyS) || rl.IsKeyPressed(rl.KeyDown) {
		play_sound(SFX_OptionMove)
		GS.SelectedPauseMenuOption += 1
		if GS.SelectedPauseMenuOption >= len(GS.PauseMenuOptions) {
			GS.SelectedPauseMenuOption = 0
		}
	}
	if rl.IsKeyPressed(rl.KeyW) || rl.IsKeyPressed(rl.KeyUp) {
		play_sound(SFX_OptionMove)
		GS.SelectedPauseMenuOption -= 1
		if GS.SelectedPauseMenuOption < 0 {
			GS.SelectedPauseMenuOption = len(GS.PauseMenuOptions) - 1
		}
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		play_sound(SFX_OptionSelect)
		GS.PauseMenuOptions[GS.SelectedPauseMenuOption].Callback(GS)
	}
}

func pause_render(GS *GameState) {
	draw_paddles_ball_and_score(GS)

	const OptionFontSize int32 = 35
	const BorderMargins = 40
	const DeltaY int32 = OptionFontSize * 2
	var CenterPosition rl.Vector2 = rl.Vector2{
		X: float32(WindowWidth) / 2.0,
		Y: float32(WindowHeight) / 2.0}
	var NumOptions int32 = int32(len(GS.PauseMenuOptions))
	var ContentHeight int32 = NumOptions*OptionFontSize + (NumOptions-1)*DeltaY + 2*BorderMargins

	// Menu Borders
	var ContentWidth float32 = 350.0
	var BorderRectangle rl.Rectangle = rl.Rectangle{
		X:      CenterPosition.X - ContentWidth/2,
		Y:      CenterPosition.Y - float32(ContentHeight)/2.0,
		Width:  ContentWidth,
		Height: float32(ContentHeight)}

	draw_rect(&BorderRectangle, rl.Black)
	rl.DrawRectangleLinesEx(BorderRectangle, 1.0, rl.White)

	// Menu options
	{
		var YPosition int32 = int32(BorderRectangle.Y) + BorderMargins + OptionFontSize/2
		for i := 0; i < len(GS.PauseMenuOptions); i++ {
			var TitleWidth int32 = rl.MeasureText(GS.PauseMenuOptions[i].Name, OptionFontSize)
			var OptionColor rl.Color = rl.Gray
			if GS.SelectedPauseMenuOption == i {
				OptionColor = rl.White
			}
			rl.DrawText(GS.PauseMenuOptions[i].Name, (WindowWidth-TitleWidth)/2, YPosition, OptionFontSize, OptionColor)
			//rl.DrawRectangleLines((WindowWidth-TitleWidth)/2, YPosition, TitleWidth, OptionFontSize, rl.Red)
			YPosition += DeltaY
		}
	}

	// Menu title
	{
		const PauseTitleFontSize int32 = 45
		const PauseTitle string = "PAUSE MENU"
		var PauseTitleWidth int32 = rl.MeasureText(PauseTitle, PauseTitleFontSize)

		var TitleBorderWidth float32 = float32(PauseTitleWidth) * 1.1
		var Diff int32 = int32(TitleBorderWidth-float32(PauseTitleWidth)) / 2

		var TitleRect rl.Rectangle = rl.Rectangle{
			X:      BorderRectangle.X + (BorderRectangle.Width-TitleBorderWidth)/2,
			Y:      BorderRectangle.Y - float32(PauseTitleFontSize)/2.0,
			Width:  TitleBorderWidth,
			Height: float32(PauseTitleFontSize)}
		draw_rect(&TitleRect, rl.Black)
		rl.DrawText(PauseTitle, TitleRect.ToInt32().X+Diff, TitleRect.ToInt32().Y, PauseTitleFontSize, rl.White)
	}
}

func on_continue(GS *GameState) {
	change_to_playing(GS)
}

func on_return_to_main_menu(GS *GameState) {
	restart_game(GS)
	change_to_menu(GS)
}
