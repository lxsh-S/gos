package blueprint

import "path/filepath"

func TsBlueprint(projectName string, projectType string) []string {
	var folders []string

	switch projectType {
	case "nxtjs":
		// nextjs
		folders = []string{
			filepath.Join(projectName, "public"),
			filepath.Join(projectName, "src", "app"),
			filepath.Join(projectName, "src", "components"),
			filepath.Join(projectName, "src", "features"),
			filepath.Join(projectName, "src", "lib"),
			filepath.Join(projectName, "src", "hooks"),
			filepath.Join(projectName, "src", "types"),
		}

	case "api":
		// api
		folders = []string{
			filepath.Join(projectName, "src", "config"),
			filepath.Join(projectName, "src", "controllers"),
			filepath.Join(projectName, "src", "middlewares"),
			filepath.Join(projectName, "src", "models"),
			filepath.Join(projectName, "src", "routes"),
			filepath.Join(projectName, "src", "services"),
			filepath.Join(projectName, "src", "types"),
			filepath.Join(projectName, "tests"),
		}

	case "lib":
		// library
		folders = []string{
			filepath.Join(projectName, "src"),
			filepath.Join(projectName, "src", "utils"),
			filepath.Join(projectName, "src", "types"),
			filepath.Join(projectName, "src", "__tests__"),
		}

	default:
		// std
		folders = []string{
			filepath.Join(projectName, "src"),
			filepath.Join(projectName, "dist"),
		}
	}

	return folders
}
