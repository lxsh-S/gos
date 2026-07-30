package gosget

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func HandelGetCommand(projectName, templateName string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w \n", err)
	}

	templateDir := filepath.Join(homeDir, ".gos", "usr_Templates", templateName)

	info, err := os.Stat(templateDir)
	if err != nil || !info.IsDir() {
		return fmt.Errorf("error: This directory '%s' already exists in the current foler ", projectName)
	}

	fmt.Printf("Creating project '%s' from template '%s'...\n", projectName, templateName)

	targetDir := filepath.Join(".", projectName)

	err = filepath.WalkDir(templateDir, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		relPath, err := filepath.Rel(templateDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(targetDir, relPath)

		if d.IsDir() {
			return os.MkdirAll(destPath, 0o755)
		}

		return copyFile(path, destPath)
	})

	err = filepath.WalkDir(templateDir, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		relPath, err := filepath.Rel(templateDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(targetDir, relPath)

		if d.IsDir() {
			return os.MkdirAll(destPath, 0o755)
		}
		return copyFile(path, destPath)
	})
	if err != nil {
		return fmt.Errorf("failed to generate project: %w", err)
	}

	fmt.Printf("Successfully created '%s'! Happy coding till it kills you =)\n", projectName)
	return nil
}

func copyFile(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}

	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}

	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
