package visualizer_test

import (
	"image/color"
	"testing"

	"github.com/kkevinchou/adventofcode/visualizer"
)

func Test(t *testing.T) {
	width := 600
	height := 300
	v := visualizer.New("out", width, height)

	red := color.RGBA{255, 0, 0, 255}
	green := color.RGBA{0, 255, 0, 255}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			v.Draw(x, y, green)
		}
	}

	offset := 0
	for i := 0; i < 5; i++ {
		for y := 0; y < height; y++ {
			for x := offset; x < offset+100; x++ {
				v.Draw(x, y, red)
			}
		}
		v.SaveToFile()
		offset += 100
	}

	v.CreateGIF(1)
}
