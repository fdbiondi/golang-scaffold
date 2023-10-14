package scaffold

import (
	"errors"
	"os"
	"os/exec"
)

func CreateProject() (Project, error) {
	var project, err = getProjectInfo()
	if err != nil {
		return Project{}, err
	}

	err = createProjectDirectory(project)
	if err != nil {
		return Project{}, err
	}

	err = createProjectStructure(project)
	if err != nil {
		return Project{}, err
	}

	return project, nil
}

func AddProjectContent(project Project) error {
	FromTemplateToFile("./templates/main.txt", getMainModFilename(project), map[string]string{
		"internalMod": INTERNAL_MOD,
		"modName":     project.modName,
	})

	FromTemplateToFile("./templates/mod.txt", getInternalModFilename(project), map[string]string{
		"internalMod": INTERNAL_MOD,
	})

	return nil
}

func createProjectDirectory(project Project) error {
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

func createProjectStructure(project Project) error {
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
