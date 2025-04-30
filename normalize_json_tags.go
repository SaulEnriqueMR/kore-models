package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory: %v", err)
	}

	dirRelPath := "models"
	rootDir := filepath.Join(cwd, dirRelPath)

	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process .go files
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		err = processFile(path)
		if err != nil {
			log.Printf("error processing file %s: %v", path, err)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("error walking path %q: %v", rootDir, err)
	}
}

func processFile(filename string) error {
	bsonRe := regexp.MustCompile(`\bbson:"([^"]+)"`)
	jsonRe := regexp.MustCompile(`\s*\bjson:"([^"]+)"`)

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var output bytes.Buffer
	scanner := bufio.NewScanner(file)

	changed := false
	currentLine := -1
	for scanner.Scan() {
		currentLine++
		line := scanner.Text()

		bsonMatch := bsonRe.FindStringSubmatch(line)
		if bsonMatch == nil {
			output.WriteString(line + "\n")
			continue
		}

		bsonFull := bsonMatch[1]
		if strings.Contains(bsonFull, "inline") {
			output.WriteString(line + "\n")
			continue
		}

		bsonParts := strings.SplitN(bsonFull, ",", 2)
		bsonOptions := ""
		if len(bsonParts) > 1 {
			bsonOptions = bsonParts[1]
		}

		jsonTagName := bsonParts[0]

		jsonMatch := jsonRe.FindStringSubmatch(line)
		if jsonMatch != nil {
			// Remove existing json tag
			line = jsonRe.ReplaceAllString(line, "")
		}

		// No json tag exists, add a new json tag with bson config
		newJsonFull := jsonTagName
		if bsonOptions != "" {
			newJsonFull = fmt.Sprintf("%s,%s", newJsonFull, bsonOptions)
		}
		newJsonTag := fmt.Sprintf(` json:"%s"`, newJsonFull)

		idx := strings.Index(line, `bson:"`)
		if idx != -1 {
			endIdx := idx + strings.Index(line[idx:], `"`) + 1 // after starting quote
			endIdx += strings.Index(line[endIdx:], `"`) + 1    // after ending quote

			pre := line[:endIdx]
			post := line[endIdx:]

			line = pre + newJsonTag + post
			changed = true
		}
		output.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if changed {
		newFilename := filename

		err = os.WriteFile(newFilename, output.Bytes(), 0644)
		if err != nil {
			return fmt.Errorf("failed to write new file %s: %w", newFilename, err)
		}
	}

	return nil
}
