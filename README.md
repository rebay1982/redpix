# redpix
Super simple OpenGL/GLFW wrapper to blit images to a framebuffer.


## Installation and Usage

Installation is as with any go module.
```
go get github.com/rebay1982/redpix
```

### General
redpix requires you to initialize the microlibrary by passing it a configuration, and providing it a draw function.
The draw function must return a `[]int8` slice that contains the pixels to be drawn on screen. The format of the data
is RGBA and is assumed to be the same size as the window size provided with the configuration during initialization.

### Confguration
The configuration requires that you set at minimum the title for the screen, the width, and the height.

Optional paramters include the Resizable window hint that prevents the window from being resized if set to `false`,
and VSync which enables vertical sync when set to `true`.

This configuration must be passed to the Init function to correctly initialize redpix.

### Running
After initlization, to get the microlibrary running, you must provide it at least one draw function. This draw function
must return a `[]uint8` slice of RGBA pixel data with a size maching that of the requested window size during the
initialization phase. This is the data the library will use to display on screen.

You can optionally pass it an update function, this can be used to update an internal state used to render things in
the draw method.

### Example
A very basic and minimal example can be studied [here](https://github.com/rebay1982/redpix/blob/main/example/main.go).

## Wayland
To build against Wayland, you must specify the `-tags=wayland` tag. This is required for GLFW.
