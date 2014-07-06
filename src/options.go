package graph

type Options struct {
	Verbose bool
}

var options Options

func SetOptions(opts Options) {
	options = opts
}
