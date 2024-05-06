package bikeshedinstall

import (
	"log"
	"runtime/debug"

	"github.com/adrg/xdg"
)

func init() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatal("bikeshedinstall: could not read build info")
	}
	GoBikeshedVersion = bi.Main.Version
}

func Run() {
	bikeshedInstall, err := xdg.DataFile("go-bikeshed/" + GoBikeshedVersion)
}
