package bikeshedinstall

import (
	"bytes"
	"context"
	_ "embed"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/adrg/xdg"
	"github.com/codeclysm/extract/v3"
)

//go:generate go run gen.go

//go:embed VERSION.txt
var versionTxt string
var Version = strings.TrimSpace(versionTxt)

var Dest = filepath.Join(xdg.DataHome, "go-bikeshed", Version)

func Install() error {
	err := extract.Archive(context.TODO(), bytes.NewBuffer(bikeshedArchive), Dest, nil)
	if err != nil {
		return err
	}

	var exeExt string
	if runtime.GOOS == "windows" {
		exeExt = ".exe"
	} else {
		exeExt = ""
	}

	err = os.Chmod(filepath.Join(Dest, "bikeshed"+exeExt), 0755)
	if err != nil {
		return err
	}

	return nil
}
