package goversion

import "runtime"

func Version() string {
	// Version returns the current Go version
	return runtime.Version()
}
