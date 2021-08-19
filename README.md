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