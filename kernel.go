package kernel

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuwenyu/kernel/ini"
	"github.com/yuwenyu/kernel/help"
)

func init() {

}

type Kernel struct {
	Ic *ini.Ic
}

func New() *Kernel {
	log.Println("test kernel")
	icfg := ini.New()
	icfg.Dir = "config" + help.Virgule
	return &Kernel{
		Ic:icfg,
	}
}

func (k *Kernel) SysLog() {
	gin.DisableConsoleColor()

	fn := "storage/logs/wyu_" + time.Now().Format("2006_01_01") + ".log"
	f, _ := os.Create(fn)
	gin.DefaultWriter = io.MultiWriter(f)
}
