package caller

import "runtime"

func GetCaller(skip int) (runtime.Frame, bool) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	return frames.Next()
}
