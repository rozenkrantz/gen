package assets

import "fmt"

func GetMainFile(projectName string) string {
	return fmt.Sprintf(`package main

import (
	"%[1]s/source/config"
	"%[1]s/source/db"
	"%[1]s/source/handlers"
	"%[1]s/source/logger"
	"%[1]s/source/multiplexer"
	"%[1]s/source/services"

	"go.uber.org/fx"
)

func main() {

	mainModules := fx.Options(
		config.Module,
		logger.Module,
		db.Module,
		multiplexer.Module,
		services.Module,
		handlers.Module,
	)

	fx.New(mainModules).Run()

}`, projectName)
}
