package kernel

func Itoa(n, base int64, buf [36]byte) int {
	var i int
	if base == 0 {
		i = len(buf) - 1
		buf[i] = '0'
		return i
	}
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
	return i
}
