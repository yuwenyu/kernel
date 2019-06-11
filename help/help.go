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
	return &Help {}
}

func (h *Help) TempCfgEnv(dir string, method string, fn string) string {
	return dir + Virgule + method + Virgule + os.Getenv("WYU_ENV") + Virgule + fn + "." + method
}





