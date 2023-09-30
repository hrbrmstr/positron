package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/blacktop/go-macho"
)

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

	for _, match := range matches {
		fmt.Println(appPattern.ReplaceAllString(machoPath, ".app")+":", match)
	}
}
