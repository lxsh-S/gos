package blueprint

import (
	"fmt"
	"path/filepath"

	"github.com/lxsh-S/gos/templates"
)

func TSBlueprint(projectName string, projectType string) (*Blueprint, error) {
	bp := &Blueprint{}
	switch projectType {
	case "nxtjs":
		// nextjs
		bp.Folders = []string{
			filepath.Join(projectName, "public"),
			filepath.Join(projectName, "src", "app"),
			filepath.Join(projectName, "src", "components"),
			filepath.Join(projectName, "src", "features"),
			filepath.Join(projectName, "src", "lib"),
			filepath.Join(projectName, "src", "hooks"),
			filepath.Join(projectName, "src", "types"),
		}
		dataReadme, err := templates.FS.ReadFile("ts/nxtjs/README.md.tmpl")
		if err != nil {
			return nil, err
		}

		dataGitignore, err := templates.FS.ReadFile("ts/common/gitignore.tmpl")
		if err != nil {
			return nil, err
		}

		bp.Files = []File{
			{
				Path:    filepath.Join(projectName, "README.md"),
				Content: string(dataReadme),
			},

			{
				Path:    filepath.Join(projectName, ".gitignore"),
				Content: string(dataGitignore),
			},
		}

	case "api":
		// api
		bp.Folders = []string{
			filepath.Join(projectName, "src", "config"),
			filepath.Join(projectName, "src", "controllers"),
			filepath.Join(projectName, "src", "middlewares"),
			filepath.Join(projectName, "src", "models"),
			filepath.Join(projectName, "src", "routes"),
			filepath.Join(projectName, "src", "services"),
			filepath.Join(projectName, "src", "types"),
			filepath.Join(projectName, "tests"),
		}

		dataIndex, err := templates.FS.ReadFile("ts/api/index.ts.tmpl")
		if err != nil {
			return nil, err
		}

		dataPackage, err := templates.FS.ReadFile("ts/api/package.json.tmpl")
		if err != nil {
			return nil, err
		}

		dataGitignore, err := templates.FS.ReadFile("ts/common/gitignore.tmpl")
		if err != nil {
			return nil, err
		}

		dataReadme, err := templates.FS.ReadFile("ts/common/README.md.tmpl")
		if err != nil {
			return nil, err
		}

		bp.Files = []File{
			{
				Path:    filepath.Join(projectName, "src", "index.ts"),
				Content: string(dataIndex),
			},

			{
				Path:    filepath.Join(projectName, "package.json"),
				Content: string(dataPackage),
			},

			{
				Path:    filepath.Join(projectName, ".gitignore"),
				Content: string(dataGitignore),
			},

			{
				Path:    filepath.Join(projectName, "README.md"),
				Content: string(dataReadme),
			},
		}

	case "lib":
		// library
		bp.Folders = []string{
			filepath.Join(projectName, "src"),
			filepath.Join(projectName, "src", "utils"),
			filepath.Join(projectName, "src", "types"),
			filepath.Join(projectName, "src", "__tests__"),
		}

		dataIndex, err := templates.FS.ReadFile("ts/lib/index.ts.tmpl")
		if err != nil {
			return nil, err
		}

		dataPackage, err := templates.FS.ReadFile("ts/lib/package.json.tmpl")
		if err != nil {
			return nil, err
		}

		dataGitignore, err := templates.FS.ReadFile("ts/common/gitignore.tmpl")
		if err != nil {
			return nil, err
		}

		dataReadme, err := templates.FS.ReadFile("ts/common/README.md.tmpl")
		if err != nil {
			return nil, err
		}

		bp.Files = []File{
			{
				Path:    filepath.Join(projectName, "src", "index.ts"),
				Content: string(dataIndex),
			},

			{
				Path:    filepath.Join(projectName, "package.json"),
				Content: string(dataPackage),
			},

			{
				Path:    filepath.Join(projectName, ".gitignore"),
				Content: string(dataGitignore),
			},

			{
				Path:    filepath.Join(projectName, "README.md"),
				Content: string(dataReadme),
			},
		}

	case "std", "":
		// std
		bp.Folders = []string{
			filepath.Join(projectName, "src"),
			filepath.Join(projectName, "dist"),
		}

		dataIndex, err := templates.FS.ReadFile("ts/std/index.ts.tmpl")
		if err != nil {
			return nil, err
		}

		dataPackage, err := templates.FS.ReadFile("ts/std/package.json.tmpl")
		if err != nil {
			return nil, err
		}

		dataGitignore, err := templates.FS.ReadFile("ts/common/gitignore.tmpl")
		if err != nil {
			return nil, err
		}

		dataTSConfig, err := templates.FS.ReadFile("ts/std/tsconfig.json.tmpl")
		if err != nil {
			return nil, err
		}

		dataReadme, err := templates.FS.ReadFile("ts/common/README.md.tmpl")
		if err != nil {
			return nil, err
		}

		bp.Files = []File{
			{
				Path:    filepath.Join(projectName, "src", "index.ts"),
				Content: string(dataIndex),
			},

			{
				Path:    filepath.Join(projectName, "package.json"),
				Content: string(dataPackage),
			},

			{
				Path:    filepath.Join(projectName, "tsconfig.json"),
				Content: string(dataTSConfig),
			},

			{
				Path:    filepath.Join(projectName, ".gitignore"),
				Content: string(dataGitignore),
			},

			{
				Path:    filepath.Join(projectName, "README.md"),
				Content: string(dataReadme),
			},
		}

	default:
		return nil, fmt.Errorf("unknown ts project type: %q", projectType)
	}

	return bp, nil
}
