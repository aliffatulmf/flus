package util

import (
	"aliffatulmf/flus/logger"
	"os"
	"path/filepath"
	"runtime"
)

var Platform map[string]string

const (
	StatusStable = "stable"
	StatusBeta   = "beta"
	StatusAlpha  = "alpha"
)

func init() {
	Platform = make(map[string]string)

	Platform["version"] = "v1.0"
	Platform["status"] = StatusStable
	Platform["architecture"] = runtime.GOARCH
	Platform["target_os"] = runtime.GOOS

	exec, err := os.Executable()
	if err != nil {
		logger.Fatal(err)
	}
	Platform["install_dir"] = filepath.Dir(exec)
}
