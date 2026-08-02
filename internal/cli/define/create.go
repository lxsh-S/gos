package define

import (
	"fmt"
	"os/exec"

	"github.com/lxsh-S/gos/internal/cli/create"
)

func Create(projectName, projectType, projectLang string) error {
	// Create the project files first
	var err error
	switch projectLang {
	case "go":
		err = create.CreateGO(projectName, projectType)
	case "ts":
		err = create.CreateTS(projectName, projectType)
	case "cpp":
		err = create.CreatCPP(projectName, projectType)
	default:
		return fmt.Errorf("unknown project language: %q (expected: go, ts, cpp)", projectLang)
	}

	if err != nil {
		return fmt.Errorf("failed to create project: %w", err)
	}

	// Initialize git inside the project directory
	cmd := exec.Command("git", "init", projectName)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to initialize git repository in %s: %w\n%s", projectName, err, string(output))
	}

	return nil
}
