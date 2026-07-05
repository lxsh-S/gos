package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/lxsh-S/gos/internal/cli/create"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gos [projectName]",
		Short: "Gos is used to create folder structures fast!",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Building project: %s\n", color.CyanString(args[0]))
			fmt.Println(color.GreenString("Done!"))
			projectName := args[0]
			create.Creatego(projectName)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
