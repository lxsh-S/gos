package blueprint

import "path/filepath"

func TsBlueprint(ProjectName string) []string {
	projectName := ProjectName

	folders := []string{
		filepath.Join(projectName, "public"),
		filepath.Join(projectName, "src/app"),
		filepath.Join(projectName, "components"),
		filepath.Join(projectName, "features"),
		filepath.Join(projectName, "lib"),
		filepath.Join(projectName, "utils"),
	}
	return folders
}
