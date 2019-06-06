package kernel

import (
	"io"
	"os"
	"time"

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

	fn := "storage/logs/wyu_" + time.Now().Format("2006_01_01") + ".log"
	f, _ := os.Create(fn)
	gin.DefaultWriter = io.MultiWriter(f)
}
