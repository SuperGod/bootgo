package kernel

import "unsafe"

var ()

type TextVideo struct {
	Buffer uintptr
	Width  int
	Height int
	Depth  int
	Color  uint8
	Column int
	Row    int
}

func (v *TextVideo) Init(width, height, depth int) {
	v.Buffer = 0xB8000
	v.Width = width
	v.Height = height
	v.Depth = depth
	v.Row = 0
	v.Column = 0
}

func (v *TextVideo) SetColor(background, front uint8) {
	v.Color = makeColor(front, background)
}

func (v *TextVideo) PrintCharAt(c byte, color uint8, x int, y int) {
	index := y*Width + x
	addr := (*uint16)(unsafe.Pointer(v.Buffer + 2*uintptr(index)))
	*addr = makeVGAEntry(c, color)
}

func (v *TextVideo) PrintChar(c byte) {
	v.PrintCharAt(c, v.Color, v.Column, v.Row)
	v.Column += 1
	if v.Column == v.Width {
		v.Column = 0
		v.Row += 1
		if v.Row == v.Height {
			v.Row = 0
		}
	}
}

func (v *TextVideo) Println(data string) {
	for i := 0; i < len(data); i += 1 {
		v.PrintChar(data[i])
	}
}

func (v *TextVideo) PrintInt64(n, base int64) {
	var buf [36]byte
	var i int
	if base == 0 {
		i = len(buf) - 1
		buf[i] = '0'
	} else {
		var cData byte
		var bMinus bool
		i = len(buf)
		if n == 0 {
			i = len(buf) - 1
			buf[i] = '0'
		} else if n < 0 {
			bMinus = true
			n = -n
		}
		for n > 0 {
			i--
			cData = "0123456789abcdef"[byte(n%base)]
			buf[i] = cData
			n = n / base
		}
		if bMinus {
			i--
			buf[i] = '-'
		}
	}
	for j := i; j != len(buf); j++ {
		v.PrintChar(buf[j])
	}

}

func (v *TextVideo) PrintBytes(data []byte) {
	for i := 0; i < len(data); i += 1 {
		v.PrintChar(data[i])
	}
}

func makeColor(fg uint8, bg uint8) uint8 {
	return fg | bg<<4
}

func makeVGAEntry(c byte, color uint8) uint16 {
	return uint16(c) | uint16(color)<<8
}
