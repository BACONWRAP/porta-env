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
		executeCmd("sh", "-c", "curl -L https://nixos.org/nix/install | sh -s -- --daemon")
		executeCmd("sh", "-c", ". '/nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh'")
		executeCmd("sh", "-c", "PATH='$HOME/.nix-profile/bin:/nix/var/nix/profiles/default/bin:$PATH'")
		executeCmd("sh", "-c", "nix-channel --add https://github.com/nix-community/home-manager/archive/master.tar.gz home-manager")
		executeCmd("sh", "-c", "nix-channel --update")
		executeCmd("sh", "-c", "nix-shell '<home-manager>' -A install")
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
