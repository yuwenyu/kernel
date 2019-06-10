package help

import (
	"os"
)

const (
	ConfigDir string = "config"

	Virgule string = "/"
	Return string = "\n"
)

type Help struct {}

func New() *Help {
	return &Help{}
}

func (h *Help) TempCfg(dir string, method string, fn string) string {
	return dir + Virgule + os.Getenv("WYU_ENV") + Virgule + method + Virgule + fn + "." + method
}



