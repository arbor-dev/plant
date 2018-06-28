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
	"os"
	"fmt"
	"path/filepath"
	
	"github.com/arbor-dev/seedling/plant"
	"github.com/spf13/cobra"
)

var plantCmd = &cobra.Command{
	Use:   "plant [arbor project name]",
    Short: "Creates an arbor project.",
    Long:  `Creates an arbor project.
    `,
    Run: plantRun,
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
	plantCmd.Flags().IntVarP(&port, "port", "p", 8000, "port for api-gateway")


	rootCmd.AddCommand(plantCmd)
}

