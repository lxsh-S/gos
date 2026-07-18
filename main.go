package main

import (
	"fmt"
	"os"

	"github.com/fatih/color" // Not the std one

	"github.com/lxsh-S/gos/internal/cli/define"

	"github.com/spf13/cobra"
)

var (
	projectLang string
	projectType string
	list        bool
)

func printSupported() {
	fmt.Println(color.HiBlueString("go") + ":  std, api, cli")
	fmt.Println(color.HiBlueString("ts") + ":  std, api, lib, nxtjs")
	fmt.Println(color.HiBlueString("cpp") + ": std, app, lib")
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "gos [projectName]",
		Short: "Gos is used to create folder structures fast!",
		Args: func(cmd *cobra.Command, args []string) error {
			if list {
				return nil
			}
			return cobra.ExactArgs(1)(cmd, args)
		},
		Version: "0.7.2",
		RunE: func(cmd *cobra.Command, args []string) error {
			if list {
				printSupported()
				return nil
			}
			fmt.Printf("Building project: %s\nProject Language: %s\nProject Type: %s\n", color.CyanString(args[0]), color.HiBlueString(projectLang), color.YellowString(projectType))
			projectName := args[0]
			if err := define.Create(projectName, projectType, projectLang); err != nil {
				return err
			}
			fmt.Println(color.GreenString("Done!"))
			return nil
		},
	}

	rootCmd.Flags().StringVarP(&projectType, "type", "t", "std", "Project structure type [go:'std', 'api', 'cli']; [ts: 'api', 'nxtjs', 'lib', 'std']; [cpp: 'app', 'lib', 'std']")

	rootCmd.Flags().StringVarP(&projectLang, "lang", "l", "go", "Project Language ['go', 'ts', 'cpp']")

	rootCmd.Flags().BoolVar(&list, "list", false, "List all the project type combinations for each language")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(color.RedString("Error: %s", err))
		os.Exit(1)
	}
}
