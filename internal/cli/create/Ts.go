package create

import (
	"fmt"
	"os"

	"github.com/lxsh-S/gos/internal/cli/blueprint"
)

func CreateTs(projectName string) {
	for _, folder := range blueprint.TsBlueprint(projectName) {
		err := os.MkdirAll(folder, 0o755)
		if err != nil {
			fmt.Printf("Errror creating folder structure %s: %v", folder, err)
		}
	}
}
