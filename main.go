package main

import (
	"path/filepath"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	var window_width int = 800
	var window_height int = 450

	rl.InitWindow(int32(window_width), int32(window_height), "MuPl")
	defer rl.CloseWindow()
	rl.InitAudioDevice()

	rl.SetTargetFPS(60)

	var time_button_x int = 8

	var music_count int = 0
	var musics [3]string = [3]string{"Music\\Chopin Torrent.mp3", "Music\\Rondo Alla Turca.mp3", "Music\\Hungarian Rhapsody 2.mp3"}

	var run_button_x int = 8
	var volume float32 = 0

	var low_border_run_button int = 7
	var low_border_music_button int = 7
	var top_border_run_button int = 745
	var top_border_music_button int = 790

	music_path := musics[music_count]
	music := rl.LoadMusicStream(music_path)
	music_name := strings.ReplaceAll(filepath.Base(music_path), ".mp3", "")
	music_length := rl.GetMusicTimeLength(music)

	music_button := rl.Rectangle{X: 50, Y: 50, Width: 100, Height: 100}
	pause_button := rl.Rectangle{X: 175, Y: 50, Width: 100, Height: 100}
	next_music_button := rl.Rectangle{X: 700, Y: 200, Width: 50, Height: 50}
	back_music_button := rl.Rectangle{X: 640, Y: 200, Width: 50, Height: 50}

	run_button := rl.Rectangle{X: float32(run_button_x), Y: 378, Width: 50, Height: 50}
	line_for_run_button := rl.Rectangle{X: 6, Y: 400, Width: 790, Height: 5}

	exit_button := rl.Rectangle{X: 720, Y: 20, Width: 50, Height: 30}

	time_music_button := rl.Rectangle{X: float32(time_button_x), Y: 302, Width: 50, Height: 50}
	line_for_music_button := rl.Rectangle{X: 6, Y: 325, Width: 790, Height: 5}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Pink)

		rl.DrawRectangle(int32(music_button.X), int32(music_button.Y), int32(music_button.Width), int32(music_button.Height), rl.Red)
		rl.DrawText("Play", int32(music_button.X)+8, int32(music_button.Y)+30, 40, rl.Black)

		rl.DrawRectangle(int32(pause_button.X), int32(pause_button.Y), int32(pause_button.Width), int32(pause_button.Height), rl.Red)
		rl.DrawText("Pause", int32(pause_button.X)+5, int32(pause_button.Y)+37, 30, rl.Black)

		rl.DrawRectangle(int32(line_for_music_button.X), int32(line_for_music_button.Y), int32(line_for_music_button.Width), int32(line_for_music_button.Height), rl.White)
		rl.DrawRectangle(int32(time_button_x), int32(time_music_button.Y), int32(time_music_button.Width), int32(time_music_button.Height), rl.Red)

		rl.DrawRectangle(6, 400, 790, 5, rl.White)

		rl.DrawRectangle(int32(run_button_x), int32(run_button.Y), int32(run_button.Width), int32(run_button.Height), rl.Red)

		rl.DrawRectangle(720, 20, 50, 30, rl.Red)
		rl.DrawText("Exit", 727, 26, 20, rl.Black)

		rl.DrawText(music_name, 10, 200, 50, rl.Black)

		rl.DrawRectangle(int32(next_music_button.X), int32(next_music_button.Y), int32(next_music_button.Width), int32(next_music_button.Height), rl.Red)
		rl.DrawText("Next", int32(next_music_button.X)+2, int32(next_music_button.Y)+15, 20, rl.Black)

		rl.DrawRectangle(int32(back_music_button.X), int32(back_music_button.Y), int32(back_music_button.Width), int32(back_music_button.Height), rl.Red)
		rl.DrawText("Back", int32(back_music_button.X)+1, int32(back_music_button.Y)+15, 20, rl.Black)

		music_time_played := rl.GetMusicTimePlayed(music)

		//Кнопка выхода
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), exit_button) {
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				rl.StopMusicStream(music)
				rl.UnloadMusicStream(music)
				rl.EndDrawing()
				rl.CloseWindow()
			}
		}

		//Остановка музыки
		if time_button_x >= 743 {
			rl.StopMusicStream(music)
			time_button_x = 8
		} else {
			rl.UpdateMusicStream(music)
		}

		//Включение следующей музыки
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), next_music_button) {
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				rl.StopMusicStream(music)

				if music_count+1 > len(musics)-1 {
					music_count = 0
				} else {
					music_count = music_count + 1
				}

				music_path = musics[music_count]
				music_name = strings.ReplaceAll(filepath.Base(music_path), ".mp3", "")
				music = rl.LoadMusicStream(music_path)
				music_length = rl.GetMusicTimeLength(music)
				time_button_x = 8
			}
		}

		//Включение предыдущей музыки
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), back_music_button) {
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				rl.StopMusicStream(music)

				if music_count-1 < 0 {
					music_count = len(musics) - 1
				} else {
					music_count = music_count - 1
				}

				music_path = musics[music_count]
				music_name = strings.ReplaceAll(filepath.Base(music_path), ".mp3", "")
				music = rl.LoadMusicStream(music_path)
				music_length = rl.GetMusicTimeLength(music)
				time_button_x = 8
			}
		}

		//Изменение координат ползунка времени
		if rl.IsMusicStreamPlaying(music) {
			time_button_x = int(735*music_time_played)/int(music_length) + 8
		}

		//Изменение координат полузунка времени один нажатием
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), line_for_music_button) {
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				if rl.GetMousePosition().X > float32(low_border_music_button) && rl.GetMousePosition().X < float32(top_border_music_button) {
					time_button_x = int(rl.GetMousePosition().X) - int(music_button.Width)/2
					rl.SeekMusicStream(music, float32(time_button_x)*music_length/735)
				}
			}
		}

		//Изменение координат ползунка громкости
		if run_button_x > low_border_run_button && run_button_x < top_border_run_button {
			if rl.IsKeyDown(rl.KeyRight) {
				run_button_x = run_button_x + 2
				if run_button_x >= top_border_run_button {
					for run_button_x >= top_border_run_button {
						run_button_x = run_button_x - 1
					}
				}
			} else if rl.IsKeyDown(rl.KeyLeft) {
				run_button_x = run_button_x - 2
				if run_button_x <= low_border_run_button {
					for run_button_x <= low_border_run_button {
						run_button_x = run_button_x + 1
					}
				}
			}
		}

		//Перемещение ползунка громкости одним кликом
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), line_for_run_button) {
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				if rl.GetMousePosition().X > float32(low_border_run_button) && rl.GetMousePosition().X < float32(top_border_run_button) {
					run_button_x = int(rl.GetMousePosition().X)
				}
			}
		}

		//Включение музыки
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), music_button) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				if !rl.IsMusicStreamPlaying(music) {
					rl.PlayMusicStream(music)
				}
			}
		}

		//Постановка музыки на паузу
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), pause_button) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				if !rl.IsMusicStreamPlaying(music) {
					rl.ResumeAudioStream(music.Stream)
				} else {
					rl.PauseAudioStream(music.Stream)
				}
			}
		}

		//Махинации со звуком
		volume = float32(run_button_x) / 740

		if run_button_x == low_border_run_button+1 {
			volume = 0
		}

		rl.SetMusicVolume(music, volume)

		rl.EndDrawing()
	}

	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()

}
