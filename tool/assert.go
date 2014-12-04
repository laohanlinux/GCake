package tool

import (
	"os"
)

const (
	ExitCode = -127
)

func AssertEqual(a, b interface{}) {
	if a != b {
		os.Exit(ExitCode)
	}
}

func Assert(b bool) {
	if b {
		return
	} else {
		os.Exit(ExitCode)
	}
}
