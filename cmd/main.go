package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kindacommander/sf-encoder/internal/counter"
)

func main() {
	str := strings.Join(os.Args[1:], "")
	fmt.Println(counter.FreqCount(str))
}
