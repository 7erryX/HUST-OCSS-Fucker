package utils

import (
	"fmt"

	"github.com/bytedance/gopkg/util/gopool"
)

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}
	// logrus.Warnf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	//os.Exit(1)
	panic(err)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	// logrus.Infof("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	// logrus.Warnf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func Go(f func()) {
	gopool.Go(f)
}
