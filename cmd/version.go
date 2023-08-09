package cmd

import (
	"aliffatulmf/flus/options"
	"aliffatulmf/flus/util"
	"fmt"
)

func Version() {
	opt := options.DefaultOptions
	opt.Status = util.StatusStable
	opt.Version.Set(2, 1, 0)

	target := fmt.Sprintf("%s_%s-%s", opt.Platform.Arch, opt.Platform.OS, opt.Status)
	fmt.Printf("Flus version %s %s\n", opt.Version.ToString(), target)
	fmt.Printf("Go: %s\n", opt.GoVersion)
}
