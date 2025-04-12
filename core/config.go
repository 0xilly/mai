package core

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Name      string   `json:"name"`
	SrcDir    string   `json:"src_dir"`
	OutDir    string   `json:"out_dir"`
	Verbose   bool     `json:"verbose,omitempty"`
	BuildType string   `json:"build_type"`
	Arch      []string `json:"arch,omitempty"`
	Target    []string `json:"os_target,omitempty"`
}

func LoadConfig() (Config, error) {
	var c Config
	// Load the config file
	f, err := os.Open("mai.build")

	if err != nil {
		return c, err
	}
	defer f.Close()
	// Decode the config file
	j := json.NewDecoder(f)
	err = j.Decode(&c)
	if err != nil {
		return c, err
	}
	// Check if the config file is valid
	if c.Name == "" {
		return c, fmt.Errorf("invalid config file: name is required")
	}
	if c.SrcDir == "" {
		return c, fmt.Errorf("invalid config file: src_dir is required")
	}

	if c.OutDir == "" {
		return c, fmt.Errorf("invalid config file: out_dir is required")
	}

	verbs := []string{"debug", "release"}
	if c.BuildType == "" {
		return c, fmt.Errorf("invalid config file: build_type is required")
	}
	if c.BuildType != "debug" && c.BuildType != "release" {
		return c, fmt.Errorf("invalid config file: build_type must be one of %v", verbs)
	}

	if len(c.Arch) >= 0 {
		for _, a := range c.Arch {
			if a != "amd64" && a != "arm64" {
				return c, fmt.Errorf("invalid config file: arch must be one of amd64 or arm64")
			}
		}
	}

	if len(c.Target) >= 0 {
		for _, t := range c.Target {
			if t != "linux" && t != "windows" && t != "macos" {
				return c, fmt.Errorf("invalid config file: target must be one of linux, windows or macos")
			}
		}
	}
	return c, nil
}
