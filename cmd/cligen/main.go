package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/first-uninteresting-username/its-always-2137-somewhere/internal/dataset"
)

func main() {
	i := flag.String("data-path", "", "Path to file to convert")
	o := flag.String("output-path", "cities.json", "Path for the output file")
	flag.Parse()
	raw, err := getCities(*i)
	if err != nil {
		log.Fatalf("get cities: %v", err)
	}
	data, err := convertToDataset(raw)
	if err != nil {
		log.Fatalf("convert to dataset: %v", err)
	}

	f, err := os.Create(*o)
	if err != nil {
		log.Fatalf("create file: %v", err)
	}
	defer f.Close()

	var w io.Writer = f
	err = dataset.Write(data, w)
	if err != nil {
		log.Fatalf("write file: %v", err)
		os.Exit(1)
	}
}
