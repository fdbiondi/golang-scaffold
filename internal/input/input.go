package input

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const DEFAULT_DIR = "."

type Input struct {
	Dir     string
	Owner   string
	Project string
}

func getCurrentDir() (string, error) {
	flag.Parse()
	path := flag.Arg(0)

	projectPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	currentPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	var dir string

	if path == DEFAULT_DIR {
		dir = currentPath
	} else {
		dir = projectPath
	}

	return dir, nil
}

func GetUserInput() (*Input, error) {
	dir, err := getCurrentDir()
	if err != nil {
		return nil, err
	}

	var owner string
	var project string

	for owner == "" {
		fmt.Print("Enter repository owner: ")
		fmt.Scanln(&owner)
	}

	for project == "" {
		fmt.Print("Enter project name: ")
		fmt.Scanln(&project)
	}

	return &Input{
		Dir:     dir,
		Owner:   owner,
		Project: project,
	}, nil
}
