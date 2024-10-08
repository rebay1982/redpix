package main 

import (
	"fmt"
	"image"
	"image/color"

	rp "github.com/rebay1982/redpix"
)

const (
	WINDOW_TITLE = "RedPix Example"
	WINDOW_WIDTH = 640 
	WINDOW_HEIGHT = 480
)

func draw() []uint8 {
	w := WINDOW_WIDTH
	h := WINDOW_HEIGHT

	var img = image.NewRGBA(image.Rect(0, 0, w, h))
	var red = color.RGBA{255, 0, 0, 0}
	var blue = color.RGBA{0, 0, 255, 0}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {

			if ((x+y) % 2 > 0) {
				img.Set(x, y, red)

			} else {
				img.Set(x, y, blue)

			}
		}
	}

	return img.Pix
}

func input(event rp.InputEvent) {

	fmt.Print("Key ")
	switch event.Key {
	case rp.IN_PLAYER_FORWARD:
		fmt.Print("forward ")
	case rp.IN_PLAYER_BACKWARD:
		fmt.Print("backward ")
	case rp.IN_PLAYER_LEFT:
		fmt.Print("left ")
	case rp.IN_PLAYER_RIGHT:
		fmt.Print("right ")
	}

	fmt.Print("was ")
	switch event.Action {
	case rp.IN_ACT_PRESSED:
		fmt.Println("pressed.")
	case rp.IN_ACT_RELEASED:
		fmt.Println("released.")
	case rp.IN_ACT_REPEATED:
		fmt.Println("repeated.")
	}
}

func main() {
	config := rp.WindowConfig {
		Title: WINDOW_TITLE,
		Width: WINDOW_WIDTH,
		Height: WINDOW_HEIGHT,
		Resizable: false,
		VSync: true,
	}

	rp.Init(config, draw, input)
	rp.Run()
}
