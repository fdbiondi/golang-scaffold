package scaffold

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"text/template"

	"github.com/fdbiondi/golang-scaffold/internal/input"
)

const (
	DEFAULT_REMOTE = "github.com"
	INTERNAL_MOD   = "sample"
)

func NewProject(input *input.Input) Project {
	return Project{
		dir:       input.Dir,
		repoOwner: input.Owner,
		name:      input.Project,
		modName:   fmt.Sprintf("%s/%s/%s", DEFAULT_REMOTE, input.Owner, input.Project),
	}
}

func getMainModFilename(project Project) string {
	return fmt.Sprintf("%s/cmd/%s/main.go", project.dir, project.name)
}

func getInternalModFilename(project Project) string {
	return fmt.Sprintf("%[1]s/internal/%[2]s/%[2]s.go", project.dir, INTERNAL_MOD)
}

func fromTemplateToFile(templateName string, outputFilename string, values map[string]string) error {
	tmp, err := template.ParseFiles(templateName)
	if err != nil {
		return errors.New("error parsing file -> " + err.Error())
	}

	file, err := os.Open(outputFilename)
	if err != nil {
		return errors.New("error opening file -> " + err.Error())
	}
	defer file.Close()

	var buf bytes.Buffer
	if err := tmp.Execute(&buf, values); err != nil {
		return errors.New("error creating file content -> " + err.Error())
	}

	err = os.WriteFile(outputFilename, buf.Bytes(), 0644)
	if err != nil {
		return errors.New("error wrinting file -> " + err.Error())
	}

	return nil
}
