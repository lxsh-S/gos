package define

import (
	"fmt"
	"os/exec"

	"github.com/lxsh-S/gos/internal/tools/mkdir"
)

func GOMkdirRun(folderName string) {
	mkdir.GOMkdir(folderName)
	cmd := exec.Command("git", "init", folderName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error initializing git repository in %s: %v\n%s\n", folderName, err, string(output))
		return
	}

	fmt.Print(string(output))
}
