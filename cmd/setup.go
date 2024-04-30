/*
Copyright © 2024 Chris Hooke chooke@unca.edu
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	//	"log"
	//	"os/exec"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup your environment for use with porta-env.",
	Long: `Setup your environment for use with porta-env. 
This will install nix and home-manager on your system. Requires sudo.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setup called")
		executeCmd("sh", "-c", "curl -L https://nixos.org/nix/install | sh -s -- --daemon")
		executeCmd("sh", "-c", ". '/nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh'")
		nixChannel, _ := filepath.Glob("/nix/store/*/bin/nix-channel")
		nixShell, _ := filepath.Glob("/nix/store/*/bin/nix-shell")

		bins, _ := filepath.Glob("/nix/store/*/bin")
		for _, bin := range bins {
			fmt.Println("setting: " + bin)
			os.Setenv("PATH", os.ExpandEnv(bin+":$PATH"))
		}
		executeCmd("sh", "-c", nixChannel[0]+" --add https://github.com/nix-community/home-manager/archive/master.tar.gz home-manager")
		executeCmd("sh", "-c", nixChannel[0]+" --update")
		executeCmd("sh", "-c", nixShell[0]+" '<home-manager>' -A install")
		fmt.Printf(`
Please restart your shell to complete the setup.
You can run porta-env create to build a new toolbox,
or porta-env replicate to download an existing one.\n`)
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
