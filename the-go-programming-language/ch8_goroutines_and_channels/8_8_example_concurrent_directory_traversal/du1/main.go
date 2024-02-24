// The du1 command computes the disk usage of the files in a directory.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// 在终端执行：
//
//	go run ./ch8_goroutines_and_channels/8_8_example_concurrent_directory_traversal/du1/main.go ~ /usr /var
func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Print the results.
	var nFiles, nBytes int64
	for size := range fileSizes {
		nFiles++
		nBytes += size
	}

	printDiskUsage(nFiles, nBytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walkDir(subDir, fileSizes)
		} else {
			info, err := entry.Info()

			var fileSize int64
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "du1: read info error: %v\n", err)
				fileSize = 0
			} else {
				fileSize = info.Size()
			}
			fileSizes <- fileSize
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries
}
