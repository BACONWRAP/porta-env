/*
Copyright © 2024 Chris Hooke chooke@unca.edu
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup your environment for replication elsewhere.",
	Long: `Setup your environment for replecation elsewhere. 
This will install nix and home-manager on your system. Requires sudo.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setup called")
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
