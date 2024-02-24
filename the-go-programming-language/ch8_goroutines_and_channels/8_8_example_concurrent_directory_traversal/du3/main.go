// The du3 command computes the disk usage of the files in a directory.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.
var vFlag = flag.Bool("v", false, "show verbose progress messages")

// 在终端执行：
//
// go run ./ch8_goroutines_and_channels/8_8_example_concurrent_directory_traversal/du3/main.go ~ /usr /var
//
// go run ./ch8_goroutines_and_channels/8_8_example_concurrent_directory_traversal/du3/main.go -v ~ /usr /var
func main() {
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nFiles, nBytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nFiles++
			nBytes += size
		case <-tick:
			printDiskUsage(nFiles, nBytes)
		}
	}

	printDiskUsage(nFiles, nBytes) // final totals
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			walkDir(subDir, n, fileSizes)
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

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.DirEntry {
	sema <- struct{}{} // acquire token
	defer func() {     // release token
		<-sema
	}()

	entries, err := os.ReadDir(dir)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries
}
