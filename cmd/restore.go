package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	restoreFlagRecursive bool

	restoreCmd = &cobra.Command{
		Use:     "restore",
		Short:   "Restore the backed-up files.",
		Long:    `Restore the backed-up files.`,
		Aliases: []string{"r", "rs"},
		Args:    cobra.MinimumNArgs(1),
		RunE:    restore,
	}
)

func init() {
	restoreCmd.Flags().BoolVarP(&restoreFlagRecursive, "recursive", "r", false, "Restore the directory recursively")
}

func restore(cmd *cobra.Command, args []string) error {
	var err error
	var result string

	// Valid targets
	fileInfos := make([]os.FileInfo, len(args))
	for idx := range args {
		tempPath, _ := filepath.Abs(args[idx])
		fileInfos[idx], err = os.Stat(filepath.Join(backupDir, tempPath))
		if err != nil {
			return fmt.Errorf("invalid filepath : %s", args[idx])
		}

		// If the target is directory
		if fileInfos[idx].IsDir() && !restoreFlagRecursive {
			return fmt.Errorf("cannot handle directories without the recursive flag : %s", err)
		}
	}

	// Restore targets
	for _, file := range fileInfos {
		result, err = _restore(file.Name())
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Restored the file(s) or directory successfully :", result)
	}

	return nil
}

func _restore(file string) (string, error) {
	var err error

	restorePath, _ := filepath.Abs(file)
	backupPath := filepath.Join(backupDir, restorePath)

	_, restoreFile := filepath.Split(restorePath)

	// Check if the file exists
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		return "", fmt.Errorf("no specific file or directory found : %s", err)
	}

	// Move the target file
	err = os.Rename(backupPath, restorePath)
	if err != nil {
		return "", fmt.Errorf("failed to move the target : %s", err)
	}

	return restoreFile, nil
}
