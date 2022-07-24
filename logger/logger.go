package logger

import (
	"fmt"
	"github.com/fajarardiyanto/flt-go-utils/caller"
	"github.com/jwalton/gchalk"
	"strconv"
)

func Logger(statusCode int, msg string) {
	call, _ := caller.GetCaller(2)
	fmt.Printf("[%s][%s][%s] %s",
		gchalk.BrightYellow(strconv.Itoa(statusCode)), gchalk.BrightWhite(call.Function),
		gchalk.BrightCyan(strconv.Itoa(call.Line)),
		msg)
}
