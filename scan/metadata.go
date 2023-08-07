package scan

import (
	"fmt"
	"io/fs"
)

type Metadata struct {
	Source      string      // Source file path
	Destination string      // Destination folder
	Info        fs.FileInfo // File info
	unsafe      bool        // Unsafe copy
	move        bool        // Copy and remove source
}

func (m *Metadata) Move() error {
	if m.move {
		return fmt.Errorf("file %s is already moved", m.Info.Name())
	}
	m.move = true
	return nil
}

func (m *Metadata) Unmove() error {
	if !m.move {
		return fmt.Errorf("file %s is not moved", m.Info.Name())
	}
	m.move = false
	return nil
}

func (m *Metadata) IsMoved() bool {
	return m.move
}

func (m *Metadata) Unsafe() error {
	if m.unsafe {
		return fmt.Errorf("file %s is already marked unsafe", m.Info.Name())
	}
	m.unsafe = true
	return nil
}

func (m *Metadata) Safe() error {
	if !m.unsafe {
		return fmt.Errorf("file %s is not marked unsafe", m.Info.Name())
	}
	m.unsafe = false
	return nil
}

func (m *Metadata) IsUnsafe() bool {
	return m.unsafe
}
