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
	texture uint32
)

func init() {
	runtime.LockOSThread()
}

func initFramebuffer(texture uint32) {
	var framebuffer uint32

	gl.GenFramebuffers(1, &framebuffer)
	gl.BindFramebuffer(gl.FRAMEBUFFER, framebuffer)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, texture, 0)

	gl.BindFramebuffer(gl.READ_FRAMEBUFFER, framebuffer)
	gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, 0)
}

func initTexture() uint32 {
	var texture uint32

	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

	gl.BindImageTexture(0, texture, 0, false, 0, gl.WRITE_ONLY, gl.RGBA8)

	return texture
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

	window, err := glfw.CreateWindow(config.Width, config.Height, config.Title, nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// Process vsync config later. It requires a OpenGL context.
	if (config.VSync) {
		glfw.SwapInterval(1)	// Enable vsync
	} else {
		glfw.SwapInterval(0)
	}

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
	texture = initTexture()
	initFramebuffer(texture)
}

func Run(update, draw func()) {
	if (window == nil) {
		log.Panic("redpix: Cannot run, engine not initalized")
	}

	if update != nil {
		update()
	}

	if draw != nil {
		draw()
	}
}
