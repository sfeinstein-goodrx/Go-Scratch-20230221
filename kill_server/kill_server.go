package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	if err := killServer("kill_server/app.pid"); err != nil {
		fmt.Printf("error type: %T\n", err)
		var perr *fs.PathError
		if errors.As(err, &perr) {
			fmt.Println("> file not found")
		}
		for e := err; e != nil; e = errors.Unwrap(e) {
			fmt.Printf("> %s\n", e)
		}
		log.Fatalf("error: %s", err)
	}
}

func killServer(pidFile string) error {
	// idion: acquire resource, check for error, defer release
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	// defer means file.Close() is called when killServer exits
	// deferred are called in reverse order
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("%q: can't close - %s", pidFile, err)
		}
	}()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return fmt.Errorf("%q: bad pid - %s", pidFile, err)
	}

	fmt.Printf("killing %d\n", pid) // simulate actually killing process
	if err := os.Remove(pidFile + "fake"); err != nil {
		log.Printf("%q: can't delete - %s", pidFile, err)
	}
	return nil
}
