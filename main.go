package main

import (
	"aliffatulmf/flus/args"
	"aliffatulmf/flus/cmd"
	"aliffatulmf/flus/logger"
)

func main() {
	arg := args.InitArgs()

	switch {
	case arg.Version:
		cmd.Version()
		return
	case arg.Target == "":
		logger.Fatal("target directory is required")
	}

	cmd.DoScan(arg)
}
