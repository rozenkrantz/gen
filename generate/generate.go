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
		if err := os.MkdirAll("./"+config.ProjectName+dir, os.ModePerm); err != nil {
			return err
		}
	}

	err := writeFile(assets.GetMainFile(config.ProjectName, config.MainFile))
	if err != nil {
		return err
	}

	for _, dataFunc := range getDataFunc() {
		data, path := dataFunc(config.ProjectName)
		err = writeFile(data, path)
		if err != nil {
			return err
		}
	}

	cmd := exec.Command("go", "mod", "init", config.ProjectName)
	cmd.Dir = "./" + config.ProjectName
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
		"/source", "/source/config", "/source/db",
		"/source/config", "/source/entity", "/source/handlers",
		"/source/logger", "/source/multiplexer", "/source/services",
	}
}

func getDataFunc() []func(string) (string, string) {
	return []func(string) (string, string){
		assets.GetENV,
		assets.GetServiceModule,
		assets.GetHandlerModule,
		assets.GetDockerfile,
		assets.GetLogger,
		assets.GetMultiplexer,
		assets.GetMysqlDB,
		assets.GetConfigFile,
		assets.GetGitIgnore,
	}
}
