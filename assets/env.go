package assets

import "fmt"

func GetENV(projectName string) string {
	return fmt.Sprintf(`DB_DRIVER="mysql"
DB_USERNAME="doro"
DB_PASSWORD="12345"
DB_PROTOCOL="tcp"
DB_ADDRESS="localhost"
DB_PORT="3306"
DB_NAME="%[1]s"
DB_PARAMS="parseTime=true"

SERVER_PORT=":7070"

BASE_API="/"
LOG_PATH="./log"
`, projectName)
}
