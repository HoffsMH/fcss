package fcss

import (
	"os"
	"path/filepath"
	"regexp"
)

var readFile = func(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	return string(bytes), err
}

var getAbs = filepath.Abs

func FindClassesInDir(dir string) []string {
	absDirPath, _ := getAbs(dir)
	result := []string{}

	filepath.Walk(dir, func(dir string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		// check if it is a regular file (not dir)
		if info.Mode().IsRegular() {
			fullPath := filepath.Join(absDirPath, info.Name())
			result = append(result, FindClassesInFile(fullPath)...)
		}
		return nil
	})
	return result
}

func FindClassesInFile(filePath string) []string {
	text, _ := readFile(filePath)
	return FindClassesInText(text)
}

func FindClassesInText(textContent string) []string {
	r := regexp.MustCompile(`[\.\#\&]([A-Za-z-_]+\w+)`)
	out := r.FindAllStringSubmatch(textContent, -1)

	result := []string{}
	for _, j := range out {
		result = append(result, j[1])
	}

	return result
}

func BulkSearch(searchTerms []string) {
	result := 
	for selector := range selectors {
		selector
	}
}

func unique(e []string) []string {
    r := []string{}

    for _, s := range e {
        if !contains(r[:], s) {
            r = append(r, s)
        }
    }
    return r
}

func contains(e []string, c string) bool {
    for _, s := range e {
        if s == c {
            return true
        }
    }
    return false
}
