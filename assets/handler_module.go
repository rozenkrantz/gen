package assets

func GetHandlerModule(projectName string) (string, string) {
	return `package handlers

import (
	"go.uber.org/fx"
)

var Module = fx.Options()
`, "./" + projectName + "/source/handlers/module.go"
}
