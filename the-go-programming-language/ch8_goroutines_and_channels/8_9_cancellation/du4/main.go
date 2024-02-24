// The du4 command computes the disk usage of the files in a directory.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// The du4 variant includes cancellation:
// it terminates quickly when the user hits return.

// 在终端执行：
//
// go run ./ch8_goroutines_and_channels/8_9_cancellation/du4/main.go ~ /usr /var
func main() {
	// Determine the initial directories.
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Cancel traversal when input is detected.
	go func() {
		_, _ = os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	// Traverse each root of the file tree in parallel.
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
	tick := time.Tick(500 * time.Millisecond)
	var nFiles, nBytes int64
loop:
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish.
			for range fileSizes {
				// Do nothing.
			}
			return
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

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDir(subDir, n, fileSizes)
		} else {
			info, err := entry.Info()

			var fileSize int64
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "du4: read info error: %v\n", err)
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
	select {
	case sema <- struct{}{}: // acquire token
	case <-done:
		return nil // cancelled
	}
	defer func() { // release token
		<-sema
	}()

	f, err := os.Open(dir)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "du4: open error: %v\n", err)
		return nil
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "du4: close error: %v\n", err)
		}
	}(f)

	entries, err := f.ReadDir(0) // 0 => no limit; read all entries
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "du4: read dir error: %v\n", err)
		// Don't return: Readdir may return partial result.
	}

	return entries
}
