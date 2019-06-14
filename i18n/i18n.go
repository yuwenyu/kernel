package i18n

import (
	"github.com/syyongx/ii18n"
)

func init() {}

type I18n struct {

}

func New() *I18n {
	cfg := map[string]ii18n.Config{
		"app": {
			SourceNewFunc: ii18n.NewJSONSource,
			OriginalLang:  "en",
			BasePath:      "resources/lang",
			FileMap: map[string]string{
				"app":   "app.json",
			},
		},
	}
	ii18n.NewI18N(cfg)
	return &I18n{}
}

func (i *I18n) T(k string, lang string) string {
	return ii18n.T("app", k, nil, lang)
}
