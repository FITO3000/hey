/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"f3s.tech/hey-cli/elasticstack"
	"f3s.tech/hey-utils/fileutil"
	"github.com/spf13/cobra"
)

// elasticstackCmd represents the elasticstack command
var elasticstackCmd = &cobra.Command{
	Use:   "elasticstack",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		installer := &elasticstack.Installer{}

		if len(args) == 0 {
			installer.InstallFolder = "."
		} else if len(args) == 1 {
			installer.InstallFolder = args[0]
		} else {
			cmd.Help()
			os.Exit(0)
		}

		if fileutil.Exists(installer.InstallFolder) {
			if err := installer.InstallAll(); err != nil {
				panic(err)
			}
		} else {
			panic(fmt.Errorf("installation-folder: %s does not exist", installer.InstallFolder))
		}
	},
}

func init() {
	installCmd.AddCommand(elasticstackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// elasticstackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// elasticstackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
