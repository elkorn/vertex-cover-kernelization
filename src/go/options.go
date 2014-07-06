package graph

import "log"

type Options struct {
	Verbose bool
}

var options Options

func SetOptions(opts Options) {
	options = opts
}

func Debug(msg string) {
	// func (message ...interface{}) { fmt.Println(...message) }
	if options.Verbose {
		log.Print(msg)
	}
}
