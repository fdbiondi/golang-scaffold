package scaffold

import (
	"fmt"
	"os"
)

func getMainModFilename(project Project) string {
	return fmt.Sprintf("%s/src/%s/main.go", project.dir, project.name)
}

func getInternalModFilename(project Project) string {
	return fmt.Sprintf("%[1]s/internal/%[2]s/%[2]s.go", project.dir, INTERNAL_MOD)
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
