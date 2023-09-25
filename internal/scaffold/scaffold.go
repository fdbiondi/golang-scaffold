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
func CreateProject() (Project, error) {
	var project, err = getProjectInfo()
	if err != nil {
		return Project{}, err
	}

	fmt.Println()

	outputDir, err := getOutputDir(project)
	if err != nil {
		return Project{}, err
	}

	err = createStructure(project, outputDir)
	if err != nil {
		return Project{}, err
	}

	return project, nil
}

func AddContent(project Project) error {
	/* tmp, err := template.ParseFiles("./templates/main.txt")
	if err != nil {
		return errors.New("error parsing file -> " + err.Error())
	}

	file, err := os.Open(getMainModFilename(project))
	if err != nil {
		return errors.New("error opening file -> " + err.Error())
	}
	defer file.Close()

	config := map[string]string{
		"internalMod": INTERNAL_MOD,
		"modName":     project.modName,
	}

	var buf bytes.Buffer
	if err := tmp.Execute(&buf, config); err != nil {
        return errors.New("error creating file content -> " + err.Error())
	}

	err = os.WriteFile(getMainModFilename(project), buf.Bytes(), 0644)
	if err != nil {
        return errors.New("error wrinting file -> " + err.Error())
	} */

	/* modFileContent := []byte(`package

	import "fmt"

	func HelloWorld() {
		fmt.Println("Hello World!")
	}`)

		err = os.WriteFile(getInternalModFilename(project), modFileContent, 0644)
		if err != nil {
			return err
		} */

	return nil
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
	if project.dir != DEFAULT_DIR {
		if err := os.MkdirAll(project.dir, 0755); err != nil {
			return errors.New("failed to create project directory")
		}
	}

	cmd := exec.Command("go", "mod", "init", project.modName)
	cmd.Dir = outputDir
	if err := cmd.Run(); err != nil {
		return errors.New("failed to create go main module")
	}

	var dirs = []string{
		project.dir + "/bin",
		project.dir + "/internal/" + INTERNAL_MOD,
		project.dir + "/src/" + project.name,
		project.dir + "/tests",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	mainFile, err := os.Create(getMainModFilename(project))
	if err != nil {
		return errors.New("failed to create main.go file")
	}
	defer mainFile.Close()

	filename := getInternalModFilename(project)
	modFile, err := os.Create(filename)
	if err != nil {
		return errors.New("failed to create internal module")
	}
	defer modFile.Close()

	return nil
}
