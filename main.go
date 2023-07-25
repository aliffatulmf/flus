package main

import (
	"aliffatulmf/flus/move"
	"aliffatulmf/flus/scan"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	target := flag.String("target", "", "target directory")
	unsafe := flag.Bool("unsafe", false, "unsafe mode")
	moveMode := flag.Bool("move", false, "move mode (default: copy)")
	flag.Parse()

	if *target == "" {
		panic("target directory cannot be empty")
	}

	// Scan files in target directory.
	fileMetas, err := scan.Files(*target)
	if err != nil {
		panic(err)
	}

	// Move files to their respective directory.
	for idx, fileMeta := range fileMetas {
		fmt.Println(strings.Repeat("-", 100))
		fmt.Println("# Copy file", idx)

		if err := move.Copy(&fileMeta, !*unsafe); err != nil {
			panic(err)
		}

		fmt.Printf("# file %d => DONE!\n", idx)
		if *moveMode {
			if err := os.Remove(fileMeta.Path); err != nil {
				panic(err)
			}
		}
	}
}
