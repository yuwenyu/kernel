package kernel

import (
	"github.com/gin-gonic/gin"
)

func init() {}

type Kernel struct {

}

func New() *Kernel {
	return &Kernel{}
}

func (k *Kernel) SysLog() {
	gin.DisableConsoleColor()
}
