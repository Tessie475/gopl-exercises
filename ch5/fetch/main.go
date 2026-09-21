// Exercise 5.18: rewrite fetch to use defer to close the writable file,
// without changing its behavior.
//
// The original closed the file with a plain call and carefully kept the error
// from io.Copy if there was one, only falling back to the close error. A bare
// "defer f.Close()" would throw the close error away, changing the behavior.
// So the deferred function is a small closure that assigns to the NAMED return
// value err, preserving the "prefer the Copy error" rule.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// fetch downloads the URL and returns the name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	// Runs when fetch returns. It closes the file, but keeps the error from
	// io.Copy (already in err) if there was one, matching the original.
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)
	return local, n, err
}

func main() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}
		fmt.Printf("%s => %s (%d bytes)\n", url, local, n)
	}
}
