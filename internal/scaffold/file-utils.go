package scaffold

import (
	"bytes"
	"errors"
	"os"
	"text/template"
)

func FromTemplateToFile(templateName string, outputFilename string, values map[string]string) error {
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
