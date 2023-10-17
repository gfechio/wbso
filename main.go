package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"wbso/cmd"
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

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Initialize a new Git repository",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Setup()
	},
}

var associateCmd = &cobra.Command{
	Use:   "associate",
	Short: "Associate something (not implemented yet)",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Associate()
		// Call associate function
	},
}

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Check out a branch",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Checkout()
		// Call checkout function
	},
}

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit changes",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Commit()
		// Call commit function
	},
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push changes to the remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Push()
		// Call push function
	},
}

var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Close an issue or task",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Close()
		// Call close function
	},
}

var stopCmd = &cobra.Command{
	Use:   "close",
	Short: "Stop an issue or task",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Stop()
		// Call close function
	},
}

var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Set up a Git alias",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Alias()
		// Call alias function
	},
}
