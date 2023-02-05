package main

import (
	"fmt"
	"io"
	"os"
)

// backup takes two arguments, src and dst, representing the source and destination files.
// It copies the contents of the source file to the destination file.
func backup(src, dst string) error {
	// Open the source file for reading
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file for writing
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

// restore takes two arguments, src and dst, representing the source and destination files.
// It copies the contents of the source file to the destination file.
func restore(src, dst string) error {
	// Open the source file for reading
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file for writing
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Set the source and destination file paths
	src := "data.txt"
	backup := "data.txt.bak"
	restore := "data-restored.txt"

	// Perform the backup
	fmt.Println("Backing up data to", backup)
	err := backup(src, backup)
	if err != nil {
		fmt.Println("Error during backup:", err)
		return
	}

	// Perform the restore
	fmt.Println("Restoring data from", backup, "to", restore)
	err = restore(backup, restore)
	if err != nil {
		fmt.Println("Error during restore:", err)
		return
	}

	// Print a success message
	fmt.Println("Backup and restore completed successfully!")
}
