package main

import (
	"image"
	"sync"

	"gocv.io/x/gocv"
)

var rows = 3
var columns = 3

func PartitionImage() []Grid {
	frame := gocv.IMRead("./bayer-simulation2.png", gocv.IMReadAnyDepth|gocv.IMReadGrayScale)

	defer frame.Close()

	frameRows := frame.Rows()
	frameColumns := frame.Cols()
	cellHeight := frameRows / rows
	cellWidth := frameColumns / columns

	grids := make([]Grid, rows*columns)
	var wg sync.WaitGroup

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			wg.Add(1)
			go func(r, c int) {
				defer wg.Done()
				y1 := r * cellHeight
				y2 := (r + 1) * cellHeight
				x1 := c * cellWidth
				x2 := (c + 1) * cellWidth
				if r == rows-1 {
					y2 = frameRows
				}
				if c == columns-1 {
					x2 = frameColumns
				}

				cell := frame.Region(image.Rect(x1, y1, x2, y2))
				pixels := cell.ToBytes()
				cell.Close()
				grids[r*columns+c] = ExtractPixelStats(pixels, cellWidth, cellHeight, r, c, "RGGB")

			}(r, c)
		}
	}

	wg.Wait()
	return grids
}
