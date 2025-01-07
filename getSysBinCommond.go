package main

import (
	"fmt"
	"os"

	// "os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		fmt.Println("This script is intended to run on Linux or macOS.")
		return
	}

	if len(os.Args) > 1 {
		fmt.Println("Usage: go run main.go")
		return
	}

	fmt.Println("GO VERSION:", runtime.Version())
	fmt.Println()

	path := os.Getenv("PATH")
	dirs := strings.Split(path, ":")

	for _, dir := range dirs {
		files, err := getFilesByDir(dir)
		if err != nil {
			// fmt.Println("Error:", err.Error())
			continue
		}
		fmt.Println("====dir:", dir, "====start====")

		for _, file := range files {
			// fileInfo, err := os.Stat(file)
			// if err != nil {
			// 	continue // Skip files that can't be accessed
			// }

			// if fileInfo.Mode()&0111 == 0 {
			// 	continue // Skip non-executable files
			// }

			// if fileInfo.IsDir() {
			// 	continue // Skip directories
			// }

			fmt.Println(file)
		}
		fmt.Println("====dir:", dir, "====end====")
	}
}

func getFilesByDir(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && (info.Mode()&0111 != 0) {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
