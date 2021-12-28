package main

import (
	"image"
	"image/color"
	"log"
	"math/rand"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type Map struct {
	img     *image.RGBA
	x       int
	y       int
	curMap  [][]uint8
	nextMap [][]uint8
}

func NewMap(x, y int) Map {
	m := Map{}
	m.img = image.NewRGBA(image.Rect(0, 0, x, y))
	m.x = x
	m.y = y

	for i := 0; i < x; i++ {
		ar1 := make([]uint8, y)
		ar2 := make([]uint8, y)
		m.curMap = append(m.curMap, ar1)
		m.nextMap = append(m.nextMap, ar2)
	}

	return m
}

func (m *Map) Reset() {
	for i := 0; i < m.x; i++ {
		for j := 0; j < m.y; j++ {
			if uint8(rand.Intn(5)) == 0 {
				m.curMap[i][j] = 1
			} else {
				m.curMap[i][j] = 0
			}
		}
	}
}

func (m *Map) GetNeighborCount(row, col int) uint8 {
	var count uint8 = 0
	if m.GetCur(row-1, col-1) != 0 {
		count++
	}
	if m.GetCur(row-1, col) != 0 {
		count++
	}
	if m.GetCur(row-1, col+1) != 0 {
		count++
	}
	if m.GetCur(row, col-1) != 0 {
		count++
	}
	if m.GetCur(row, col+1) != 0 {
		count++
	}
	if m.GetCur(row+1, col-1) != 0 {
		count++
	}
	if m.GetCur(row+1, col) != 0 {
		count++
	}
	if m.GetCur(row+1, col+1) != 0 {
		count++
	}
	return count
}

func (m *Map) SetNext(row, col int, val uint8) {
	if row < 0 || col < 0 || row >= m.x || col >= m.y {
		return
	}
	m.nextMap[row][col] = val
}

func (m *Map) GetCur(row, col int) uint8 {
	if row < 0 || col < 0 || row >= m.x || col >= m.y {
		return 0
	}
	return m.curMap[row][col]
}

func (m *Map) GameCycle() {
	for x := 0; x < m.x; x++ {
		for y := 0; y < m.y; y++ {
			count := m.GetNeighborCount(x, y)
			switch {
			case count >= 4:
				m.SetNext(x, y, 0)
			case count == 3:
				m.SetNext(x, y, 1)
			case count == 2:
				m.SetNext(x, y, m.GetCur(x, y))
			case count <= 1:
				m.SetNext(x, y, 0)
			default:
				log.Println("error count", count)
			}
		}
	}
	for x := 0; x < m.x; x++ {
		for y := 0; y < m.y; y++ {
			m.curMap[x][y] = m.nextMap[x][y]
		}
	}
}

func (m *Map) PrintMap() {
	for x := 0; x < m.x; x++ {
		for y := 0; y < m.y; y++ {
			if m.GetCur(x, y) == 0 {
				m.img.Set(x, y, color.White)
			} else {
				m.img.Set(x, y, color.Black)
			}
		}
	}
}

const (
	dx = 400
	dy = 400
)

func main() {
	m := NewMap(dx, dy)
	m.Reset()
	m.PrintMap()

	lifeApp := app.New()
	lifeWindows := lifeApp.NewWindow("lifeGame")

	lifeImg := canvas.NewImageFromImage(m.img)
	lifeImg.FillMode = canvas.ImageFillOriginal

	content := container.New(layout.NewCenterLayout(), lifeImg)
	lifeWindows.SetContent(content)

	go func() {
		for {
			m.GameCycle()
			m.PrintMap()
			content.Refresh()
			// time.Sleep(100 * time.Millisecond)
		}
	}()
	lifeWindows.CenterOnScreen()
	lifeWindows.ShowAndRun()
}
