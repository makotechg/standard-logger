# standard-logger
This is the minimum configuration logger package for standard output.

## How to use
- import "github.com/makotechg/standard-logger/stdlogger"

- Initialize
    - `logger := stdlogger.NewStdLogger()`

- standard output
    - `logger.StdPrint("your std output string here")`

- standard error output
    - `logger.ErrPrint("your err output string here")`

- fatal error output
    - If you use this, program will exit usind os.Exit(1).
    - `logger.FatalPrint("your err output string here")`

## Check output destination
- command `go run . 1> stdout.log 2> errout.log`
    - stdout.log
        - ```
            [INFO]2021/08/20 std-logger.go:22: stdout
            [INFO]2021/08/20 std-logger.go:27: stdprintf:std output
            ```
    - errout.log
        - ```
            [ERROR]2021/08/20 std-logger.go:32: stderr
            [ERROR]2021/08/20 std-logger.go:36: errprintf:err output
            ```
    