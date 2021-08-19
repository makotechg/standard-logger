package main

import (
	"github.com/makotechg/standard-logger/stdlogger"
)

func main() {
	//logger := stdlogger.NewStdLogger()

	stdlogger.SL.StdPrint("stdout")
	stdlogger.SL.ErrPrint("stderr")

	stdlogger.SL.StdPrintf("stdprintf:%s", "std output")
	stdlogger.SL.ErrPrintf("errprintf:%s", "err output")
}
