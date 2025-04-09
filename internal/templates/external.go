package templates

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/catmorte/go-remarks/internal/vars"
)

type externalTemplate struct {
	Name        string
	RunTemplate string
	Vars        string
}

func (d externalTemplate) GetName() string {
	return d.Name
}

type TemplateData struct {
	Vars map[string]string
}

func (d externalTemplate) Compile(vrs vars.Vars) error {
	now, ok := vrs[string(InternalSimpleTimeField)]
	if !ok {
		now = time.Now().Format("2006-01-02 15:04:05")
		vrs[string(InternalSimpleTimeField)] = now
	}
	t, err := template.New(d.Name).Funcs(templateFuncs).Parse(d.RunTemplate)
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

func (d externalTemplate) GetVars() []string {
	return strings.Split(d.Vars, "\n")
}
