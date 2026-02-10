package main

import (
	"bytes"
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

var rows = 3
var columns = 3

var cellBuffer *bytes.Buffer

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

	count := 0
	// let's see how long this operation takes
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

			cell := frame.Region(image.Rect(x1, y1, x2, y2))
			defer cell.Close()

			// with image write = ~1.11ms, without = ~3.8µs
			//filename := fmt.Sprintf("cell_%d.jpg", count)
			//gocv.IMWrite(filename, cell)

			// what we need to do here is load the specific grid into a buffer
			// then hand it over to pixel analysis
			//
			// let's eventually parallelize our buffers... we will work on that tomorrow. should speed things up significantly!
			buffer, err := gocv.IMEncode(".jpg", cell)

			if err != nil {
				fmt.Printf("trouble loading cell into buffer")
			}

			defer buffer.Close()

			cellBuffer = bytes.NewBuffer(buffer.GetBytes())
			fmt.Printf(">> Buffer size: %d bytes\n", cellBuffer.Len())

			count++
		}
	}
}
