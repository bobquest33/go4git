package go4git

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"os"
)

/////// Begin changes by Ashani ///////////////////

func GetArgInputFile() (*os.File, error) {
	args := flag.Args()
	if len(args) > 0 {
		return os.Open(args[0])
	} else {
		return os.Stdin, nil
	}
}

func GenSHA1(in io.Reader, objType string) ([]byte, error) {
	buf := new(bytes.Buffer)
	length, err := buf.ReadFrom(in)
	if err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("%s %d", objType, length)

	h := sha1.New()
	io.WriteString(h, prefix)
	h.Write([]byte{0})
	buf.WriteTo(h)
	return h.Sum(nil), nil
}

func UnzlibToBuffer(in io.Reader) ([]byte, error) {
	var b bytes.Buffer
	err := Unzlib(in, &b)
	return b.Bytes(), err
}

func Unzlib(in io.Reader, out io.Writer) error {
	r, err := zlib.NewReader(in)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, r)
	r.Close()
	return err
}

func Zlib(in io.Reader, out io.Writer) error {
	w := zlib.NewWriter(out)
	_, err := io.Copy(w, in)
	w.Close()
	return err
}

///////////////////End changes by Ashani////////////
