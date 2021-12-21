# Unmountpoint

**Unmountpoint** is Go library to wait for the detached/unmounted state of a path.

> **DISCLAIMER**: This project is under development and fully experimental mode.

## Installation

```console
$ go get -v github.com/dwisiswant0/unmountpoint/pkg/unmount
```

## Usage

```golang
package main

import (
	"log"
	"os"
	"time"

	"github.com/dwisiswant0/unmountpoint/pkg/unmount"
)

func main() {
	p := "/media/dw1/USB-BUS1" // Mountpoint path to watch
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
```

Basically like [`signal.Notify`](https://pkg.go.dev/os/signal#Notify). If `/media/dw1/USB-BUS1` directory is unmounted; `Wait` assigns channel value then do some stuff _(if any)_.

See [examples](https://github.com/dwisiswant0/unmountpoint/blob/master/examples).

## TODO

See `TODO`.

## Acknowledgements

- [BusKill](https://docs.buskill.in/buskill-app/en/stable/index.html) served as my inspiration.

## License

**Unmountpoint** is distributed under Apache-2.0. See `LICENSE`.