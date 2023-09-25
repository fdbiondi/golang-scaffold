package scaffold

import (
	"flag"
	"fmt"
	"os"
)

type Project struct {
	dir       string
	repoOwner string
	name      string
	modName   string
}

func getProjectDir() (string, error) {
	flag.Parse()

	projectDir := flag.Arg(0)
	if projectDir == "" {
		projectDir = DEFAULT_DIR

	} else if _, err := os.Stat(projectDir); !os.IsNotExist(err) {
		return "", fmt.Errorf("a directory named '%s' already exists", projectDir)
	}

	if _, err := os.Stat(projectDir + "/go.mod"); !os.IsNotExist(err) {
		pwd, err := os.Getwd()
		if err != nil {
			pwd = projectDir
		}

		return "", fmt.Errorf("go module already exists on '%s'", pwd)
	}

	return projectDir, nil
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
