package assets

func GetServiceModule(projectName string) (string, string) {
	return `package services

import (
	"go.uber.org/fx"
)

var Module = fx.Options()
`, "./" + projectName + "/source/services/module.go"
}
