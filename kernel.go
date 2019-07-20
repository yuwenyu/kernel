/**
 * Copyright 2019 YuwenYu.  All rights reserved.
**/

package kernel

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/multitemplate"
)

func init() {}

type Kernel struct {
	Ini *ini
}

func New() *Kernel {
	var ini INI = &ini{}
	return &Kernel{
		Ini:ini.Loading(),
	}
}

func (k *Kernel) Run() *gin.Engine {
	k.ginInitialize()

	r := gin.Default()
	r  = k.ginTemplateStatic(r)

	return r
}

func (k *Kernel) GinTemplate() multitemplate.Renderer {
	var templates templates = &template{
		directory:k.Ini.K(
			MapConfLists[ConfTemplates][1],
			MapConfParam[MapConfLists[ConfTemplates][1]][0],
		).String(),
	}
	return templates.Tpl()
}

func (k *Kernel) GinTemplateLoadByView(skeleton string, view string) []string {
	var templates templates = &template{
		directory:k.Ini.K(
			MapConfLists[ConfTemplates][1],
			MapConfParam[MapConfLists[ConfTemplates][1]][0],
		).String(),
	}
	return templates.LoadingTPL(skeleton, view)
}

func (k *Kernel) ginInitialize() {
	bLog, _ := k.Ini.K(
		MapConfLists[ConfCommons][2],
		MapConfParam[MapConfLists[ConfCommons][2]][0],
	).Bool()
	if bLog {
		gin.DisableConsoleColor()

		logRoot := k.Ini.K(
			MapConfLists[ConfCommons][2],
			MapConfParam[MapConfLists[ConfCommons][2]][1],
		).String()
		if logRoot == "" {logRoot = KCommLogRoot}

		_, err := os.Stat(logRoot)
		if err != nil {
			panic(err.Error())
		}

		logPrefixFN := k.Ini.K(
			MapConfLists[ConfCommons][2],
			MapConfParam[MapConfLists[ConfCommons][2]][2],
		).String()
		if logPrefixFN == "" {logPrefixFN = KCommLogPrefixFn}

		fn := logRoot + StrVirgule + logPrefixFN + StrUL + MapTimeFormat["DDF"] + ".log"
		f, _ := os.Create(fn)
		gin.DefaultWriter = io.MultiWriter(f)
	} else {
		gin.ForceConsoleColor()
	}
}

func (k *Kernel) ginTemplateStatic(r *gin.Engine) *gin.Engine {
	bTplStatic, _ := k.Ini.K(
		MapConfLists[ConfCommons][1],
		MapConfParam[MapConfLists[ConfCommons][1]][1],
	).Bool()
	if bTplStatic {
		static 		:= k.Ini.K(
			MapConfLists[ConfTemplates][0],
			MapConfParam[MapConfLists[ConfTemplates][0]][0],
		).String()
		if static == "" {static = KTempStatic}

		staticFile 	:= k.Ini.K(
			MapConfLists[ConfTemplates][0],
			MapConfParam[MapConfLists[ConfTemplates][0]][1],
		).String()
		if staticFile == "" {staticFile = KTempStaticFile}

		r.Static("/assets", static)
		r.StaticFile("/favicon.ico", staticFile)
	}

	return r
}
