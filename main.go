package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/adityamkk/cc-isa-emulator/core"
	"github.com/adityamkk/cc-isa-emulator/mem"
)

const MAX_CORES = 16

var nCores int
var binPath string
var ram *mem.RandomAccessMemory
var cores []*core.Core

var coreMessages chan struct {
	CoreId int
	Err    error
} = make(chan struct {
	CoreId int
	Err    error
})

func RunCore(coreId int) {
	go func() {
		for {
			coreMessages <- struct {
				CoreId int
				Err    error
			}{CoreId: coreId, Err: (<-cores[coreId].Stop)}
		}
	}()
	go cores[coreId].Run()
}

func main() {
	flag.IntVar(&nCores, "cores", 1, "number of cores to use")
	flag.StringVar(&binPath, "bin", "", "path to the binary file")
	flag.Parse()

	if binPath == "" {
		fmt.Println("Usage: cc-isa-emulator --bin=<filepath> [--cores=N]")
		os.Exit(1)
	}

	if nCores <= 0 || nCores > MAX_CORES {
		fmt.Printf("Error: cores must be between 1 and %d\n", MAX_CORES)
		os.Exit(1)
	}

	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		fmt.Println("Error: provided bianry file does not exist:", binPath)
		os.Exit(1)
	}

	binFile, err := os.Open(binPath)
	if err != nil {
		fmt.Println("Error: failed to open binary file:", err)
		os.Exit(1)
	}
	defer binFile.Close()

	fmt.Println("Hello, World!")
	fmt.Println("File path:", binPath)
	fmt.Println("Cores:", nCores)

	ram = mem.NewRandomAccessMemory(binFile)
	cores = make([]*core.Core, nCores)
	for coreId := range nCores {
		cores[coreId] = core.NewCore(ram)
		RunCore(coreId)
		fmt.Printf("Core %d Started\n", coreId)
	}

	for range nCores {
		msg := (<-coreMessages)
		if msg.Err != nil {
			fmt.Printf("Error (core %d): %v", msg.CoreId, msg.Err)
			os.Exit(1)
		}
		fmt.Printf("Core %d Completed\n", msg.CoreId)
	}
	fmt.Printf("All Cores Complete, stopping emulator...\n")
}
