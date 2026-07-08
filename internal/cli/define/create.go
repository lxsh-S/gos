package define

import "github.com/lxsh-S/gos/internal/cli/create"

func Create(projectName, projectType, projectLang string) {
	if projectLang == "go" {
		create.Creatego(projectName, projectType)
	} else if projectLang == "ts" {
		create.CreateTs(projectName)
	}
}
