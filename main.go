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
	fl := flag.NewFlagSet("Flus", flag.ExitOnError)
	target := fl.String("target", "", "Specify the target directory")
	unsafe := fl.Bool("unsafe", false, "Enable unsafe mode (skipping file verification)")
	moveMode := fl.Bool("move", false, "Use move mode instead of copy (default: copy)")
	fl.Parse(os.Args[1:])

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
		fmt.Println("# Moving file", idx+1)

		if err := move.Copy(&fileMeta, *moveMode, !*unsafe); err != nil {
			fmt.Println("# Skipping file", idx+1, "due to error:", err)
		}

		fmt.Printf("# File %d => DONE!\n", idx+1)
		// if *moveMode {
		// 	if err := os.Remove(fileMeta.Path); err != nil {
		// 		panic(err)
		// 	}
		// }
	}
}
