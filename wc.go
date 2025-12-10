package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	var isDir bool
	var path string
	flag.StringVar(&path, "p", "", "file or directory path")
	flag.BoolVar(&isDir, "d", false, "is directory?")
	flag.Parse()
	if isDir {
		fmt.Println(countDir(path))
	} else {
		fmt.Println(count(path))
	}

}

func countDir(p string) int {
	wc := 0
	filepath.WalkDir(p, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		wc += count(path)
		return nil
	})
	return wc
}

func count(path string) int {
	wc := 0
	f, _ := os.Open(path)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wc++
	}

	return wc

}
