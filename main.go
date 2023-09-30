package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/Masterminds/semver/v3"
	"github.com/blacktop/go-macho"
)

func versionIsSafe(versionStr string) string {

	requiredVersions := []string{"22.3.24", "24.8.3", "25.8.1", "26.2.1"}

	version, err := semver.NewVersion(versionStr)
	if err != nil {
		return "❔"
	}

	for _, reqVersionStr := range requiredVersions {
		reqVersion, err := semver.NewVersion(reqVersionStr)
		if err != nil {
			continue
		}

		if version.Major() == reqVersion.Major() && !version.LessThan(reqVersion) {
			return ("🟢")
		}
	}

	return ("🔴")

}

func main() {

	// Check if the file path is provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: positron /path/to/your/macho/library")
		return
	}
	machoPath := os.Args[1]

	m, err := macho.Open(machoPath)
	if err != nil {
		os.Exit(1)
	}
	defer m.Close()

	cstringSeg := m.Section("__TEXT", "__cstring")
	if cstringSeg == nil {
		fmt.Println("No __TEXT,__cstring segment found")
		return
	}

	content, err := cstringSeg.Data()
	if err != nil {
		panic(err)
	}

	contentStr := string(content)

	pattern := regexp.MustCompile(`Chrome/\d[^ ]+ Electron/[\d\.]+`)

	matches := pattern.FindAllString(contentStr, -1)

	appPattern := regexp.MustCompile(`\.app.*`)
	versPattern := regexp.MustCompile(`^.*Electron/`)

	for _, match := range matches {
		versionStr := versPattern.ReplaceAllString(match, "")
		fmt.Println(appPattern.ReplaceAllString(machoPath, ".app")+":", match, versionIsSafe(versionStr))
	}
}
