package kernel

import (
	"github.com/syyongx/ii18n"
	"strings"
)

type I18N interface {
	T(cg string, key string, ln string) string
}

type i18n struct {
	cfg cfgII18N
}

var _ I18N = &i18n{}

type cfgII18N map[string]ii18n.Config

func NewI18n() *i18n {
	return &i18n{}
}

func (thisI18n *i18n) initialize() *i18n {
	var c INI = NewIni().LoadByFN(ConfCommons)

	arrFileMap := make(map[string]string)
	strFileMap := c.K(
		MapConfLists[ConfCommons][0],
		MapConfParam[MapConfLists[ConfCommons][0]][2],
	).String()
	if strFileMap == "" {panic("Error Translate File")}

	for _, strFM := range strings.Split(strFileMap, ",") {
		arrFM := strings.Split(strFM, ":")
		arrFileMap[arrFM[0]] = arrFM[1]
	}

	thisI18n.cfg = cfgII18N{
		"app": {
			SourceNewFunc: ii18n.NewJSONSource,
			OriginalLang:  c.K(
							MapConfLists[ConfCommons][0],
							MapConfParam[MapConfLists[ConfCommons][0]][0],
						   ).String(),
			BasePath:      c.K(
							MapConfLists[ConfCommons][0],
							MapConfParam[MapConfLists[ConfCommons][0]][1],
						   ).String(),
			FileMap:       arrFileMap,
		},
	}
	ii18n.NewI18N(thisI18n.cfg)
	return thisI18n
}

func (thisI18n *i18n) T(cg string, key string, ln string) string {
	thisI18n.initialize()
	return ii18n.T(cg, key, nil, ln)
}
