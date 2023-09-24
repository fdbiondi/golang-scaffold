package scaffold

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	DEFAULT_REMOTE = "github.com"
	DEFAULT_DIR    = "."
)

type Project struct {
	dir       string
	repoOwner string
	name      string
}

func CreateProject() {
	var project, err = getUserInput()
	if err != nil {
		// fmt.Println(err)
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println()

	outputDir, err := getOutputDir(project)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = createStructure(project, outputDir)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}

func getUserInput() (Project, error) {
	flag.Parse()

	projectDir := flag.Arg(0)
	if projectDir == "" {
		projectDir = DEFAULT_DIR
	} else if _, err := os.Stat("./" + projectDir); !os.IsNotExist(err) {
		return Project{}, fmt.Errorf("directory already exists: '%s'", projectDir)
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

	return Project{
		dir:       projectDir,
		repoOwner: repoOwner,
		name:      projectName,
	}, nil
}

func getOutputDir(project Project) (string, error) {
	var outputDir string

	currentPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if project.dir != DEFAULT_DIR {
		outputDir = fmt.Sprintf("%s/%s", currentPath, project.dir)
	}

	return outputDir, nil
}

func createStructure(project Project, outputDir string) error {

	fmt.Println("Creating project directory...")
	if project.dir != DEFAULT_DIR {
		cmd := exec.Command("mkdir", "-p", project.dir)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	fmt.Println("Initializing project...")
	cmd := exec.Command("go", "mod", "init", fmt.Sprintf("%s/%s/%s", DEFAULT_REMOTE, project.repoOwner, project.name))
	cmd.Dir = outputDir
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("Directories scaffolding...")
	cmd = exec.Command("mkdir", "-p", "bin", "src", "internal/sample", "tests", fmt.Sprintf("src/%s", project.name))
	cmd.Dir = outputDir
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("Creating main.go file...")
	cmd = exec.Command("touch", fmt.Sprintf("src/%s/main.go", project.name))
	cmd.Dir = outputDir
	if err := cmd.Run(); err != nil {
		return err
	}

	var module = "sample"

	fmt.Printf("Creating %s module...\n", module)
	cmd = exec.Command("touch", fmt.Sprintf("internal/%s/%s.go", module, module))
	cmd.Dir = outputDir
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
