package args

import (
	"flag"
	"os"
)

type Args struct {
	Target  string
	Move    bool
	Buffer  uint
	Version bool
}

func InitArgs() *Args {
	var args = new(Args)

	fl := flag.NewFlagSet("Flus", flag.ExitOnError)
	fl.StringVar(&args.Target, "target", "", "Target directory to scan")
	fl.BoolVar(&args.Move, "move", false, "Use move mode instead of copy")
	fl.UintVar(&args.Buffer, "buffer", 64*1024, "Buffer size to use when copying files")
	fl.BoolVar(&args.Version, "version", false, "Print version information")
	fl.Parse(os.Args[1:])

	return args
}

// func init() {
// 	fl := flag.NewFlagSet("Flus", flag.ExitOnError)
// 	fl.StringVar(&args.Target, "target", "", "Target directory to scan")
// 	fl.BoolVar(&args.Move, "move", false, "Use move mode instead of copy")
// 	fl.UintVar(&args.Buffer, "buffer", 64*1024, "Buffer size to use when copying files")
// 	fl.BoolVar(&args.Version, "version", false, "Print version information")
// 	fl.Parse(os.Args[1:])
// }

// func DoScan() {
// 	if args.Target == "" {
// 		logger.Error("target directory is not specified")
// 		os.Exit(1)
// 	}

// 	args.Target, err := filepath.Abs(args.Target)
// 	if err != nil {
// 		logger.Errorf("failed to get absolute path: %v", err)
// 	}

// 	metas, err := scan.All(args.Target)
// 	if err != nil {
// 		logger.Errorf("failed to scan files: %v", err)
// 	}

// 	for _, meta := range *metas {
// 		src := meta.Info.Name()
// 		dst := filepath.Join(strings.ToLower(meta.Destination), meta.Info.Name())
// 		logger.Infof("transferring file: %s => %s", src, dst)

// 		if err := processFile(args.Target, meta); err != nil {
// 			logger.Error(err)
// 			continue
// 		}
// 	}
// }

// func Command() {
// 	if version {
// 		Version()
// 		os.Exit(0)
// 	}

// 	if target == "" {
// 		logger.Error("target directory is not specified")
// 		os.Exit(1)
// 	}

// 	target, err := filepath.Abs(target)
// 	if err != nil {
// 		logger.Errorf("failed to get absolute path: %v", err)
// 	}

// 	metas, err := scan.All(target)
// 	if err != nil {
// 		logger.Errorf("failed to scan files: %v", err)
// 	}

// 	for _, meta := range *metas {
// 		src := meta.Info.Name()
// 		dst := filepath.Join(strings.ToLower(meta.Destination), meta.Info.Name())
// 		logger.Infof("transferring file: %s => %s", src, dst)

// 		if err := processFile(target, meta); err != nil {
// 			logger.Error(err)
// 			continue
// 		}
// 	}
// }

// func processFile(target string, meta scan.Metadata) error {
// 	dir := filepath.Join(target, meta.Destination)
// 	if _, err := os.Lstat(dir); err != nil {
// 		if os.IsNotExist(err) {
// 			if err := os.MkdirAll(dir, fs.ModeDir); err != nil {
// 				return fmt.Errorf("failed to create destination directory: %v", err)
// 			}
// 		} else {
// 			return fmt.Errorf("failed to get destination directory info: %v", err)
// 		}
// 	}

// 	if unsafe {
// 		logger.Infof("marking file as unsafe: %s", meta.Info.Name())
// 		if err := meta.Unsafe(); err != nil {
// 			return fmt.Errorf("failed to mark file as unsafe: %v", err)
// 		}
// 	}

// 	if move {
// 		if err := meta.Move(); err != nil {
// 			return fmt.Errorf("failed to mark file as moved: %v", err)
// 		}
// 	}

// 	if err := io.Relocate(dir, meta, buffSize); err != nil {
// 		return fmt.Errorf("failed to transfer file: %v", err)
// 	}

// 	return nil
// }
