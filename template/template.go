package template

import (
	"io/fs"
	"strings"
	"text/template"
)

func TemplatesFS(fsys fs.FS, sub string, patterns ...string) Templates {
	if sub, err := fs.Sub(fsys, sub); err != nil {
		panic(err)
	} else if tmpls, err := template.ParseFS(sub, patterns...); err != nil {
		panic(err)
	} else {
		return Templates{
			tmpls: tmpls,
		}
	}
}

type Templates struct {
	tmpls *template.Template
}

func (t *Templates) Render(name string, data any) string {
	var result strings.Builder
	if err := t.tmpls.ExecuteTemplate(&result, name, data); err != nil {
		panic(err)
	} else {
		return result.String()
	}
}
