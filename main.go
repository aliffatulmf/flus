package main

import (
	"aliffatulmf/flus/args"
	"aliffatulmf/flus/cmd"
	"aliffatulmf/flus/logger"
)

func main() {
	av := args.NewArgs()

	switch {
	case av.Version:
		cmd.Version()
		return
	case av.Target == "":
		logger.Fatal("target directory is required")
	}

	cmd.DoScan(av)
}
