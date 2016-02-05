package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	flag.Parse()
	f, err := GetArgInputFile()
	if err != nil {
		panic(err)
	}
	t, _, err := ReadType(f)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "Type:%s\n", t)
}
