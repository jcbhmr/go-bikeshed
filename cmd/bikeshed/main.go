package main

import (
	"log"
	"runtime/debug"

	"github.com/adrg/xdg"
)

func main() {
	log.SetFlags(0)
	_, err := xdg.DataFile("go-bikeshed/" + Version)
	if err != nil {
		log.Fatal(err)
	}
	// bikeshedinstall.ExtractTo(bikeshedInstall)
	// log.Printf("Bikeshed installed to %s", bikeshedInstall)
}

