// Creatego
package create

import (
	"fmt"
	"os"

	"github.com/lxsh-S/gos/internal/cli/blueprint"
)

func Creatego(ProjectName string) {
	projectName := ProjectName

	for _, folder := range blueprint.Goblueprint(projectName) {
		err := os.MkdirAll(folder, 0o755)
		if err != nil {
			fmt.Printf("Errror creating folder structure %s: %v", &folder, err)
		}
	}
}
