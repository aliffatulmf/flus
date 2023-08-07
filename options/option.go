package options

import (
	"fmt"
	"runtime"
	"strconv"
)

const AppName = "Flus"

const (
	StatusStable = "stable" // The program is stable.
	StatusBeta   = "beta"   // The program is beta.
	StatusAlpha  = "alpha"  // The program is alpha.
)

// Version represents a version number with major, minor, and patch components.
type Version struct {
	Major uint8 // The major version number.
	Minor uint8 // The minor version number.
	Patch uint8 // The patch version number.
}

// ToString returns a string representation of the version number in the format "vX.Y.Z".
func (v *Version) ToString() string {
	return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// Get returns the major, minor, and patch components of the version number.
func (v *Version) Get() (uint8, uint8, uint8) {
	return v.Major, v.Minor, v.Patch
}

// Set sets the major, minor, and patch components of the version number.
func (v *Version) Set(major, minor, patch uint8) {
	v.Major = major
	v.Minor = minor
	v.Patch = patch
}

// LessThan returns true if the receiver version number is less than the argument version number.
func (v *Version) LessThan(version Version) bool {
	if v.Major < version.Major {
		return true
	}

	if v.Major > version.Major {
		return false
	}

	if v.Minor < version.Minor {
		return true
	}

	if v.Minor > version.Minor {
		return false
	}

	if v.Patch < version.Patch {
		return true
	}

	if v.Patch > version.Patch {
		return false
	}

	return false
}

// FromString sets the major, minor, and patch components of the version number from a string in the format "vX.Y.Z".
// Returns an error if the string is not in the correct format or if the resulting version number is less than the current version number.
func (v *Version) FromString(version string) error {
	var ver Version

	_, err := fmt.Sscanf(version, "v%d.%d.%d", &ver.Major, &ver.Minor, &ver.Patch)
	if err != nil {
		return err
	}

	if ver.LessThan(*v) {
		return fmt.Errorf("version %s is less than current version %s", ver.ToString(), v.ToString())
	}

	v.Major = ver.Major
	v.Minor = ver.Minor
	v.Patch = ver.Patch

	return nil
}

func NewVersion(major, minor, patch uint8) Version {
	return Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

// Platform represents the operating system and architecture of a system.
type Platform struct {
	OS   string // The name of the operating system.
	Arch string // The name of the architecture.
}

// GlobalOptions represents a set of global options for a program.
type GlobalOptions struct {
	Name      string   // The name of the program.
	Status    string   // The status of the program.
	Version   Version  // The version number of the program.
	Platform  Platform // The platform the program is running on.
	GoVersion string   // The version of Go used to build the program.
	Debug     bool     // Whether debug mode is enabled.
}

// Get returns the value of the specified key in the global options.
// Returns an error if the key is not recognized.
func (g *GlobalOptions) Get(key string) (string, error) {
	switch key {
	case "name":
		return g.Name, nil
	case "status":
		return g.Status, nil
	case "version":
		return g.Version.ToString(), nil
	case "os":
		return g.Platform.OS, nil
	case "arch":
		return g.Platform.Arch, nil
	case "goversion":
		return g.GoVersion, nil
	case "debug":
		return strconv.FormatBool(g.Debug), nil
	default:
		return "", fmt.Errorf("invalid key: %s", key)
	}
}

// Set sets the value of the specified key in the global options.
// Returns an error if the key is not recognized or if the value is invalid.
func (g *GlobalOptions) Set(key, value string) error {
	switch key {
	case "name":
		g.Name = value
	case "status":
		g.Status = value
	case "version":
		version := Version{}
		if err := version.FromString(value); err != nil {
			return err
		}

		g.Version = version
	case "os":
		g.Platform.OS = value
	case "arch":
		g.Platform.Arch = value
	default:
		return fmt.Errorf("invalid key: %s", key)
	}

	return nil
}

func defaultOptions() *GlobalOptions {
	return &GlobalOptions{
		Name:    AppName,
		Status:  StatusAlpha,
		Version: NewVersion(0, 0, 0),
		Platform: Platform{
			OS:   runtime.GOOS,
			Arch: runtime.GOARCH,
		},
		GoVersion: runtime.Version(),
		Debug:     false,
	}
}

var DefaultOptions = defaultOptions()
