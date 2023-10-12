package main

import (
	"fmt"
	"log"
	"os"
	"poc/shared/generic"

	"github.com/gocarina/gocsv"
)

var (
	csvFilepath = "../../../../shared/assets/csv/file.csv"
)

func main() {
	file, err := os.Open(csvFilepath)
	if err != nil {
		log.Fatalf("failed to read file: %v\n", err)
	}

	var ss []generic.Struct

	if err := gocsv.UnmarshalFile(file, &ss); err != nil {
		log.Fatalf("failed to unmarshal file: %v\n", err)
	}

	for _, s := range ss {
		fmt.Printf("s.Foo: %v\n", s.Foo) // s.Foo: bar
		fmt.Printf("s.Baz: %v\n", s.Baz) // s.Baz: 1
	}
}
