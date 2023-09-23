package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// func Spinner(delay time.Duration) {
//     for !StopSpinner{
//         for _, r := range `-\|/` {
//             fmt.Printf("\r%c", r)
//             time.Sleep(delay)
//         }
//     }
// }

const (
	DEFAULT_REMOTE = "github.com"
	DEFAULT_DIR    = "."
)

func main() {
	flag.Parse()

	projectDir := flag.Arg(0)
	if projectDir == "" {
		projectDir = DEFAULT_DIR
	} else if _, err := os.Stat("./" + projectDir); !os.IsNotExist(err) {
		fmt.Printf("Directory already exists: '%s'\n", projectDir)
		os.Exit(1)
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

	fmt.Println()

	fmt.Println("Creating project directory...")
	if projectDir != DEFAULT_DIR {
		cmd := exec.Command("mkdir", "-p", projectDir)
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}

		// cmd = exec.Command("cd", projectDir)
		// cmd.Run()
	}

	currentPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	var outputDir string

	if (projectDir != DEFAULT_DIR) {
		outputDir = fmt.Sprintf("%s/%s", currentPath, projectDir)
	}

	fmt.Println("Initializing project...")
	cmd := exec.Command("go", "mod", "init", fmt.Sprintf("%s/%s/%s", DEFAULT_REMOTE, repoOwner, projectName))
	cmd.Dir = outputDir
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Directories scaffolding...")
	cmd = exec.Command("mkdir", "-p", "bin", "src", "internal", "tests")
	cmd.Dir = outputDir
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating main.go file...")
	filename := fmt.Sprintf("src/%s/main.go", projectName)
	cmd = exec.Command("touch", filename)
	cmd.Dir = outputDir
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating sample module...")
	cmd = exec.Command("touch", "internal/sample/sample.go")
	cmd.Dir = outputDir
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}
