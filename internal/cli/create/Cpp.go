package create

import (
	"fmt"
	"os"

	"github.com/lxsh-S/gos/internal/cli/blueprint"
)

func CreatCpp(projectName, projectType string) error {
	bp, err := blueprint.CppBlueprint(projectName, projectType)
	if err != nil {
		return err
	}
	for _, folder := range bp.Folders {
		err := os.MkdirAll(folder, 0o755)
		if err != nil {
			return fmt.Errorf("error creating folder structure %s: %w", folder, err)
		}
	}

	for _, file := range bp.Files {
		err := os.WriteFile(file.Path, []byte(file.Content), 0o644)
		if err != nil {
			return fmt.Errorf("Error creating files inside the directories %s: %w", file, err)
		}
	}
	return nil
}
