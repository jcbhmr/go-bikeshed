package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jcbhmr/go-bikeshed/internal/bikeshedinstall"
	"github.com/jcbhmr/go-bikeshed/internal/replacexec"
)

func main() {
	log.SetFlags(0)

	if _, err := os.Stat(bikeshedinstall.Dest); err != nil { // if not exist
		err := bikeshedinstall.Install()
		if err != nil {
			log.Fatal(err)
		}
	}

	var exeExt string
	if runtime.GOOS == "windows" {
		exeExt = ".exe"
	} else {
		exeExt = ""
	}
	bikeshed := filepath.Join(bikeshedinstall.Dest, "bikeshed"+exeExt)

	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	// try to mv $0 $0.bak && ln -s bikeshed $0
	// revert if failed
	err = os.Rename(exe, exe+".bak")
	if err == nil {
		err = os.Symlink(bikeshed, exe)
		if err != nil {
			err = os.Rename(exe+".bak", exe)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			_ = os.Remove(exe + ".bak")
		}
	}

	err = replacexec.Replacexec(bikeshed, os.Args...)
	if err != nil {
		log.Fatal(err)
	}
}
