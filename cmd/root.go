package cmd

import (
	"log"
	"os/user"

	"github.com/spf13/cobra"
)

var (
	backupDir string

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
	user, err := user.Current()
	if err != nil {
		log.Fatal(err.Error())
	}

	homeDir := user.HomeDir
	backupDir = homeDir + "/.sbu"

	rootCmd.AddCommand(backupCmd)
	rootCmd.AddCommand(restoreCmd)
	rootCmd.AddCommand(versionCmd)
}
