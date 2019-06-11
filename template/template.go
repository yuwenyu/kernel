package template

import (
	"fmt"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

func init() {}

type Template struct {
	dir string
}

func New() *Template {
	return &Template {}
}

func (t *Template) SetDir(dir string) *Template {
	t.dir = dir
	return t
}

func (t *Template) Tpl() multitemplate.Renderer {
	if t.dir == "" {
		panic("Error: Empty Template Dir")
	}

	tpl := multitemplate.NewRenderer()

	layout, err := filepath.Glob(t.dir + "/" + "layouts/wyu.html")
	if err != nil {
		panic(fmt.Sprintf("Template Layout-wyu Error: %s", err.Error()))
	}

	shareds, err := filepath.Glob(t.dir + "/" + "shared/*.html")
	if err != nil {
		panic(fmt.Sprintf("Template Shared-wyu Error: %s", err.Error()))
	}

	arrTPL := make([]string, 1)
	arrTPL  = append(layout, t.dir + "/views/index.html")

	for _, shared := range shareds {
		arrTPL = append(arrTPL, shared)
	}

	tpl.AddFromFiles("index.html", arrTPL ...)

	return tpl
}

//	//views, err := filepath.Glob(dir + "/" + "views/*.html")
//	//if err != nil {
//	//	panic(fmt.Sprintf("Template view Error: %s", err.Error()))
//	//}
//	//
//	//for _, view := range views {
//	//	layoutCopy := make([]string, len(layouts))
//	//	copy(layoutCopy, layouts)
//	//	log.Println(layoutCopy)
//	//	fs := append(layoutCopy, view)
//	//	tpl.AddFromFiles(filepath.Base(view), fs ...)
//	//	log.Println(fs)
//	//}
