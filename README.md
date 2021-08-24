# gen
Generate project skeleton

# Getting Started

1. Download gen by using
```shell
$ go get -u github.com/dequinox/gen
```
2. Run `gen init -p [PROJECT-NAME]` to generate project structure.

```shell
$ gen init -p [PROJECT-NAME]
```

# gen cli 
```shell
NAME:
   gen init - Create skeleton

USAGE:
   gen init [command options] [arguments...]

OPTIONS:
   --mainInfo value, -m value     Main go file (default: "main.go")
   --projectInfo value, -p value  Name of the project (default: "my-project")
   --help, -h                     show help (default: false)
```