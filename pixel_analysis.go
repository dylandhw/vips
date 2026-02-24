package main

type Pixel struct {
	mean            float32
	variance        float32
	gradient_energy uint16
	bit_entropy     int16
}

type cfaPattern struct {
	r  Pixel
	g1 Pixel
	g2 Pixel
	b  Pixel
}

type Grid struct {
	row            int16
	col            int16
	pattern        cfaPattern
	x1, y1, x2, y2 int
}

func ExtractPixelStats(cellPixels []byte, cellWidth, cellHeight, rol, col int, bayerPattern string) Grid {
	var rPixels, g1Pixels, g2Pixels, bPixels []byte

	for y := 0; y < cellHeight; y++ {
		for x := 0; x < cellWidth; x++ {

		}
	}

	return Grid{}
}
