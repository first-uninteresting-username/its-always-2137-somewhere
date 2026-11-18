package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.String("data-path", "data/cities5000.txt", "Relative path to file to convert")
	flag.String("output-path", "cities.json", "Relative path for output ")
	fmt.Print("Hello, world")

}
