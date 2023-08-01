package scan

type Scanner interface {
	Scan() error
}

type Skipper interface {
	Skip() error
	Unskip() error
	IsSkipped() bool
}
