package command

import (
	"encoding/json"
	"fmt"
	"mai/core"
	"os"
	"path/filepath"
	"strings"
)

type InitBuild struct {
}

func NewInitBuild() *InitBuild { return &InitBuild{} }

func (*InitBuild) Name() string        { return "init" }
func (*InitBuild) ShortName() string   { return "i" }
func (*InitBuild) Description() string { return "Initialize a new Mai project" }
func (*InitBuild) Usage() string       { return "init [<name>] - Initialize a new Mai project" }

func (*InitBuild) Run(args []string) error {
	if len(args) == 0 {
		wd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current working directory: %w", err)
		}

		dirName := filepath.Base(wd)
		if err := createProjectDir(dirName); err != nil {
			return err
		}
	}

	name := args[0]

	if _, err := os.Stat(name); !os.IsNotExist(err) {
		return fmt.Errorf("project directory already exists: %s", name)
	}

	if err := createProjectDir(name); err != nil {
		return err
	}

	fmt.Printf("Created new Mai project in %s\n", name)
	return nil
}

func createProjectDir(name string) error {
	// Create the project directory
	if err := os.MkdirAll(name, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Create the mai.build file
	config := core.Config{
		Name:      name,
		SrcDir:    "src",
		OutDir:    "out",
		BuildType: "debug",
	}
	configFile, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	if err := os.WriteFile(filepath.Join(name, "mai.build"), configFile, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	//create gitignore file
	var ignore strings.Builder
	ignore.WriteString("out")

	writeGitIgnore := []byte(ignore.String())
	if err := os.WriteFile(filepath.Join(name, ".gitignore"), writeGitIgnore, 0644); err != nil {
		return fmt.Errorf("failed to write .gitignore file: %w", err)
	}

	// Create the src directory
	if err := os.MkdirAll(filepath.Join(name, "src"), 0755); err != nil {
		return fmt.Errorf("failed to create src directory: %w", err)
	}

	if err := os.MkdirAll(filepath.Join(name, "out"), 0755); err != nil {
		return fmt.Errorf("failed to create out directory: %w", err)
	}

	// Create a main.mai file
	mainMai := `pkg main
use std:fmt

@main
ex main :: fn(): void -> {
	@println("Hello, Mai!")
}
`
	if err := os.WriteFile(filepath.Join(name, "src", "main.mai"), []byte(mainMai), 0644); err != nil {
		return fmt.Errorf("failed to write main.mai file: %w", err)
	}

	return nil
}
