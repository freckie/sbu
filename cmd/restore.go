package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(restoreCmd)
}

var restoreCmd = &cobra.Command{
	Use:     "restore",
	Short:   "Restore the backed-up files.",
	Long:    `Restore the backed-up files.`,
	Aliases: []string{"r", "rs"},
	RunE:    restore,
}

func restore(cmd *cobra.Command, args []string) error {
	restorePath, _ := filepath.Abs(args[0])
	backupPath := filepath.Join(backupDir, restorePath)

	_, restoreFile := filepath.Split(restorePath)

	// Check if the file exists
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		fmt.Println("No specific file found.", err)
		return err
	}

	// Move the target file
	err := os.Rename(backupPath, restorePath)
	if err != nil {
		fmt.Println("moving file failed.", err)
		return err
	}

	fmt.Printf("%s has been restored successfully.\n", restoreFile)
	return nil
}
