package main

import (
	"fmt"
	"github.com/atotto/clipboard"
)

const (
	ZWSP = "\u200B"
)

func main() {
	clipboard.WriteAll(ZWSP)
	fmt.Println("Zero-width space (U+200B) copied to clipboard!")
}
