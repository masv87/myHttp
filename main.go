package main

import (
	"flag"
)

func main() {
	numParallel := flag.Int("parallel", 10, "parallel workers count")
	flag.Parse()
	urlsList := flag.Args()

	worker := NewWorker(*numParallel, NewClient(), NewUrlModifier())
	worker.Run(urlsList)
}
