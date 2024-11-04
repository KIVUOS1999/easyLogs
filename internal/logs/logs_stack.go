package logs

import (
	"bytes"
	"runtime"
)

func getStackTrace() string {
	buf := make([]byte, 1<<15)
	runtime.Stack(buf, true)

	str := bytes.TrimRightFunc(buf, func(r rune) bool {
		return r == '\x00'
	})

	return string(str)
}
