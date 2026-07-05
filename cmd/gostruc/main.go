package main

import (
	"fmt"
	"os"

	"github.com/lxsh-S/gostruc/internal/cli/create"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gos",
		Short: "Gos is used to create folder structures fast!",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome...")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Les check if the name is given
	if len(os.Args) < 2 {
		fmt.Println("Error: Please specify a project name")
		fmt.Println("Usage: 'gos <your-project-name>'")
		return
	} else { // if name is given then let's call Creatego func in Golang.go
		projectName := os.Args[1]
		create.Creatego(projectName)
	}
}
