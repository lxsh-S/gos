package blueprint

import "path/filepath"

func Goblueprint(ProjectName string) []string {
	projectName := ProjectName
	folders := []string{
		filepath.Join(projectName, "cmd"),
		filepath.Join(projectName, "internal"),
		filepath.Join(projectName, "pkg"),
	}
	return folders
}
