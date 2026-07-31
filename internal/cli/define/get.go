package define

import "github.com/lxsh-S/gos/internal/tools/gosget"

func GoGetRun(projectName, templateName string) {
	gosget.HandelGetCommand(projectName, templateName)
}
