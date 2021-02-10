package configs

import "fmt"

func write() {

}

func (f *RawConfig) update() {
	if !f.Parsed {
		panic(fmt.Errorf("can't update unparsed file"))
	}
	//f.Parse()
	write()
}
