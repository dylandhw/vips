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
	row     int16
	col     int16
	pattern cfaPattern
}

// remember to port this over for live video processing
var ImageFrames = []string{
	"cell_0.jpg",
	"cell_1.jpg",
	"cell_2.jpg",
	"cell_3.jpg",
	"cell_4.jpg",
	"cell_5.jpg",
	"cell_6.jpg",
	"cell_7.jpg",
	"cell_8.jpg",
}

func ProcessImageFrames() {
	if cellBuffer == nil {
		panic("buffer failed")
	}
	for _, frame := range ImageFrames {
		ExtractPixelStats(frame)
	}
}

func ExtractPixelStats(frame string) {}
