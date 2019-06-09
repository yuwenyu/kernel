package ini

import (
	"fmt"
	"os"
	"github.com/yuwenyu/kernel/help"
	//"path/filepath"
	//"gopkg.in/ini.v1"
)

const (
	sIni string = "ini"
	file string = "config" + "/"
)

func init() {}

type Ic struct {
	Dir string
}

func New() *Ic {
	return &Ic{}
}

func (i *Ic) Loading() {
	fn := i.Dir + os.Getenv("WYU_ENV") + help.Virgule + sIni + help.Virgule + "*." + sIni
	//cfgs, err := filepath.Glob(fn)
	//if (err != nil) {
	//	panic(fmt.Sprintf("Configure Error: %s", err.Error()))
	//}
	//
	//for _, cfg = range cfgs {
	//
	//}
	fmt.Println(fn)
}


