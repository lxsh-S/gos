package mkdir

import (
	"log"
	"os"
	"path/filepath"
)

func GOMkdir(folderName string) {
	targetPath := filepath.Join(".", folderName)
	err := os.Mkdir(targetPath, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
