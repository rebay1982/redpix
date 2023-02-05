package redpix

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"log"
	"runtime"
//	"image"
//	"image/color"
)

// WindowConfig: Structure describing starting parameters for the application's window.
type WindowConfig struct {
	
	// The window's title bar caption.
	Title string

	// The window's starting width, in pixels.
	Width int

	// The window's starting height, in pixels.
	Height int

	// Specifies if the window is resizable or not.
	Resizable bool

	// Enable or disable VSych
	VSync bool
}


// Init: Function to initialize GLFW, OpenGL and setup the window configuration.
func Init(config WindowConfig) {
	runtime.LockOSThread()

	//window := initGLFW(config)
	initGLFW(config)
	initOpenGL()
}

func initGLFW(config WindowConfig) *glfw.Window {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	// Process configurations
	if config.Resizable {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	} else {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}

	if (config.VSync) {
		glfw.SwapInterval(1)	// Enable vsync
	} else {
		glfw.SwapInterval(0)
	}

	window, err := glfw.CreateWindow(config.Width, config.Height, config.Title, nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func initOpenGL() {
	err := gl.Init()
	if err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version: ", version)
}
