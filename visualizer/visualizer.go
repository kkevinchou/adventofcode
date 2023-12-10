package visualizer

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"os/exec"
)

type Pixel struct {
	color color.RGBA
}

type Visualizer struct {
	width, height int
	pixels        [][]Pixel
	prefix        string
	count         int
}

func New(prefix string, width, height int) *Visualizer {
	v := &Visualizer{
		width:  width,
		height: height,
		prefix: prefix,
	}

	for i := 0; i < height; i++ {
		pixelRow := make([]Pixel, width)
		v.pixels = append(v.pixels, pixelRow)
	}

	return v
}

func (v *Visualizer) Draw(x, y int, color color.RGBA) {
	v.pixels[y][x].color = color
}

func (v *Visualizer) SaveToFile() {
	fmt.Println(len(v.pixels))
	fmt.Println(len(v.pixels[0]))

	img := image.NewRGBA(image.Rect(0, 0, v.width, v.height))

	for y, pixelRow := range v.pixels {
		for x, pixel := range pixelRow {
			img.Set(x, y, pixel.color)
		}
	}

	// Create a new file to save the image
	err := os.MkdirAll(v.prefix, os.ModeDir)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(fmt.Sprintf("%s/%s%d.png", v.prefix, v.prefix, v.count))
	v.count += 1
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Encode the image and write it to the file
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

func (v *Visualizer) CreateGIF() {
	videoFile := fmt.Sprintf("%s.avi", v.prefix)
	gifFile := fmt.Sprintf("%s.gif", v.prefix)

	cmd := exec.Command("./ffmpeg.exe", "-framerate", "1", "-i", fmt.Sprintf("%s/%s%%d.png", v.prefix, v.prefix), videoFile, "-y")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))

	cmd = exec.Command("./ffmpeg.exe", "-i", videoFile, "-pix_fmt", "rgb32", "-loop", "0", gifFile, "-y")
	stdout, err = cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}
