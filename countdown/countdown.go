package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const startCount = 3

func Countdown(out io.Writer) {
	for i := startCount; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
