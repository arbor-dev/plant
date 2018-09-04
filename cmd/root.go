/**
* Copyright © 2018, ACM@UIUC
*
* This file is part of the Arbor Project.
*
* The Arbor Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/arbor-dev/plant/plant"
	"github.com/spf13/cobra"
	//"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "plant [arbor project name]",
	Short: "A command line tool to create new arbor projects",
	Long:  `plant is a command line tool that can create and manage arbor projects`,
	Run:   plantRun,
}

var port int

func plantRun(cmd *cobra.Command, args []string) {
	if len(args) > 1 || len(args) == 0 {
		os.Stderr.WriteString("Error: plant command expects one name for the project\n")
		os.Exit(-1)
	}

	fmt.Println("Planting ...")

	project := args[0]
	dir, _ := os.Getwd()
	root := filepath.Base(dir)

	plant.CreateMainFile(port, project, root)
	plant.CreateServicesFiles(project)
	plant.CreateConfigFile(project)
	fmt.Println("Finshied planting: " + project)
	fmt.Println("Correct import statements in generated go files!")
}

func init() {

	// --port or -p for api-gateway port
	rootCmd.Flags().IntVarP(&port, "port", "p", 8000, "port for api-gateway")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
