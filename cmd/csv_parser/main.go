package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	var filename = flag.String("f", "", "File name to be parsed")
	flag.Parse()
	if *filename == "" {
		fmt.Println("Enter the file name")
		os.Exit(1)
	}
	fd, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
	}
	fileReader := csv.NewReader(fd)
	records, err := fileReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(records)
	defer fd.Close()
}
