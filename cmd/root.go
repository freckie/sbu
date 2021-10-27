package cmd

import (
	"github.com/spf13/cobra"
)

var (
	backupDir = "~/.sbu"

	rootCmd = &cobra.Command{
		Use:   "sbu",
		Short: "A Simple File Backup Tool",
		Long:  `sbu is a simple file backup tool for unix systems.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(backupCmd)
	rootCmd.AddCommand(restoreCmd)
	rootCmd.AddCommand(versionCmd)
}
