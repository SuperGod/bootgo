package kernel

import "unsafe"

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

var (
	BackgroundColor uint8
	FrontColor      uint8
	Width           int
	Height          int
	Buffer          uintptr = 0xA0000
	Color           uint8
)

func makeColor(fg uint8, bg uint8) uint8 {
	return fg | bg<<4
}

func makeVGAEntry(c byte, color uint8) uint16 {
	return uint16(c) | uint16(color)<<8
}

func Init(width, height, fontColor, background int) {
	BackgroundColor = uint8(background)
	FrontColor = uint8(fontColor)
	Width = width
	Height = height
	Color = makeColor(FrontColor, BackgroundColor)
	Color = BackgroundColor
	// var i int
	// for i = 0; i != 0x1ffff; i++ {
	// 	addr := (*uint8)(unsafe.Pointer(uintptr(0xA0000 + i)))
	// 	*addr = 15
	// }
	terminalInit()
}

func terminalInit() {

	var buf []byte
	buf[1] = 1
	var str string
	str = "ooo"
	buf[2] = str[1]

	for y := 0; y < Height; y += 1 {
		for x := 0; x < Width; x += 1 {
			drawAt(x, y, Color)
			// terminalPutEntryAt(0, Color, x, y)
		}
	}
}

func drawAt(x, y int, color uint8) {
	index := y*Width + x
	WriteMemory(Buffer, index, color)
	// index := y*Width*3 + x*3
	// ptr := Buffer + uintptr(index)
	// addr := (*uint8)(unsafe.Pointer(ptr))
	// *addr = 0xff
	// addr = (*uint8)(unsafe.Pointer(ptr + 1))
	// *addr = 0xff
	// addr = (*uint8)(unsafe.Pointer(ptr + 2))
	// *addr = 0xff
}

func terminalPutEntryAt(c byte, color uint8, x int, y int) {
	index := y*Width + x
	addr := (*uint16)(unsafe.Pointer(Buffer + 2*uintptr(index)))
	*addr = makeVGAEntry(c, color)
}

// func terminalPutChar(c byte) {
// 	terminalPutEntryAt(c, color, column, row)
// 	column += 1
// 	if column == VGA_WIDTH {
// 		column = 0
// 		row += 1
// 		if row == VGA_HEIGHT {
// 			row = 0
// 		}
// 	}
// }
// func writeString(data string) {
// 	for i := 0; i < len(data); i += 1 {
// 		terminalPutChar(data[i])
// 	}
// }
