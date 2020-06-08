package main

import (
	"log"
	"os"

	"github.com/uinta-labs/configdir"
)

func main() {
	cfgDir := configdir.New("example-vendor", "example-app")

	allTheDirs := cfgDir.QueryFolders(configdir.All)
	allTheDirs = append(allTheDirs, cfgDir.QueryCacheFolder())

	for _, dir := range allTheDirs {
		info, _ := os.Stat(dir.Path)
		exists := info != nil
		if exists {
			log.Printf("Found %s path at %s", dir.Type, dir.Path)
		} else {
			log.Printf("Target (non-existent) %s path: %s", dir.Type, dir.Path)
		}
	}
}