package main

import (
	"fmt"
	"image"
	"image/color"
	"sync"
)

type world struct {
	w, h    int
	r       float64
	cells   [][]bool
	cellsMu sync.RWMutex
}

func newWorld(w, h int) *world {
	// [ [..w..] ..h.. [..w..]]
	row := make([][]bool, h)
	for i := 0; i < h; i++ {
		row[i] = make([]bool, w)
	}

	return &world{
		w:     w,
		h:     h,
		cells: row,
	}
}

// for color interface
func (w *world) ColorModel() color.Model {
	return color.GrayModel
}

// for color interface
func (w *world) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{},
		Max: image.Point{X: w.w, Y: w.h},
	}
}

func (w *world) outside(x, y, b int) bool {
	if b > w.w/2 || b > w.h/2 {
		return true
	}
	return x < b || x >= w.w-b || y < b || y >= w.h-b
}

// for color interface
func (w *world) At(x, y int) color.Color {
	if w.outside(x, y, 0) || !w.cells[y][x] {
		return color.Gray{255}
	}
	return color.Gray{0}
}

func (w *world) String() string {
	if w.w*w.h > 200 {
		return fmt.Sprintf("image %d x %d. too big into string", w.w, w.h)
	}

	var m = make([]byte, 0)

	for _, row := range w.cells {
		for _, cell := range row {
			if cell {
				m = append(m, '*')
			} else {
				m = append(m, ' ')
			}
		}
		m = append(m, '|', '\n')
	}
	return string(m)
}

func (w *world) set(x, y int, b bool) error {
	w.cellsMu.Lock()
	defer w.cellsMu.Unlock()
	if w.outside(x, y, 0) {
		return fmt.Errorf("out of bound: %d %d", x, y)
	}

	w.cells[y][x] = b
	return nil
}

func (w *world) get(x, y int) (bool, error) {
	if w.outside(x, y, 0) {
		return false, fmt.Errorf("out of bound: %d %d", x, y)
	}

	return w.cells[y][x], nil
}

func (w *world) isTouch(x, y int) (bool, error) {
	if w.outside(x, y, 1) {
		return false, fmt.Errorf("out of bound %d %d", x, y)
	}
	return w.cells[y-1][x] || w.cells[y+1][x] || w.cells[y][x-1] || w.cells[y][x+1], nil
}
