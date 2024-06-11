package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	port := *flag.Int("PORT", 3000, "The port that the application runs on")
	flag.Parse()

	treeSvr := NewTreeServer(":" + fmt.Sprint(port))
	treeSvr.Run()
}

type Item struct {
	Name     string `json:"name"`
	IsDir    bool   `json:"isDir"`
	AbsPath  string `json:"absPath"`
	Children []Item `json:"children,omitempty"`
}

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
