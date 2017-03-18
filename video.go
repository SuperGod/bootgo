package kernel

const (
	COLOR_BLACK = iota
	COLOR_BLUE
	COLOR_GREEN
	COLOR_CYAN
	COLOR_RED
	COLOR_MAGENTA
	COLOR_BROWN
	COLOR_LIGHT_GREY
	COLOR_DARK_GREY
	COLOR_LIGHT_BLUE
	COLOR_LIGHT_GREEN
	COLOR_LIGHT_CYAN
	COLOR_LIGHT_RED
	COLOR_LIGHT_MAGENTA
	COLOR_LIGHT_BROWN
	COLOR_WHITE
)

// GraphicsVideo graphics mode
type GraphicsVideo struct {
	Buffer uintptr
	Width  int
	Height int
	Depth  int
}

var (
	BackgroundColor uint8
	FrontColor      uint8
	Width           int
	Height          int
	Buffer          uintptr = 0xA0000
	Color           uint8
	graphicsVideo   GraphicsVideo
)

// Init set the graphics video params
func (v *GraphicsVideo) Init(width, height, depth int) {
	v.Buffer = Buffer
	v.Width = width
	v.Height = height
	v.Depth = depth
}

func terminalInit() {

	var buf []byte
	buf[1] = 1
	var str string
	str = "ooo"
	buf[2] = str[1]

	for y := 0; y < Height; y += 1 {
		for x := 0; x < Width; x += 1 {
			// drawAt(x, y, Color)
			// terminalPutEntryAt(0, Color, x, y)
		}
	}
}

func (v *GraphicsVideo) DrawBox(x1, y1, x2, y2 int, color int) {
	var index int
	var nColor uint8
	nColor = uint8(color)
	for y := y1; y != y2; y++ {
		for x := x1; x != x2; x++ {
			index = y*v.Width + x
			WriteMemory(v.Buffer, index, nColor)
		}
	}
}
