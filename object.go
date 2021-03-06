package go4git

// holds all functions common to all object types

import (
	"bytes"
	"os"
	"strconv"
	"strings"
)

type Object struct {
	id      string
	objType string
	size    int
	data    []byte
}

func parseHeader(buff *bytes.Buffer) (int, string, error) {
	header, err := buff.ReadString(0)

	if err != nil {
		return 0, "", err
	}

	xs := strings.Split(header, " ")
	objType, objSize := xs[0], xs[1]

	objSize = objSize[:len(objSize)-1] // remove trailing null

	size, err := strconv.Atoi(objSize)

	if err != nil {
		return 0, objType, err
	}

	return size, objType, nil
}

func ReadObjectType(in *os.File) (int, string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(in)
	if err != nil {
		return 0, "", err
	}
	return parseHeader(buf)
}
