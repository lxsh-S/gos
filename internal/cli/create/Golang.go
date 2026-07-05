// Creatego
package create

import (
	"fmt"
	"os"

	"github.com/lxsh-S/gostruc/internal/cli/blueprint"
)

func Creatego(ProjectName string) {
	projectName := ProjectName

	// folders := []string{
	// 	filepath.Join(projectName, "cmd"),
	// 	filepath.Join(projectName, "internal"),
	// 	filepath.Join(projectName, "pkg"),
	// }
	for _, folder := range blueprint.Goblueprint(projectName) {
		err := os.MkdirAll(folder, 0o755)
		if err != nil {
			fmt.Printf("Errror creating folder structure %s: %v", &folder, err)
		}
	}
}
