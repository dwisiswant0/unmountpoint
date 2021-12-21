package main

import (
	"log"
	"os"
	"time"

	"path/filepath"

	"github.com/dwisiswant0/unmountpoint/pkg/unmount"
)

func main() {
	p := getCurrentPath()
	c := make(chan bool, 1)
	e := unmount.Wait(c, p)

	if e != nil {
		log.Fatal(e)
	}

	go func() {
		<-c
		// Unmounted!
		// Do stuff e.g. rm -rf /

		log.Println("Path unmounted!")
		os.Exit(1)
	}()

	log.Printf("Wait for %s path to detach/unmounted...\n", p)
	for {
		time.Sleep(10 * time.Second)
	}
}

func getCurrentPath() string {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	path, err := filepath.Abs(filepath.Dir(exe))
	if err != nil {
		log.Fatal(err)
	}

	return path
}
