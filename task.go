//go:build ignore

package main

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"
)

//go:embed go.mod
var goModBytes []byte
var Version = string(regexp.MustCompile(`// VERSION: (.+)`).FindSubmatch(goModBytes)[1])
var BikeshedBuilderVersion = strings.SplitN(Version, "+", 2)[1]
var BikeshedVersion = strings.SplitN(BikeshedBuilderVersion, "+", 2)[1]

func WgetTemplate() error {
	t, err := template.New("wget-template").Parse(os.Args[2])
	if err != nil {
		return err
	}
	urlBuilder := &strings.Builder{}
	err = t.Execute(urlBuilder, map[string]string{
		"Version":                Version,
		"BikeshedBuilderVersion": BikeshedBuilderVersion,
		"BikeshedVersion":        BikeshedVersion,
	})
	if err != nil {
		return err
	}
	urlString := urlBuilder.String()

	urlObject, err := url.Parse(urlString)
	if err != nil {
		return err
	}

	log.Printf("Downloading %s", urlObject)
	res, err := http.Get(urlObject.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download %s: %s", urlObject, res.Status)
	}
	log.Printf("Got status %d %s", res.StatusCode, res.Status)

	log.Printf("Saving to %s", path.Base(urlObject.Path))
	filename := path.Base(urlObject.Path)
	dest, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dest.Close()
	_, err = io.Copy(dest, res.Body)
	if err != nil {
		return err
	}
	log.Printf("Saved to %s", filename)

	return nil
}

func main() {
	log.SetFlags(0)
	var taskName string
	if len(os.Args) >= 2 {
		taskName = os.Args[1]
	} else {
		log.Fatal("no task")
	}
	task, ok := map[string]func() error{
		"wget-template": WgetTemplate,
	}[taskName]
	if !ok {
		log.Fatal("no such task")
	}
	err := task()
	if err != nil {
		log.Fatal(err)
	}
}
