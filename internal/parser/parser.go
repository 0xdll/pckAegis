package parser

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
)

type Dependency struct {
	Name    string
	Version string
}

func ParseRequirements(path string) ([]Dependency, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var deps []Dependency
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Split(line, "==")
		if len(parts) == 2 {
			deps = append(deps, Dependency{Name: parts[0], Version: parts[1]})
		}
	}
	return deps, nil
}

func ParsePackageJSON(path string) ([]Dependency, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data struct {
		Deps    map[string]string `json:"dependencies"`
		DevDeps map[string]string `json:"devDependencies"`
	}

	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	var deps []Dependency
	for name, ver := range data.Deps {
		deps = append(deps, Dependency{Name: name, Version: strings.TrimLeft(ver, "^~")})
	}
	return deps, nil
}