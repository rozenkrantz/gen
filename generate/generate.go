package generate

import (
	"github.com/dequinox/gen/assets"
	"os/exec"

	"os"
)

type Config struct {

	// MainFile the main file
	MainFile string

	// ProjectName name of the project/base directory
	ProjectName string
}

func Build(config *Config) error {

	for _, dir := range getDirs() {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	err := writeFile(assets.GetMainFile(config.ProjectName), "./"+config.MainFile)
	if err != nil {
		return err
	}

	err = writeFile(assets.GetDockerfile(config.ProjectName), "./Dockerfile")
	if err != nil {
		return err
	}

	err = writeFile(assets.GetENV(config.ProjectName), "./.env")
	if err != nil {
		return err
	}

	err = writeFile(assets.GetConfigFile(), "./source/config/config.go")
	if err != nil {
		return err
	}

	err = writeFile(assets.GetMysqlDB(config.ProjectName), "./source/db/db.go")
	if err != nil {
		return err
	}

	err = writeFile(assets.GetLogger(), "./source/logger/logger.go")
	if err != nil {
		return err
	}

	err = writeFile(assets.GetMultiplexer(config.ProjectName), "./source/multiplexer/multiplexer.go")
	if err != nil {
		return err
	}

	err = writeFile(assets.GetHandlerModule(), "./source/handlers/module.go")
	if err != nil {
		return err
	}

	err = writeFile(assets.GetServiceModule(), "./source/services/module.go")
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "init", config.ProjectName)
	err = cmd.Run()

	if err != nil {
		return err
	}

	return nil
}

func writeFile(b string, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(b))
	return err
}

func getDirs() []string {
	return []string{
		"./source", "./source/config", "./source/db",
		"./source/config", "./source/entity", "./source/handlers",
		"./source/logger", "./source/multiplexer", "./source/services",
	}
}
