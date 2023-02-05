package redpix

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"log"
	"runtime"
//	"image"
//	"image/color"
)

var (
	window *glfw.Window
)

func init() {
	runtime.LockOSThread()
}

func initGLFW(config WindowConfig) *glfw.Window {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	// Process configurations
	if (config.Resizable) {
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

// Init: Function to initialize GLFW, OpenGL and setup the window configuration.
func Init(config WindowConfig) {
	if !config.validate() {
		log.Panic("redpix: Failed to init, incomplete configuration.")
	}

	window = initGLFW(config)
	initOpenGL()
}

func Run(update func()) {
	if (window == nil) {
		log.Panic("redpix: Cannot run, engine not initalized")
	}
}
