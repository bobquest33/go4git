package main

import (
	"os"
	"fmt"
	"flag"
	"path/filepath"
	"errors"
)

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}

func folderExists(path string) bool {
	if s, err := os.Stat(path); err == nil {
		return s.IsDir()
	} else {
		return false
	}
}

func IsRepo(root string) bool {
	dotGit := filepath.Join(root, ".git")
	return folderExists(dotGit)
}

func IsBareRepo(root string) bool {
	objFolder := filepath.Join(root, "objects")
	headFile := filepath.Join(root, "HEAD")

	return folderExists(objFolder) && fileExists(headFile)
}

func GitDir(root string) (string, error) {
	switch {
	case IsRepo(*repoRoot):
		return filepath.Join(root, ".git"), nil
	case IsBareRepo(*repoRoot):
		return root, nil
	default:
		return "", errors.New("not a git repository")
	}
}

func GetObjPath(sha string, root string) (string, error) {
	gd, err := GitDir(root)

	if err != nil {
		return "", err
	}

	d := sha[0:2]
	f := sha[2:]
	return filepath.Join(gd, d, f), nil
}


var repoRoot = flag.String("d", ".", "path to repository root")

func main() {
	flag.Parse()

	for _, sha := range os.Args[1:] {
		p, err := GetObjPath(sha, *repoRoot)

		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
		} else {
			fmt.Fprintln(os.Stdout, p)
		}
	}
}
