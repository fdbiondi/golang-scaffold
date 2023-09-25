package main

import (
	"log"
	"os"

	"github.com/fdbiondi/golang-scaffold/internal/scaffold"
)

func main() {
	project, err := scaffold.CreateProject()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = scaffold.AddProjectContent(project)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
