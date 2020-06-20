package main

import (
	"flag"
	"fmt"
)

func main() {
	var count int
	flag.IntVar(&count, "count", 0, "Count Members")
	flag.Parse()
	fmt.Println(count)
}
