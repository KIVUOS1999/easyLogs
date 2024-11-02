package logs

import (
	"runtime"
)

func printStackTrace() string {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	return string(buf)
}
