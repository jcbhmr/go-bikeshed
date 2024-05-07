//go:build ignore

package main

import (
	_ "embed"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	var taskName string
	if len(os.Args) >= 2 {
		taskName = os.Args[1]
	} else {
		log.Fatal("no task")
	}
	task, ok := map[string]func() error{}[taskName]
	if !ok {
		log.Fatal("no such task")
	}
	err := task()
	if err != nil {
		log.Fatal(err)
	}
}
