package luhn2

import (
    "fmt"
    "runtime"
)

func debugPrintf(fmt_ string, args ...interface{}) {
	programCounter, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(programCounter)
	prefix := fmt.Sprintf("[%s:%s %d] %s", file, fn.Name(), line, fmt_)
	fmt.Printf(prefix, args...)
	fmt.Println()
}

