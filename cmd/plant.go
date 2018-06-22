/**
* Copyright © 2018, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package cmd

import (
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
	plant.CreateMainFile(port, "test", "test-root")
}

func init() {

	// --port or -p for api-gateway port
	plantCmd.Flags().IntVarP(&port, "port", "p", 8000, "port for api-gateway")


	rootCmd.AddCommand(plantCmd)
}

