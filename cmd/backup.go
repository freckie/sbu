package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(backupCmd)
}

var backupCmd = &cobra.Command{
	Use:     "backup",
	Short:   "Backup files.",
	Long:    `Backup the target files.`,
	Aliases: []string{"b", "bu"},
	RunE:    backup,
}

func backup(cmd *cobra.Command, args []string) error {
	oldPath, _ := filepath.Abs(args[0])
	newPath := filepath.Join(backupDir, oldPath)

	newDir, newFile := filepath.Split(newPath)

	// Make directories
	err := os.MkdirAll(newDir, os.ModePerm)
	if err != nil {
		fmt.Println("making directories failed.", err)
		return err
	}

	// Move the target file
	err = os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Println("moving file failed.", err)
		return err
	}

	fmt.Printf("%s moved to %s successfully.\n", args[0], newFile)
	return nil
}
