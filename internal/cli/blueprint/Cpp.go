package blueprint

import (
	"fmt"
	"path/filepath"

	"github.com/lxsh-S/gos/templates"
)

func CppBlueprint(projectName, projectType string) (*Blueprint, error) {
	bp := &Blueprint{}
	switch projectType {

	case "app":
		// Game
		bp.Folders = []string{
			projectName,
			filepath.Join(projectName, "src"),
			filepath.Join(projectName, "include"),
			filepath.Join(projectName, "tests"),
			filepath.Join(projectName, "assets"),
			filepath.Join(projectName, "cmake"),
		}
		dataMain, err := templates.FS.ReadFile("cpp/app/main.cpp.tmpl")
		if err != nil {
			return nil, err
		}

		dataCMakeLists, err := templates.FS.ReadFile("cpp/app/CMakeLists.txt.tmpl")
		if err != nil {
			return nil, err
		}

		dataGitignore, err := templates.FS.ReadFile("cpp/app/gitignore.tmpl")
		if err != nil {
			return nil, err
		}

		dataCompilerWarnings, err := templates.FS.ReadFile("cpp/app/CompilerWarnings.cmake.tmpl")
		if err != nil {
			return nil, err
		}

		dataReadme, err := templates.FS.ReadFile("cpp/app/README.md.tmpl")
		if err != nil {
			return nil, err
		}

		bp.Files = []File{
			{
				Path:    filepath.Join(projectName, "src", "main.cpp"),
				Content: string(dataMain),
			},

			{
				Path:    filepath.Join(projectName, "CMakeLists.txt"),
				Content: string(dataCMakeLists),
			},

			{
				Path:    filepath.Join(projectName, ".gitignore"),
				Content: string(dataGitignore),
			},

			{
				Path:    filepath.Join(projectName, "cmake", "CompilerWarnings.cmake"),
				Content: string(dataCompilerWarnings),
			},

			{
				Path:    filepath.Join(projectName, "README.md"),
				Content: string(dataReadme),
			},
		}

	case "lib":
		// for building a reusabel library
		bp.Folders = []string{
			projectName,
			filepath.Join(projectName, "src"),
			filepath.Join(projectName, "include", projectName),
			filepath.Join(projectName, "tests"),
			filepath.Join(projectName, "examples"),
		}

		dataLibCpp, err := templates.FS.ReadFile("cpp/lib/library.cpp.tmpl")
		if err != nil {
			return nil, err
		}

		dataLibHpp, err := templates.FS.ReadFile("cpp/lib/library.hpp.tmpl")
		if err != nil {
			return nil, err
		}

		dataExampleCpp, err := templates.FS.ReadFile("cpp/lib/example.cpp.tmpl")
		if err != nil {
			return nil, err
		}

		dataTestCpp, err := templates.FS.ReadFile("cpp/lib/test.cpp.tmpl")
		if err != nil {
			return nil, err
		}

		dataCMakeLists, err := templates.FS.ReadFile("cpp/lib/CMakeLists.txt.tmpl")
		if err != nil {
			return nil, err
		}

		dataGitignore, err := templates.FS.ReadFile("cpp/lib/gitignore.tmpl")
		if err != nil {
			return nil, err
		}

		dataReadme, err := templates.FS.ReadFile("cpp/lib/README.md.tmpl")
		if err != nil {
			return nil, err
		}

		bp.Files = []File{
			{
				Path:    filepath.Join(projectName, "src", "library.cpp"),
				Content: string(dataLibCpp),
			},

			{
				Path:    filepath.Join(projectName, "include", projectName, "library.hpp"),
				Content: string(dataLibHpp),
			},

			{
				Path:    filepath.Join(projectName, "examples", "example.cpp"),
				Content: string(dataExampleCpp),
			},

			{
				Path:    filepath.Join(projectName, "tests", "test.cpp"),
				Content: string(dataTestCpp),
			},

			{
				Path:    filepath.Join(projectName, "CMakeLists.txt"),
				Content: string(dataCMakeLists),
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
			projectName,
			filepath.Join(projectName, "src"),
			filepath.Join(projectName, "tests"),
			filepath.Join(projectName, "build"),
		}
		dataMain, err := templates.FS.ReadFile("cpp/std/main.cpp.tmpl")
		if err != nil {
			return nil, err
		}

		dataCMakeLists, err := templates.FS.ReadFile("cpp/std/CMakeLists.txt.tmpl")
		if err != nil {
			return nil, err
		}

		dataGitignore, err := templates.FS.ReadFile("cpp/std/gitignore.tmpl")
		if err != nil {
			return nil, err
		}

		dataReadme, err := templates.FS.ReadFile("cpp/std/README.md.tmpl")
		if err != nil {
			return nil, err
		}

		bp.Files = []File{
			{
				Path:    filepath.Join(projectName, "src", "main.cpp"),
				Content: string(dataMain),
			},

			{
				Path:    filepath.Join(projectName, "CMakeLists.txt"),
				Content: string(dataCMakeLists),
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
		return nil, fmt.Errorf("unknown cpp project type: %q", projectType)
	}
	return bp, nil
}
