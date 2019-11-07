package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func spin() {
	fh, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	for {
		fmt.Fprintf(fh, ".")
		time.Sleep(500000 * time.Nanosecond)
	}
}

func main() {
	numCores := os.Getenv("NUM_CORES")
	speed := os.Getenv("BURN_SPEED")
	var nc int
	var err error
	if numCores == "" || numCores == "0" {
		nc = runtime.NumCPU()
	} else {
		nc, err = strconv.Atoi(numCores)
		if err != nil {
			panic(err)
		}
	}

	var rt int
	switch speed {
	case "light":
		rt = 1
	case "rediculous":
		rt = 5
	case "ludicrous":
		rt = 10
	default:
		speed = "light"
		rt = 1 // default to light
	}

	if nc > runtime.NumCPU() {
		fmt.Printf("downgrading NUM_CORES=%d to equal max cores available of %d", nc, runtime.NumCPU())
		nc = runtime.NumCPU()
	}
	fmt.Printf("Using %d cpu cores at %s speed\n", nc, speed)
	runtime.GOMAXPROCS(nc)

	for i := 0; i < nc; i++ {
		for t := 0; t < rt; t++ {
			go spin()
		}
	}

	// sleep for ever
	for {
		time.Sleep(10 * time.Second)
	}
}
