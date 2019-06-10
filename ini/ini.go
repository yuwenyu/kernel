package ini

import (
	"fmt"
	"path/filepath"

	"github.com/yuwenyu/kernel/help"
	"gopkg.in/ini.v1"
)

const (
	sIni string = "ini"
)

func init() {}

type ip struct {
	section string
}

type Ic struct {
	dir string
	h *help.Help
	ip ip
	cfg *ini.File
}

func New() *Ic {
	h := help.New()
	return &Ic{
		h:h,
	}
}

func (i *Ic) SetDir(dir string) *Ic{
	i.dir = dir
	return i
}

func (i *Ic) SetIp(section string) *Ic {
	i.ip.section = section
	return i
}

func (i *Ic) Loading() *Ic {
	fns, err := filepath.Glob(i.h.TempCfg(i.dir, sIni, "*"))
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err.Error()))
	}

	if len(fns) == 0 {
		panic(fmt.Sprintf("Error: %s", err.Error()))
	}

	arrFns := make([]interface{}, len(fns))
	for k, fn := range fns {
		arrFns[k] = fn
	}

	cfg, err := ini.Load(arrFns[0], arrFns ...)
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err.Error()))
	}

	i.cfg = cfg

	return i;
}

/**
 * Todo: Initialize the ic.IniFile in the first place
 *
 * Method:Get
 * kernel.Ic.Loading()
 * kernel.Ic.K("Key")
 * kernel.Ic.K("Key").In("str", []string{"str1","str2"})
 * kernel.Ic.K("Key").MustInt(9999)
 * kernel.Ic.K("Key").MustBool(false)
 *
 * Method:Set
 * kernel.Ic.K("Key").SetValue("SetValue")
 * i.cfg.SaveTo(Path)
 * cfg.Section("").Key("app_mode").SetValue("production")
 * cfg.SaveTo(path+"my.ini.local")
 */
func (i *Ic) K(key string) *ini.Key {
	if i.cfg == nil {
		panic("Error nil")
	}

	return i.cfg.Section(i.ip.section).Key(key)
}

func (i *Ic) SaveTo(key string, val string, fn string) *Ic {
	if i.cfg == nil {
		panic("Error nil")
	}

	i.K(key).SetValue(val)
	i.cfg.SaveTo(i.h.TempCfg(i.dir, sIni, fn))
	return i
}


