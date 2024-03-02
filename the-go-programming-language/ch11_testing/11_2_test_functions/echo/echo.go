// Echo prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "separator")
)
var out io.Writer = os.Stdout

// 在终端执行：
//
// go run ./ch11_testing/11_2_test_functions/echo/echo.go -n=false -s="," "a" "b" "c" "d"
func main() {
	flag.Parse()
	if err := echo(!*n, *s, flag.Args()); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func echo(newline bool, sep string, args []string) error {
	_, _ = fmt.Fprint(out, strings.Join(args, sep))
	if newline {
		_, _ = fmt.Fprintln(out)
	}
	return nil
}
