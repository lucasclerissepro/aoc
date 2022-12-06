package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var (
	day     = flag.Int("d", 0, "advent of code day")
	year    = flag.Int("y", 2022, "advent of code year")
	session = flag.String("s", "", "session from cookie called `session`")
	dest    = flag.String("o", "", "file where to register the input")
)

func fetchInput(ctx context.Context) ([]byte, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", *year, *day)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: *session})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to do http request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(url +
			"\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func main() {
	flag.Parse()

	log.Printf("Recovering input for day %d of year %d", *day, *year)
	log.Printf("Using session %s", *session)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

  // catch signals
  go func() {
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
    <-sig
    cancel()
  }()

	input, err := fetchInput(ctx)
	if err != nil {
		log.Fatalf("Failed to fetch input: %v", err)
	}

	if err := os.WriteFile(*dest, input, 0644); err != nil {
		log.Fatalf("Failed to write input to destination: %v", err)
	}
}
