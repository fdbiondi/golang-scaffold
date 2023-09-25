package scaffold

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Project struct {
	dir       string
	repoOwner string
	name      string
	modName   string
}

func getProjectDir() (string, error) {
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

func getProjectInfo() (Project, error) {
	projectDir, err := getProjectDir()
	if err != nil {
		return Project{}, err
	}

	var repoOwner = ""
	var projectName = ""

	for repoOwner == "" {
		fmt.Print("Enter repository owner: ")
		fmt.Scanln(&repoOwner)
	}

	for projectName == "" {
		fmt.Print("Enter project name: ")
		fmt.Scanln(&projectName)
	}

	modName := fmt.Sprintf("%s/%s/%s", DEFAULT_REMOTE, repoOwner, projectName)

	return Project{
		dir:       projectDir,
		repoOwner: repoOwner,
		name:      projectName,
		modName:   modName,
	}, nil
}
