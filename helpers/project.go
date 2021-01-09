package helpers

import (
	"fmt"
)

// Project contains name and paths to projects.
type Project struct {
	AbsolutePath string
	Name         string
}

func (p *Project) Create() error {
	fmt.Printf("%v", p)
	if err := createDirIfNotExists(p.AbsolutePath); err != nil {
		return err
	}
	// Create project files ...

	return nil
}
