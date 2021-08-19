package main

import (
	"github.com/makotechg/standard-logger/stdlogger"
)

func main() {
	logger := stdlogger.NewStdLogger()

	logger.StdPrint("stdout")
	logger.ErrPrint("stderr")
}
