package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	backupFlagRecursive bool

	backupCmd = &cobra.Command{
		Use:     "backup [flags] [TARGETS]",
		Short:   "Backup files.",
		Long:    `Backup the target files.`,
		Aliases: []string{"b", "bu"},
		Args:    cobra.MinimumNArgs(1),
		RunE:    backup,
	}
)

func init() {
	backupCmd.Flags().BoolVarP(&backupFlagRecursive, "recursive", "r", false, "Back up the directory recursively")
}

func backup(cmd *cobra.Command, args []string) error {
	var err error
	var result string

	// Valid targets
	fileInfos := make([]os.FileInfo, len(args))
	for idx := range args {
		fileInfos[idx], err = os.Stat(args[idx])
		if err != nil {
			return fmt.Errorf("invalid filepath : %s", args[idx])
		}

		// If the target is directory
		if fileInfos[idx].IsDir() && !backupFlagRecursive {
			return fmt.Errorf("cannot handle directories without the recursive flag : %s", args[idx])
		}
	}

	// Back up targets
	for _, file := range fileInfos {
		result, err = _backup(file.Name())
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Backed up the file(s) or directory successfully :", result)
	}

	return nil
}

func _backup(file string) (string, error) {
	var err error

	oldPath, _ := filepath.Abs(file)
	newPath := filepath.Join(backupDir, oldPath)

	newDir, newFile := filepath.Split(newPath)

	// Make directories
	err = os.MkdirAll(newDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create the directory : %s", err)
	}

	// Move the target file or directory
	err = os.Rename(oldPath, newPath)
	if err != nil {
		return "", fmt.Errorf("failed to move the target : %s", err)
	}

	return newFile, nil
}
