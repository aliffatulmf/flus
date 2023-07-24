package main

import (
	pkg "aliffatulmf/flus/move"
	"aliffatulmf/flus/scan"
	"flag"
	"fmt"
)

func main() {
	target := flag.String("target", "", "target directory")
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
		if err := pkg.SafeCopy(&fileMeta); err != nil {
			panic(err)
		}
		fmt.Println("Moving", fileMeta.Info.Name())
	}
}
