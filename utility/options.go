package utility

type Options struct {
	Verbose bool
}

var options Options = Options{
	Verbose: false,
}

func SetOptions(opts Options) {
	options = opts
}
