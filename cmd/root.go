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

	"github.com/spf13/cobra"
	//"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "seedling",
	Short: "A command line tool for arbor projects",
	Long: `seedling is a command line tool that can create and manage arbor projects`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
