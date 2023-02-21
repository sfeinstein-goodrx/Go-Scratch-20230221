package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sig, err := fileSHA1("sha1/http.log.gz")
	if err != nil {
		log.Fatalf("ERROR %e", err)
	}
	log.Println("sig: ", sig)
}

// $ cat http.log.gz | gunzip | sha1sum
// exercise: If the file does not end with .gz, don't decompress it
// Hint: strings.HasSuffix
func fileSHA1(fileName string) (string, error) {
	if err := validateFilename(fileName); err != nil {
		return "", err
	}

	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	r, err := gzip.NewReader(file)
	if err != nil {
		return "", fmt.Errorf("%q: not gzip - %s", fileName, err)
	}

	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("%q: can't copy - %s", fileName, err)
	}
	data := w.Sum(nil)

	return fmt.Sprintf("%x", data), nil
}

func validateFilename(fileName string) error {
	if !strings.HasSuffix(fileName, ".gz") {
		return fmt.Errorf("%q: wrong file extension", fileName)
	}
	return nil
}
