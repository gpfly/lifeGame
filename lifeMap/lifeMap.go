package lifeMap

import (
	"image"
	"image/color"
	"log"
	"math/rand"
)

type Map struct {
	Img     *image.RGBA
	dx      int
	dy      int
	dw      int
	curMap  [][]uint8
	nextMap [][]uint8
}

func NewMap(x, y, w int) Map {
	m := Map{
		dx: x,
		dy: y,
		dw: w,
	}
	m.createMap(x, y, w)
	m.reset()
	return m
}

func (m *Map) createMap(x, y, w int) {
	m.Img = image.NewRGBA(image.Rect(0, 0, x*w, y*w))
	m.dx = x
	m.dy = y

	for i := 0; i < x; i++ {
		ar1 := make([]uint8, y)
		ar2 := make([]uint8, y)
		m.curMap = append(m.curMap, ar1)
		m.nextMap = append(m.nextMap, ar2)
	}
}

func (m *Map) reset() {
	for i := 0; i < m.dx; i++ {
		for j := 0; j < m.dy; j++ {
			if uint8(rand.Intn(2)) == 0 {
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
	if row < 0 || col < 0 || row >= m.dx || col >= m.dy {
		return
	}
	m.nextMap[row][col] = val
}

func (m *Map) GetCur(row, col int) uint8 {
	if row < 0 || col < 0 || row >= m.dx || col >= m.dy {
		return 0
	}
	return m.curMap[row][col]
}

func (m *Map) GameCycle() {
	for x := 0; x < m.dx; x++ {
		for y := 0; y < m.dy; y++ {
			count := m.GetNeighborCount(x, y)
			switch {
			case count >= 4:
				m.SetNext(x, y, 0)
			case count >= 3:
				m.SetNext(x, y, 1)
			case count >= 2:
				m.SetNext(x, y, m.GetCur(x, y))
			case count <= 1:
				m.SetNext(x, y, 0)
			default:
				log.Println("error count", count)
			}
		}
	}
	for x := 0; x < m.dx; x++ {
		for y := 0; y < m.dy; y++ {
			m.curMap[x][y] = m.nextMap[x][y]
		}
	}
}

func (m *Map) PrintMap() {
	for x := 0; x < m.dx; x++ {
		for y := 0; y < m.dy; y++ {
			if m.GetCur(x, y) == 0 {
				m.setimg(x, y, color.White)
			} else {
				m.setimg(x, y, color.Black)
			}
		}
	}
}

func (m *Map) setimg(x, y int, c color.Color) {
	for i := 1; i < m.dw-1; i++ {
		for j := 1; j < m.dw-1; j++ {
			m.Img.Set(x*m.dw+i, y*m.dw+j, c)
		}
	}
}
