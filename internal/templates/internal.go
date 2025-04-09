package templates

import (
	"github.com/catmorte/go-remarks/internal/vars"
)

type FieldVar string

func (f FieldVar) Get(vrs vars.Vars) (string, bool) {
	v, ok := vrs[string(f)]
	return v, ok
}
