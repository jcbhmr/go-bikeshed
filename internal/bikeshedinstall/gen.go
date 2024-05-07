//go:build ignore

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	_ "embed"
)

//go:embed VERSION.txt
var versionTxt string
var version = strings.TrimSpace(versionTxt)

func main() {
	log.SetFlags(0)

	log.Printf("version: %s", version)

	targets := []string{
		"arm64-apple-darwin",
		"x86_64-apple-darwin",
		"x86_64-pc-windows-msvc",
		"x86_64-unknown-linux-gnu",
	}
	for _, target := range targets {
		var url string
		if strings.Contains(target, "windows") {
			url = "https://github.com/jcbhmr/bikeshed-builder/releases/download/v" + version + "/bikeshed-" + target + ".zip"
		} else {
			url = "https://github.com/jcbhmr/bikeshed-builder/releases/download/v" + version + "/bikeshed-" + target + ".tar.gz"
		}

		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			log.Fatal(fmt.Errorf("unexpected status code: %s", res.Status))
		}

		dest := path.Base(url)
		file, err := os.Create(dest)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = io.Copy(file, res.Body)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("downloaded %s to %s", url, dest)
	}
}
