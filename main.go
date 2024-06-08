package main

import (
	"time"

	"github.com/joshryandavis/songbird/cmd"
)

func main() {
	start := time.Now()
	cmd.Main()
	elapsed := time.Since(start)
	println("took:", elapsed.String())
}
