package kernel

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuwenyu/kernel/ini"
	"github.com/yuwenyu/kernel/help"
)

func init() {}

type Kernel struct {
	Ic *ini.Ic
}

func New() *Kernel {
	return &Kernel{
		Ic:ini.New().SetDir(help.ConfigDir + help.Virgule),
	}
}

func (k *Kernel) SysLog() {
	gin.DisableConsoleColor()

	fn := "storage/logs/wyu_" + time.Now().Format("2006_01_01") + ".log"
	f, _ := os.Create(fn)
	gin.DefaultWriter = io.MultiWriter(f)
}
