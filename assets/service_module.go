package assets

func GetServiceModule() string {
	return `package services

import (
	"go.uber.org/fx"
)

var Module = fx.Options()
`
}
