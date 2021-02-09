package configs

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// parses config file according to it 'Definition'
func (f *FileRaw) Parse() {
	f.read()
	if len(f.Content) == 0 {
		panic(fmt.Errorf("file `" + f.Path + "` content is empty or failed to read"))
	}

	switch f.Definition {

	case "workspace":
		parse(f, WorkspaceYaml{})
	case "application":
		parse(f, ApplicationYaml{})
	default:
		panic(fmt.Errorf("unknow definition"))
	}

	fmt.Printf("%s config parsed : %v \n", f.Definition, f.Config)
	f.Parsed = true
}

func parse(f *FileRaw, i interface{}) {
	if err := yaml.Unmarshal(f.Content, &i); err != nil {
		//return nil, nil, errors.New("failed to Parse workspace config: " + WorkspaceFiles.Path)
		panic(fmt.Errorf("failed to Unmarshall `" + f.Path + "`"))
	}

	f.Config = i
}

func (f *FileRaw) read() {
	content, err := ioutil.ReadFile(f.Path)
	if err != nil {
		panic(fmt.Errorf("failed to read file: " + f.Path))
	}
	if len(content) == 0 {
		panic(fmt.Errorf("file is empty: " + f.Path))
	}

	f.Content = content
}
