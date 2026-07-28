package gosadd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func HandelAddCommand(folderName string) error {
	info, err := os.Stat(folderName)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("error: folder '%s' does not exist", folderName)
		}
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("This: %s is a file not a folder!", folderName)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Failed to reach the user directory: %w", err)
	}

	templateName := filepath.Base(folderName)
	destDir := filepath.Join(homeDir, ".gos", "usr_Templates", templateName)

	if _, err := os.Stat(destDir); err == nil {
		fmt.Printf("Warning: This template %s already exists, Overwriting...\n", templateName)
		if err := os.RemoveAll(destDir); err != nil {
			return fmt.Errorf("failed to overwrite existing template: %w", err)
		}
	}

	fmt.Printf("Saving %s as template %s....\n", folderName, templateName)

	err = filepath.WalkDir(folderName, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		relPath, err := filepath.Rel(folderName, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(destDir, relPath)

		if d.IsDir() {
			return os.MkdirAll(targetPath, 0o755)
		}
		return copyFile(path, targetPath)
	})
	if err != nil {
		return fmt.Errorf("failed to copy template files: %w", err)
	}

	fmt.Printf("Successfully saved template!: %s\n", templateName)
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
