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
    
## Update package
- Check if the latest version available exists
    - in PACKAGE : `go list -u -m ${PACKAGE}`
    - in all : `go list -u -m all`
- Update to the latest version
    - in PACKAGE : `go get PACKAGE@latest`
    - in all : `go get -u`