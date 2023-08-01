package io

type Skipper interface {
	Skip() error
	Unskip() error
	IsSkipped() bool
}

type Mover interface {
	Move() error
	Unmove() error
	IsMoved() bool
}

type Safer interface {
	Safe() error
	Unsafe() error
	IsSafe() bool
}
