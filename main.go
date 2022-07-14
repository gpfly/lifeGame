package main

import (
	"lifeGame/lifeMap"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

const (
	dx = 80
	dy = 80
)

func main() {
	m := lifeMap.NewMap(dx, dy, 5)
	m.PrintMap()

	lifeApp := app.New()
	lifeWindows := lifeApp.NewWindow("lifeGame")

	lifeImg := canvas.NewImageFromImage(m.Img)
	lifeImg.FillMode = canvas.ImageFillOriginal

	content := container.New(layout.NewCenterLayout(), lifeImg)
	lifeWindows.SetContent(content)

	go func() {
		for {
			m.GameCycle()
			m.PrintMap()
			content.Refresh()
			time.Sleep(100 * time.Millisecond)
		}
	}()
	lifeWindows.CenterOnScreen()
	lifeWindows.ShowAndRun()
}
