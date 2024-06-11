package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
			log.Fatal(ServerError{Error: "error getting work directory"})
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
	Children []Item `json:"children"`
}

// Returns the file system tree, reprenented as Items
func readDir(dirP string) []Item {
	// Open the directory
	dir, err := os.Open(dirP)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	// Read the directory entries
	entries, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	var items []Item
	for _, entry := range entries {
		var (
			name    = entry.Name()
			isDir   = entry.IsDir()
			absPath = dirP + "\\" + name
		)

		children := []Item{}
		if isDir {
			children = readDir(absPath)
		}

		items = append(items, Item{
			Name:     name,
			IsDir:    isDir,
			AbsPath:  absPath,
			Children: children,
		})
	}

	return items
}
