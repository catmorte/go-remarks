package templates

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"
	"time"

	"github.com/catmorte/go-remarks/internal/vars"
)

type internalSimple string

const (
	InternalSimpleNameField FieldVar = "name"
	InternalSimpleTagsField FieldVar = "tags"
	InternalSimpleTimeField FieldVar = "time"
)

//go:embed internal_simple_template.md
var internalSimpleTemplate internalSimple

func (d internalSimple) GetName() string {
	return "simple"
}

func (d internalSimple) Compile(vrs vars.Vars) error {
	now, ok := vrs[string(InternalSimpleTimeField)]
	if !ok {
		now = time.Now().Format("2006-01-02 15:04:05")
		vrs[string(InternalSimpleTimeField)] = now
	}
	t, err := template.New(d.GetName()).Funcs(templateFuncs).Parse(string(internalSimpleTemplate))
	if err != nil {
		return err
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, vrs)
	if err != nil {
		return err
	}
	fmt.Println(tpl.String())
	return nil
}

func (d internalSimple) GetVars() []string {
	return []string{
		string(InternalSimpleNameField),
		string(InternalSimpleTagsField),
		string(InternalSimpleTimeField),
	}
}
