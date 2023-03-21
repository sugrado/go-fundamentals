package interfaces

import "fmt"

type Logger interface {
	Log(text string)
}

type FileLogger struct {
}

type DbLogger struct {
}

func (f FileLogger) Log(text string) {
	fmt.Print(text + " logged by file logger")
}

func (db DbLogger) Log(text string) {
	fmt.Print(text + " logged by database logger")
}

func Demo1() {
	var logger Logger = FileLogger{}
	logger.Log("Info")
}
