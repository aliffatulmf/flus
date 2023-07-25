package main

import (
	"aliffatulmf/flus/move"
	"aliffatulmf/flus/scan"
	"aliffatulmf/flus/util"
	"flag"
	"fmt"
)

func main() {
	target := flag.String("target", "", "target directory")
	unsafe := flag.Bool("unsafe", false, "unsafe mode")
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
	for _, fileMeta := range fileMetas {
		if *unsafe {
			fmt.Print("Unsafe copying ", util.TrimText(fileMeta.Info.Name()))
		} else {
			fmt.Print("Safe copying ", util.TrimText(fileMeta.Info.Name()))
		}

		if err := move.Copy(&fileMeta, !*unsafe); err != nil {
			panic(err)
		}

		fmt.Print(" => DONE!\n")
	}
}
