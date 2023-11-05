package main

import (
	"log"
	"os"

	"github.com/fdbiondi/golang-scaffold/internal/input"
	"github.com/fdbiondi/golang-scaffold/internal/scaffold"
)

func main() {
	userInput, err := input.GetUserInput()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	project := scaffold.NewProject(userInput)

	err = project.CreateDirectory()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = project.CreateStructure()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = project.AddContent()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
