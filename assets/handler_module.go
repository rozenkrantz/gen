package assets

func GetHandlerModule() string {
	return `package handlers

import (
	"go.uber.org/fx"
)

var Module = fx.Options()
`
}
