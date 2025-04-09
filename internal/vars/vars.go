package vars

import "strings"

const (
	currentTime = "CURTIME"
	currentName = "CURNAME"
)

type Vars map[string]string

func ReplacePatterns(text string, allFields map[string]string) string {
	for variableName, value := range allFields {
		placeholder := "{{" + variableName + "}}"
		text = strings.ReplaceAll(text, placeholder, value)
	}
	return text
}

func (v Vars) GetCurrentTime() string {
	return v[currentTime]
}

func (v Vars) GetCurrentName() string {
	return v[currentName]
}

func (v Vars) SetCurrentTime(time string) {
	v[currentTime] = time
}

func (v Vars) SetCurrentName(name string) {
	v[currentName] = name
}
