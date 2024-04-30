/*
Copyright © 2024 Chris Hooke chooke@unca.edu
*/
package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

// replicateCmd represents the replicate command
var replicateCmd = &cobra.Command{
	Use:   "replicate",
	Short: "Replicate an environment on this machine.",
	Long: `From a remote repository, replicate the defined environment in home.nix on this machine.
Requires sudo.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No repository specified.")
			return
		}
		executeCmd("sh", "-c", "curl -OL "+args[0])
		fmt.Println("Downloaded main.tar.gz")
		fmt.Println("Moving main.tar.gz to " + os.Getenv("HOME") + ".config/home-manager/main.tar.gz")
		os.Rename("main.tar.gz", os.Getenv("HOME")+"/.config/home-manager/main.tar.gz")
		fmt.Println("Extracting main.tar.gz to " + os.Getenv("HOME") + ".config/home-manager")
		executeCmd("sh", "-c", "tar -xzf "+os.Getenv("HOME")+"/.config/home-manager/main.tar.gz -C "+os.Getenv("HOME")+"/.config/home-manager")
		fmt.Println("Moving toolbox-main to " + os.Getenv("HOME") + ".config/home-manager")
		executeCmd("sh", "-c", "mv "+os.Getenv("HOME")+"/.config/home-manager/toolbox-main/* "+os.Getenv("HOME")+"/.config/home-manager")
		fmt.Println("Modifying home.nix to use right env vars")
		replaceInFile(os.Getenv("HOME")+"/.config/home-manager/home.nix", *regexp.MustCompile(`home.username.*`),
			"home.username = \""+os.Getenv("USER")+"\";")
		replaceInFile(os.Getenv("HOME")+"/.config/home-manager/home.nix", *regexp.MustCompile(`home.homeDirectory.*`),
			"home.homeDirectory = \""+os.Getenv("HOME")+"\";")
		executeCmd("sh", "-c", "home-manager switch")
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
