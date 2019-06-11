package kernel

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuwenyu/kernel/ini"
	"github.com/yuwenyu/kernel/help"
	"github.com/yuwenyu/kernel/template"
)

func init() {}

type Kernel struct {
	G *gin.Engine
	Ic *ini.Ic
	tpl *template.Template
}

func New() *Kernel {
	return &Kernel{
		G:gin.Default(),
		Ic:ini.New().SetDir(help.ConfigDir + help.Virgule).Loading(),
		tpl:template.New(),
	}
}

func (k *Kernel) Run(addr string) {
	sLog, _ := k.Ic.SetIp("common_cfg").K("log_status").Bool()
	if sLog {
		k.log()
	}

	sTpl, _ := k.Ic.SetIp("common_cfg").K("template_status").Bool()
	if sTpl {
		root := k.Ic.SetIp("template_root").K("directory").String()
		k.G.HTMLRender = k.tpl.SetDir(root).Tpl()
		k.static()
	}

	k.G.Run(addr)
}

func (k *Kernel) log() {
	gin.DisableConsoleColor()

	fn := "storage/logs/wyu_" + time.Now().Format("2006_01_01") + ".log"
	f, _ := os.Create(fn)
	gin.DefaultWriter = io.MultiWriter(f)
}

func (k *Kernel) static() {
	static := k.Ic.SetIp("template_statics").K("static").String()
	staticFile := k.Ic.SetIp("template_statics").K("static_file").String()

	k.G.Static("/assets", static)
	k.G.StaticFile("/favicon.ico", staticFile)
}
