package configs

import "fmt"

func write() {

}

func (cfg *RawConfig) update() {
	if !cfg.Parsed {
		panic(fmt.Errorf("can't update unparsed file"))
	}
	//f.Parse()
	write()
}
