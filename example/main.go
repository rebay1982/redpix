package main 

import (
	"fmt"
	rp "github.com/rebay1982/redpix"
)

func update() {
	fmt.Println("I am update")
}

func draw() {
	fmt.Println("I am draw")
}

func main() {
	config := rp.WindowConfig {
		Width: 640,
		Height: 480,
	}

	rp.Init(config)
	rp.Run(update, draw)

}
