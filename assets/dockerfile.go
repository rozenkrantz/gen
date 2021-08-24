package assets

import "fmt"

func GetDockerfile(projectName string) string {
	return fmt.Sprintf(`FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o /%[1]s

CMD [ "/%[1]s" ]`, projectName)
}
