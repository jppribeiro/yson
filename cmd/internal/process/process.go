package process

import (
	"yson.com/yson/cmd/internal/input"
)

func Translate(file input.FileData) string {
	switch file.Type {
	case input.InputJSON:
		return ReadJSON(file)
	case input.InputYaml:
		return ReadYAML(file)
	default:
		return ""
	}
}
