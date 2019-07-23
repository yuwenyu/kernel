/**
 * Copyright 2019 YuwenYu.  All rights reserved.
**/

package kernel

import (
	"os"
	"time"
)

const (
	SysName		string = "WYU"
	SysOsEnv	string = "WYU_ENV"
	sysLoadLocation	string = "Asia/Shanghai"

	StrCD      			string = "config" // directory of config (string)
	StrVirgule 			string = "/"
	StrDoubleVirgule	string = "//"
	StrUL      			string = "_"
	StrDOT     			string = "."
	StrColon   			string = ":"
	StrHttp				string = "http"
	StrHttps			string = "https"

	SysTimeFormat string = "2006-01-02 00:00:00"
	SysDateFormat string = "2006-01-02"
	DirDateFormat string = "20060102" // Directory Time Format

	ConfDB          string = "db"
	ConfRedis       string = "redis"
	ConfCommons     string = "commons"
	ConfTemplates	string = "templates"

	KRedisAddr		string = "127.0.0.1:6379"
	KRedisDB		int = 0
	KRedisPoolSize	int = 10

	KCommLogRoot		string = "storage" + StrVirgule + "logs" + StrVirgule
	KCommLogPrefixFn	string = SysName

	KDbPort		int = 3306
	KDbMaxOpen	int = 50
	KDbMaxIdle	int = 200
	KDbShowedSQL	bool = false
	KDbCachedSQL	bool = false

	KTempStatic		string = "./resources/assets"
	KTempStaticFile	string = "./resources/favicon.ico"
)

var (
	SysTimeLocation, _                   	= time.LoadLocation(sysLoadLocation)
	MapTimeFormat      map[string]string 	= map[string]string{
		"STF": time.Now().Format(SysTimeFormat),
		"SDF": time.Now().Format(SysDateFormat),
		"DDF": time.Now().Format(DirDateFormat),
	}
	MapConfLists map[string]map[int]string	= map[string]map[int]string{
		ConfDB: map[int]string{
			0:"db_engine",
		},
		ConfRedis:		map[int]string{
			0:"redis",
		},
		ConfCommons:	map[int]string{
			0:"common_translate",
			1:"common_cfg",
			2:"common_log",
		},
		ConfTemplates:	map[int]string{
			0:"template_statics",
			1:"template_root",
		},
	}
	MapConfParam map[string]map[int]string	= map[string]map[int]string{
		MapConfLists[ConfDB][0]:		map[int]string{
			0:"driver",
			1:"host",
			2:"port",
			3:"table",
			4:"username",
			5:"password",
			6:"max_open",
			7:"max_idle",
			8:"showed_sql",
			9:"cached_sql",
		},
		MapConfLists[ConfRedis][0]:		map[int]string{
			0:"address",
			1:"password",
			2:"db",
			3:"pool_size",
		},
		MapConfLists[ConfCommons][0]: 	map[int]string{
			0:"origin_language",
			1:"base_path",
			2:"file_map",
		},
		MapConfLists[ConfCommons][1]:	map[int]string{
			0:"template_status",
			1:"template_static_status",
		},
		MapConfLists[ConfCommons][2]:	map[int]string{
			0:"log_status",
			1:"log_root",
			2:"log_fn_prefix",
		},
		MapConfLists[ConfTemplates][0]:	map[int]string{
			0:"static",
			1:"static_file",
		},
		MapConfLists[ConfTemplates][1]:	map[int]string{
			0:"directory",
			1:"directory_view",
			2:"resources",
		},
	}
)

type Helpers interface {
	TempCfgEnv(fn string) string
}

type Helper struct {
	directory string
	method    string
}

var _ Helpers = &Helper{}

func (h *Helper) TempCfgEnv(fn string) string {
	if h.directory == "" || h.method == "" {
		panic("Error Empty Helper ...")
	}

	strEnv := os.Getenv(SysOsEnv)
	if strEnv == "" {
		panic("Error ENV(WYU_ENV) Helper ...")
	}

	return h.directory + StrVirgule + h.method + StrVirgule + fn + StrDOT + strEnv + StrDOT + h.method
}
