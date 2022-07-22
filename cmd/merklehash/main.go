package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hipper/merklehash/internal/merkle"
)

const usageTemplate = `
                      _    _      _   _           _
 _ __ ___   ___ _ __| | _| | ___| | | | __ _ ___| |__
| '_   _ \ / _ \ '__| |/ / |/ _ \ |_| |/ _  / __| '_ \
| | | | | |  __/ |  |   <| |  __/  _  | (_| \__ \ | | |
|_| |_| |_|\___|_|  |_|\_\_|\___|_| |_|\__,_|___/_| |_|

Usage:
	merklehash [-flag]`

var path string

func init() {
	flag.StringVar(&path, "path", "", "path to the input file")
}

func printUsage() {
	fmt.Println(usageTemplate)
	fmt.Println()
	fmt.Println("Available flags:")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if path == "" {
		printUsage()
		return
	}

	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Unable to read file: %s", err)
	}

	tree := merkle.NewTreeHash(&merkle.SHA256Hasher{})
	root, err := tree.CalculateRoot(bytes.Split(f, []byte("\n")))
	if err != nil {
		log.Fatalf("Failed to calculate: %s", err)
	}

	fmt.Printf("calculated root hash: %s\n", root)

	os.Exit(0)
}
