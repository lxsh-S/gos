package blueprint

import (
	"fmt"
	"path/filepath"

	"github.com/lxsh-S/gos/templates"
)

type File struct {
	Path    string
	Content string
}

type Blueprint struct {
	Files   []File
	Folders []string
}

func Goblueprint(projectName string, projectType string) (*Blueprint, error) {
	bp := &Blueprint{}

	switch projectType {
	case "api":
		bp.Folders = []string{
			projectName,
			filepath.Join(projectName, "cmd", "api"),
			filepath.Join(projectName, "internal", "handlers"),
			filepath.Join(projectName, "internal", "models"),
			filepath.Join(projectName, "pkg"),
		}

		dataMain, err := templates.FS.ReadFile("go/main.go.tmpl")
		if err != nil {
			return nil, err
		}

		dataMod, err := templates.FS.ReadFile("go/go.mod.tmpl")
		if err != nil {
			return nil, err
		}

		bp.Files = []File{
			{
				Path:    filepath.Join(projectName, "cmd", "api", "main.go"),
				Content: string(dataMain),
			},

			{
				Path:    filepath.Join(projectName, "go.mod"),
				Content: string(dataMod),
			},
		}

	case "cli":
		bp.Folders = []string{
			projectName,
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "internal", "config"),
			filepath.Join(projectName, "pkg", "utils"),
		}
		dataMain, err := templates.FS.ReadFile("go/main.go.tmpl")
		if err != nil {
			return nil, err
		}

		dataMod, err := templates.FS.ReadFile("go/go.mod.tmpl")
		if err != nil {
			return nil, err
		}
		bp.Files = []File{
			{
				Path:    filepath.Join(projectName, "cmd", "main.go"),
				Content: string(dataMain),
			},

			{
				Path:    filepath.Join(projectName, "go.mod"),
				Content: string(dataMod),
			},

			{
				Path: filepath.Join(projectName, "go.mod"),
			},
		}

	case "std", "":
		// Our OG (std)
		bp.Folders = []string{
			projectName,
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "internal"),
			filepath.Join(projectName, "pkg"),
		}

		dataMain, err := templates.FS.ReadFile("go/main.go.tmpl")
		if err != nil {
			return nil, err
		}

		dataMod, err := templates.FS.ReadFile("go/go.mod.tmpl")
		if err != nil {
			return nil, err
		}
		bp.Files = []File{
			{
				Path:    filepath.Join(projectName, "cmd", "main.go"),
				Content: string(dataMain),
			},

			{
				Path:    filepath.Join(projectName, "go.mod"),
				Content: string(dataMod),
			},
		}

	default:
		return nil, fmt.Errorf("unknown go project type: %q", projectType)
	}
	return bp, nil
}
