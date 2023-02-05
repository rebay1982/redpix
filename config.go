package redpix

const (
	MAX_WINDOW_WIDTH = 1980 
	MAX_WINDOW_HEIGHT = 1080
)

// WindowConfig: Structure describing starting parameters for the application's window.
type WindowConfig struct {
	
	// The window's title bar caption.
	Title string

	// The window's starting width, in pixels. Min: 1, Max: 1980
	Width int

	// The window's starting height, in pixels. Min: 1, Max: 1080
	Height int

	// Specifies if the window is resizable or not.
	Resizable bool

	// Enable or disable VSych
	VSync bool
}

func (c *WindowConfig) validate() bool {

	if !(c.Width > 0 && c.Width <= MAX_WINDOW_WIDTH) {
		return false
	}

	if !(c.Height > 0 && c.Height <= MAX_WINDOW_HEIGHT) {
		return false
	}

	return true
}


