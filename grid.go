package main

import (
	"bytes"
	"fmt"
	"image"
	"sync"

	"gocv.io/x/gocv"
)

var rows = 3
var columns = 3

type CellBuffer struct {
	index  int
	buffer *bytes.Buffer
}

func PartitionImage() {

	frame := gocv.IMRead("./bayer-simulation2.png", gocv.IMReadAnyDepth|gocv.IMReadGrayScale)

	if frame.Empty() {
		fmt.Println("error reading image")
		return
	}
	defer frame.Close()

	frameRows := frame.Rows()
	frameColumns := frame.Cols()
	cellHeight := frameRows / rows
	cellWidth := frameColumns / columns

	totalCells := rows * columns
	cellBuffers := make([]*bytes.Buffer, totalCells)

	type cellData struct {
		index int
		mat   gocv.Mat
	}
	cells := make([]cellData, 0, totalCells)

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
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

			index := r*columns + c
			cell := frame.Region(image.Rect(x1, y1, x2, y2))
			cells = append(cells, cellData{index, cell})
		}
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, cd := range cells {
		wg.Add(1)
		go func(cd cellData) {
			defer wg.Done()
			defer cd.mat.Close()

			buffer, err := gocv.IMEncode(".jpg", cd.mat)
			if err != nil {
				fmt.Printf("trouble encoding cell %d\n", cd.index)
				return
			}
			buf := bytes.NewBuffer(buffer.GetBytes())
			buffer.Close()

			_ = mu
			cellBuffers[cd.index] = buf
			fmt.Printf("cell %d buffer size: %d bytes\n", cd.index, buf.Len())
		}(cd)
	}

	wg.Wait()

	fmt.Printf("cellBuffers %v", cellBuffers)
}
