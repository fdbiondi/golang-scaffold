package scaffold

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

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
	FromTemplateToFile("./templates/main.txt", getMainModFilename(project), map[string]string{
		"internalMod": INTERNAL_MOD,
		"modName":     project.modName,
	})

	FromTemplateToFile("./templates/mod.txt", getInternalModFilename(project), map[string]string{
		"internalMod": INTERNAL_MOD,
	})

	return nil
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
