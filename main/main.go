package main

import "time"

func measure(action func()) time.Duration {
	start := time.Now()
	action()
	return time.Since(start)
}

func main() {
}
