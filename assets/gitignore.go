package assets

func GetGitIgnore(projectName string) (string, string) {
	return `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

*.log
.idea
*.idea
*.iml
*.xml
`, "./" + projectName + "/.gitignore"
}
