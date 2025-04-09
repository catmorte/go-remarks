package templates

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/catmorte/go-remarks/internal/vars"
)

var ErrNotExist = errors.New("no template defined")

type (
	DefinedTemplate interface {
		GetName() string
		Compile(vars.Vars) error
		GetVars() []string
	}

	DefinedTemplates []DefinedTemplate
)

func GetDefinedTemplates(cfgPath string) (DefinedTemplates, error) {
	templates := InternalTemplates()

	info, err := os.Lstat(cfgPath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return templates, nil
		}
		return nil, err
	}

	if info.Mode()&os.ModeSymlink != 0 {
		cfgPath, err = filepath.EvalSymlinks(cfgPath)
		if err != nil {
			return nil, err
		}
	}

	err = filepath.Walk(cfgPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		folderName := filepath.Base(path)
		runTemplatePath := filepath.Join(cfgPath, folderName, "remark.tmpl")
		varsPath := filepath.Join(cfgPath, folderName, "vars")
		_, err = os.Stat(runTemplatePath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return nil
			}
			return err
		}
		_, err = os.Stat(varsPath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return nil
			}
			return err
		}
		runTemplate, err := os.ReadFile(runTemplatePath)
		if err != nil {
			return err
		}
		vars, err := os.ReadFile(varsPath)
		if err != nil {
			return err
		}

		templates = append(templates, externalTemplate{
			Name:        folderName,
			RunTemplate: string(runTemplate),
			Vars:        string(vars),
		})
		return nil
	})

	return templates, err
}

func (dts DefinedTemplates) FindByName(name string) (DefinedTemplate, error) {
	for _, v := range dts {
		if v.GetName() == name {
			return v, nil
		}
	}
	return nil, ErrNotExist
}

func InternalTemplates() []DefinedTemplate {
	return []DefinedTemplate{
		internalSimpleTemplate,
	}
}
