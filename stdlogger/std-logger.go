package stdlogger

import (
	"log"
	"os"
)

type stdLogger struct {
	stdout *log.Logger
	stderr *log.Logger
}

func NewStdLogger() *stdLogger {
	l := new(stdLogger)
	l.stdout = log.New(os.Stdout, "[INFO]", log.Ldate|log.Lshortfile)
	l.stderr = log.New(os.Stderr, "[ERROR]", log.Ldate|log.Lshortfile)
	return l
}

func (sl *stdLogger) StdPrint(v_array ...interface{}) {
	for i := 0; i < len(v_array); i++ {
		sl.stdout.Printf("%+v", v_array[i])
	}
}

func (sl *stdLogger) StdPrintf(format string, v_array ...interface{}) {
	sl.stdout.Printf(format, v_array...)
}

func (sl *stdLogger) ErrPrint(v_array ...interface{}) {
	for i := 0; i < len(v_array); i++ {
		sl.stderr.Printf("%+v", v_array[i])
	}
}
func (sl *stdLogger) ErrPrintf(format string, v_array ...interface{}) {
	sl.stderr.Printf(format, v_array...)
}

func (sl *stdLogger) FatalPrint(v_array ...interface{}) {
	for i := 0; i < len(v_array); i++ {
		sl.stderr.Print(v_array[i])
	}
	log.Fatal("Fatal Error is occured")
}

func (sl *stdLogger) FatalPrintf(format string, v_array ...interface{}) {
	log.Fatalf(format, v_array...)
}
