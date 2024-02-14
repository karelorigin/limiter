package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

// opt is a struct containing application options.
var opt struct {
	Duration time.Duration
	Rate     int
	Splitter rune
}

// run runs the application.
func run() error {
	if opt.Rate <= 0 {
		return errors.New("-r must be larger than 0")
	}

	if !hasStdin() {
		return errors.New("no stdin was provided")
	}

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)

	t := time.NewTicker(opt.Duration)
	defer t.Stop()

	for range t.C {
		if !process(s, opt.Rate) {
			return nil
		}
	}

	return nil
}

// process reads from a given bufio.Scanner up to n times.
func process(s *bufio.Scanner, n int) (c bool) {
	for i := 0; i < n; i++ {
		if !s.Scan() {
			return false
		}

		err := s.Err()
		if err == nil {
			fmt.Println(s.Text())
		} else {
			log.Println("Error scanning input:", err)
		}
	}

	return true
}

// hasStdin returns whether or not the application was provided standard input.
func hasStdin() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false // Should we panic here?
	}

	return fi.Mode()&os.ModeNamedPipe != 0
}

func main() {
	flag.DurationVar(&opt.Duration, "d", time.Second, "The time to wait after each processed batch. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.")
	flag.IntVar(&opt.Rate, "r", 1, "The max processing rate per unit of time.")
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal("A fatal error occured: ", err)
	}
}
