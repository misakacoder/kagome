package errs

import (
	"fmt"
	"runtime"
	"strings"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func GetStackTrace(err any) string {
	var pc [64]uintptr
	n := runtime.Callers(2, pc[:])
	frames := runtime.CallersFrames(pc[:n])
	stackTrace := strings.Builder{}
	stackTrace.WriteString(fmt.Sprintf("%v", err))
	for {
		frame, more := frames.Next()
		if !strings.HasPrefix(frame.Function, "runtime.") {
			stackTrace.WriteString(fmt.Sprintf("\n - %s:%d (0x%x)", frame.File, frame.Line, frame.PC))
		}
		if !more {
			break
		}
	}
	return stackTrace.String()
}
