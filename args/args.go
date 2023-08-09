package args

import (
	"flag"
	"os"
)

type Value struct {
	Target  string `name:"target"`
	Move    bool   `name:"move"`
	Version bool   `name:"version"`
}

func NewArgs() *Value {
	var args = new(Value)

	fl := flag.NewFlagSet("Flus", flag.ExitOnError)
	fl.StringVar(&args.Target, "target", "", "Target directory to scan")
	fl.BoolVar(&args.Move, "move", false, "Use move mode instead of copy")
	fl.BoolVar(&args.Version, "version", false, "Print version information")
	fl.Parse(os.Args[1:])

	return args
}
