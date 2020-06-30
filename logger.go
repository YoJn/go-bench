package go_bench

import (
	"io"
	"log"
)

var logger Logger

func DefaultLogger() BenchHandle {
	return LoggerWithWriter(nil)
}

func LoggerWithWriter(writer *io.Writer) BenchHandle {
	if writer ==nil{
		writer = &defaultWriter
	}
	logger = Logger(log.New(*writer,"bench",log.LstdFlags))
	return func(engine *BenchEngine) {

	}
}

// Logger is used for logging formatted messages.
type Logger interface {
	// Printf must have the same semantics as log.Printf.
	Printf(format string, args ...interface{})
}