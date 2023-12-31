package cmd

import (
	"fmt"

	"github.com/containerd/continuity/commands"
	"github.com/gfechio/wbso/config"
	"github.com/gfechio/wbso/pkg/commands"
	"github.com/gfechio/wbso/pkg/ticket"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wbso",
	Short: "wbso command line to interact with Git and Ticket managenment systems",
	Run: func(cmd *cobra.Command, args []string) {
		printUsage()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(associateCmd)
	rootCmd.AddCommand(checkoutCmd)
	rootCmd.AddCommand(commitCmd)
	rootCmd.AddCommand(pushCmd)
	rootCmd.AddCommand(closeCmd)
	rootCmd.AddCommand(aliasCmd)
}

func printUsage() {
	fmt.Println("Usage: cmdline_app <subcommand>")
	fmt.Println("Available subcommands:")
	fmt.Println("  - setup: Initialize a new Git repository.")
	fmt.Println("  - associate: Associate something (not implemented yet).")
	fmt.Println("  - checkout: Check out a branch.")
	fmt.Println("  - commit: Commit changes.")
	fmt.Println("  - push: Push changes to the remote repository.")
	fmt.Println("  - close: Close an issue or task.")
	fmt.Println("  - alias: Set up a Git alias.")
	fmt.Println("  - stop: Stop the clock for a given task.")
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Initialize a new Git repository",
	Run: func(cmd *cobra.Command, args []string) {
		commands.Setup(args)
	},
}

var associateCmd = &cobra.Command{
	Use:   "associate",
	Short: "Associate something (not implemented yet)",
	Run: func(cmd *cobra.Command, args []string) {
		commands.Associate(args)
		// Call associate function
	},
}

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Check out a branch",
	Run: func(cmd *cobra.Command, args []string) {
		commands.Checkout(args)
		// Call checkout function
	},
}

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit changes",
	Run: func(cmd *cobra.Command, args []string) {
		commands.Commit(args)
		// Call commit function
	},
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push changes to the remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		commands.Push(args)
		// Call push function
	},
}

var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Close an issue or task",
	Run: func(cmd *cobra.Command, args []string) {
		commands.Close(args)
		// Call close function
	},
}

var stopCmd = &cobra.Command{
	Use:   "close",
	Short: "Stop an issue or task",
	Run: func(cmd *cobra.Command, args []string) {
		commands.Stop()
		// Call close function
	},
}

var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Set up a Git alias",
	Run: func(cmd *cobra.Command, args []string) {
		commands.Alias()
		// Call alias function
	},
}
