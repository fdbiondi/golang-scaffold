package scaffold

import (
	"errors"
	"os"
	"os/exec"
)

type Project struct {
	dir       string
	repoOwner string
	name      string
	modName   string
}

func (project Project) CreateDirectory() error {
	if _, err := os.Stat(project.dir); os.IsNotExist(err) {
		if err := os.MkdirAll(project.dir, 0755); err != nil {
			return errors.New("failed to create project directory")
		}
	}

	files, err := os.ReadDir(project.dir)
	if err != nil {
		return errors.New("bad path to project directory")
	}

	if len(files) > 0 {
		return errors.New("directory is not empty")
	}

	return nil
}

func (project Project) CreateStructure() error {
	cmd := exec.Command("go", "mod", "init", project.modName)
	cmd.Dir = project.dir
	if err := cmd.Run(); err != nil {
		return errors.New("failed to create go main module")
	}

	var dirs = []string{
		project.dir + "/bin",
		project.dir + "/internal/" + INTERNAL_MOD,
		project.dir + "/cmd/" + project.name,
		project.dir + "/tests",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return errors.New("failed to create directory " + dir + " : " + err.Error())
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

func (project Project) AddContent() error {
	fromTemplateToFile("./templates/main.txt", getMainModFilename(project), map[string]string{
		"internalMod": INTERNAL_MOD,
		"modName":     project.modName,
	})

	fromTemplateToFile("./templates/mod.txt", getInternalModFilename(project), map[string]string{
		"internalMod": INTERNAL_MOD,
	})

	return nil
}
