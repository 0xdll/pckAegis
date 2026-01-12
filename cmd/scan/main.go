package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"vuln-scanner/internal/parser"
	"vuln-scanner/internal/report"
	"vuln-scanner/internal/scanner"
)

func main() {
	filePath := flag.String("file", "", "Path to requirements.txt or package.json")
	ecosystem := flag.String("eco", "PyPI", "Ecosystem (PyPI or npm)")
	jsonPath := flag.String("out", "", "Path to save JSON report (e.g. report.json)")
	strictMode := flag.Bool("strict", true, "Exit 1 if vulns are found")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Usage: pckaegis -file <path> -eco <PyPI|npm> [-out report.json]")
		os.Exit(1)
	}

	var deps []parser.Dependency
	var err error
	if *ecosystem == "PyPI" {
		deps, err = parser.ParseRequirements(*filePath)
	} else {
		deps, err = parser.ParsePackageJSON(*filePath)
	}

	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}

	var results []map[string]string
	vulnerabilitiesFound := false

	for _, d := range deps {
		vulns, _ := scanner.QueryOSV(d.Name, d.Version, *ecosystem)
		if len(vulns) > 0 {
			vulnerabilitiesFound = true
			for _, v := range vulns {
				results = append(results, map[string]string{
					"pkg": d.Name, "ver": d.Version, "id": v.ID, "summary": v.Summary,
				})
			}
		}
	}

	if *jsonPath != "" {
		file, _ := json.MarshalIndent(results, "", "  ")
		_ = os.WriteFile(*jsonPath, file, 0644)
		fmt.Printf("JSON report saved to: %s\n", *jsonPath)
	}

	report.DisplayResults(results)

	if vulnerabilitiesFound && *strictMode {
		fmt.Println("\n[!] Blocking CI/CD due to vulnerabilities.")
		os.Exit(1)
	}
}
