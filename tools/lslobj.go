package main

import (
	"flag"
	"fmt"
	"os"
)

var repoRoot = flag.String("d", ".", "path to repository root")

func main() {
	flag.Parse()

	repo, err := NewRepository(*repoRoot)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

	objs, err := repo.LooseObjects()

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

	for _, obj := range objs {
		fmt.Fprintln(os.Stdout, obj)
	}
}
