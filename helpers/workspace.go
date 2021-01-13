package helpers

type DApp struct {
	RelativePath string
}

// Workspace contains name and paths to dApps.
type Workspace struct {
	AbsolutePath string
	Name         string
	Apps         []DApp
}

func (p *Workspace) Create() error {
	if err := createDirIfNotExists(p.AbsolutePath); err != nil {
		return err
	}
	// Create project files ...

	return nil
}
