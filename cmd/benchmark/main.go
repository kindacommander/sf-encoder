package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	rand_strings "github.com/kindacommander/sf-encoder/internal/rand-strings"
	"github.com/kindacommander/sf-encoder/internal/tree"
)

var (
	minLen     = flag.Int("min", 15, "Minimal length of the generated string.")
	maxLen     = flag.Int("max", 35, "Maximum length of the generated string.")
	strsNumber = flag.Int("n", 10000, "Number of strings to generate.")
)

func main() {
	flag.Parse()

	strs := make(chan string, 30)
	go stringsGenerator(strs)

	fmt.Printf(" > Benchmark:\nNumber of strings: %d\nStrings min length: %d\nStrings max length: %d\n",
		*strsNumber, *minLen, *maxLen)

	fmt.Printf("\nWith concurrency...")

	start := time.Now()

	for str := range strs {
		tree.BuildCodeTree(str)
	}

	pureTotalTiming := time.Since(start)

	strs = make(chan string, 20)

	go stringsGenerator(strs)

	fmt.Printf("\r                     ")
	fmt.Printf("\rWith concurrency:\n")

	fmt.Printf("Total: %v\n", pureTotalTiming)
	fmt.Printf("Average: %v\n", pureTotalTiming/time.Duration(*strsNumber))

	fmt.Print("\nWithout concurrency...")

	start = time.Now()

	for str := range strs {
		tree.BuildCodeTreeUnparallel(str)
	}

	pureTotalTiming = time.Since(start)

	fmt.Print("\r                      ")
	fmt.Printf("\rWithout concurrency:\n")

	fmt.Printf("Total: %v\n", pureTotalTiming)
	fmt.Printf("Average: %v\n", pureTotalTiming/time.Duration(*strsNumber))
}

func stringsGenerator(strs chan<- string) {
	for i := 0; i < *strsNumber; i++ {
		var seededRand *rand.Rand = rand.New(
			rand.NewSource(time.Now().UnixNano()))
		strs <- rand_strings.String(seededRand.Intn(*maxLen-*minLen) + *minLen)
	}
	close(strs)
}
