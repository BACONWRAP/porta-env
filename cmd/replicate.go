/*
Copyright © 2024 Chris Hooke chooke@unca.edu
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// replicateCmd represents the replicate command
var replicateCmd = &cobra.Command{
	Use:   "replicate",
	Short: "Replicate an environment on this machine.",
	Long: `From a remote repository, replicate the defined environment in home-manager.nix on this machine.
Installs nix and home-manager if not already installed. Requires sudo.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("replicate called")
	},
}

func init() {
	rootCmd.AddCommand(replicateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// replicateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// replicateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
