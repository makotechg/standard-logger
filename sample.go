package sample

func main() {
	logger := stdlogger.NewStdLogger()

	logger.StdPrint("stdout")
	logger.ErrPrint("stderr")
}
