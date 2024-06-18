package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// The default port the application will run on
// if the PORT flag is NOT specified.
const defaultPort = 3000

func main() {
	var (
		port = flag.Int("PORT", defaultPort, "The port that the application runs on")
		wd   = flag.String("WD", "", "The working directory from where all nested files are served")
	)
	flag.Parse()

	if *wd == "" {
		// If the WD flag is omitted,
		// the directory where the application is running
		// becomes the working directory.
		_wd, err := os.Getwd()
		if err != nil {
			log.Fatal("error getting work directory: " + err.Error())
		}
		*wd = _wd
	}

	treeSvr := NewTreeServer(":"+fmt.Sprint(*port), *wd)
	treeSvr.Run()
}

// An Item represents either a file or a folder in the file system.
// An Item can have more Items as it's Children if it's a folder (IsDir == true).
type Item struct {
	Name     string `json:"name"`
	IsDir    bool   `json:"isDir"`
	AbsPath  string `json:"absPath"`
	URL      string `json:"url"`
	Children []Item `json:"children"`
}

// Returns the file system tree, reprenented as Items
func readDir(dirP string, wd string) []Item {
	// Open the directory
	dir, err := os.Open(dirP)
	if err != nil {
		log.Fatal("error opening directory: " + err.Error())
	}
	defer dir.Close()

	// Read the directory entries
	entries, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal("error reading directory contents: " + err.Error())
	}

	var items []Item
	for _, entry := range entries {
		var (
			name     = entry.Name()
			isDir    = entry.IsDir()
			absPath  = formatAbsPath(dirP, name)
			url      = formatURL(absPath, wd)
			children []Item
		)

		if isDir {
			children = readDir(absPath, wd)
		} else {
			children = nil
		}

		items = append(items, Item{
			Name:     name,
			IsDir:    isDir,
			AbsPath:  absPath,
			URL:      url,
			Children: children,
		})
	}

	return items
}

func formatAbsPath(path string, fileName string) string {
	return replaceAll(rmvTrailingSlash(path), "/", "\\") + "\\" + fileName
}

func rmvTrailingSlash(path string) string {
	if len(path) > 0 && path[len(path)-1] == '/' {
		return path[:len(path)-1]
	}
	return path
}

func formatURL(absPath string, wd string) string {
	return replaceAll(replaceAll(absPath, "\\", "/"), replaceAll(wd, "\\", "/"), "")
}

func replaceAll(input string, oldChar string, newChar string) string {
	return strings.ReplaceAll(input, string(oldChar), string(newChar))
}
