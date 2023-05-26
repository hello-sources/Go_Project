package logs

import "io"

var writeAdapter = make(map[string]InitLogWriterFunc, 0)

type InitLogWriterFunc func() LogWriter

type LogWriter interface {
	Flush()
	io.Writer
}

func RegisterInitWriterFunc(adapterName string, writeFunc InitLogWriterFunc) {
	writeAdapter[adapterName] = writeFunc
}
