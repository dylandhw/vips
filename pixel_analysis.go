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

func ExtractPixelStats(cellPixels []byte, cellWidth, cellHeight, row, col int, bayerPattern string) Grid {
	var rPixels, g1Pixels, g2Pixels, bPixels []byte

	for y := 0; y < cellHeight; y++ {
		for x := 0; x < cellWidth; x++ {
			pixel := cellPixels[y*cellWidth*x]
			even_row := y%2 == 0
			even_col := x%2 == 0

			switch bayerPattern {
			case "RGGB":
				if even_row && even_col {
					rPixels = append(rPixels, pixel)
				} else if even_row && !even_col {
					g1Pixels = append(g1Pixels, pixel)
				} else if !even_row && even_col {
					g2Pixels = append(g2Pixels, pixel)
				} else {
					bPixels = append(bPixels, pixel)
				}
			}
		}
	}

	return Grid{
		row: int16(row),
		col: int16(col),
		pattern: cfaPattern{
			r:  computePixelStats(rPixels),
			g1: computePixelStats(g1Pixels),
			g2: computePixelStats(g2Pixels),
			b:  computePixelStats(bPixels),
		},
	}
}

// here we will be computing mean, variance, gradient energy, and bit entropy
func computePixelStats(pixels []byte) Pixel {
	if len(pixels) == 0 {
		return Pixel{}
	}

	sum := 0
	for _, p := range pixels {
		sum += int(p)
	}

	return Pixel{}
}
