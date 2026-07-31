package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color" // Not the std one

	"github.com/lxsh-S/gos/internal/cli/define"

	"github.com/spf13/cobra"
)

var (
	projectLang          string
	projectType          string
	gomkdirFile          string
	goaddFile            string
	projectTemplate      string
	projectNameForGosget string
	list                 bool
)

func printSupported() {
	fmt.Println(color.HiBlueString("gosdir") + "'projectName'")
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

		Version: "0.9.5",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Les create a folder to store custom temolates by the user
			//
			//Get the user's dir:
			homedir, err := os.UserHomeDir()
			if err != nil {
				log.Fatal(err)
			}

			targetPath := filepath.Join(homedir, ".gos")
			templateFolder := filepath.Join(targetPath, "usr_Templates")

			err1 := os.MkdirAll(targetPath, 0o755)
			if err1 != nil {
				log.Fatal(err1)
			}

			err3 := os.MkdirAll(templateFolder, 0o755)
			if err3 != nil {
				log.Fatal(err3)
			}

			if list {
				printSupported()
				return nil
			}
			projectName := args[0]

			if projectName == "gosget" {
				templateName := projectTemplate
				projectName := projectNameForGosget
				define.GoGetRun(projectName, templateName)
			} else if projectName == "gosadd" {
				folderName := goaddFile
				define.GOAddRun(folderName)
			} else if projectName == "mkdir" {
				folderName := gomkdirFile
				define.GOMkdirRun(folderName)
				fmt.Printf("Empty dir: %s created!", folderName)
			} else {

				fmt.Printf("Building project: %s\nProject Language: %s\nProject Type: %s\n", color.CyanString(args[0]), color.HiBlueString(projectLang), color.YellowString(projectType))

				if err := define.Create(projectName, projectType, projectLang); err != nil {
					return err
				}
				fmt.Println(color.GreenString("Done!"))
			}
			return nil
		},
	}

	rootCmd.Flags().StringVarP(&projectType, "type", "t", "std", "Project structure type [go:'std', 'api', 'cli']; [ts: 'api', 'nxtjs', 'lib', 'std']; [cpp: 'app', 'lib', 'std']")

	rootCmd.Flags().StringVarP(&projectLang, "lang", "l", "go", "Project Language ['go', 'ts', 'cpp']")

	rootCmd.Flags().BoolVar(&list, "list", false, "List all the project type combinations for each language")

	rootCmd.Flags().StringVarP(&gomkdirFile, "gomkdir", "m", "gomkdir", "Makes a dir")

	rootCmd.Flags().StringVarP(&goaddFile, "add", "a", "gosadd", "Make your current project a template!")

	rootCmd.Flags().StringVarP(&projectTemplate, "template", "e", "", "Make a new folder with the stored template")

	rootCmd.Flags().StringVarP(&projectNameForGosget, "projectName", "p", "", "The project Name for the folder created using custom templates that are saved by you")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(color.RedString("Error: %s", err))
		os.Exit(1)
	}
}
