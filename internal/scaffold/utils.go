package scaffold

import (
	"fmt"
)

func getMainModFilename(project Project) string {
	return fmt.Sprintf("%s/src/%s/main.go", project.dir, project.name)
}

func getInternalModFilename(project Project) string {
	return fmt.Sprintf("%[1]s/internal/%[2]s/%[2]s.go", project.dir, INTERNAL_MOD)
}
