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
	modName   string
}

func CreateProject() {
	var project, err = getProjectInfo()
	if err != nil {
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
		if err := os.MkdirAll(project.dir, 0755); err != nil {
			return err
		}
	}

	fmt.Println("Initializing project...")
	cmd := exec.Command("go", "mod", "init", project.modName)
	cmd.Dir = outputDir
	if err := cmd.Run(); err != nil {
		return err
	}

	var internalModule = "sample"
	var dirs = []string{
		project.dir + "/bin",
		project.dir + "/internal/" + internalModule,
		project.dir + "/src/" + project.name,
		project.dir + "/tests",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	fmt.Println("Creating main.go file...")
	_, err := os.Create(fmt.Sprintf("%s/src/%s/main.go", project.dir, project.name))
	if err != nil {
		return err
	}

	fmt.Printf("Creating %s module...\n", internalModule)
	_, err = os.Create(fmt.Sprintf("%[1]s/internal/%[2]s/%[2]s.go", project.dir, internalModule))
	if err != nil {
		return err
	}

	return nil
}
