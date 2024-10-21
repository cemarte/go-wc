package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

// Welcome to the Golang implementation of wc command line tool
func main() {
	// parse command line args
	commandLineArgs := os.Args

	var err error
	var result []string

	wd, error := os.Getwd()
	if error != nil {
		panic(error)
	}

	// if the second argument is -c, count the number of bytes of each subsequent file in args
	switch commandLineArgs[1] {
	case "-m":
		result, err = CountCharacters(os.DirFS(wd), commandLineArgs[2:])
	case "-c":
		result, err = CountBytes(os.DirFS(wd), commandLineArgs[2:])
	case "-l":
		result, err = CountLines(os.DirFS(wd), commandLineArgs[2:])
	case "-w":
		result, err = CountWords(os.DirFS(wd), commandLineArgs[2:])
	}
	if err != nil {
		panic(err)
	}
	printResult(result)
}

func CountCharacters(fsys fs.FS, fileNames []string) ([]string, error) {
	result := make([]string, 0)
	for i := 0; i < len(fileNames); i++ {
		fileName := fileNames[i]
		file, err := fsys.Open(fileName)
		if err != nil {
			return nil, fmt.Errorf("Error reading file %s", fileName)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanRunes)
		charCount := 0
		for scanner.Scan() {
      charCount += 1
		}

		result = append(result, fmt.Sprintf("%d %s", charCount, fileName))

	}
	return result, nil
}

func CountWords(fsys fs.FS, fileNames []string) ([]string, error) {
	result := make([]string, 0)
	for i := 0; i < len(fileNames); i++ {
		fileName := fileNames[i]
		file, err := fsys.Open(fileName)
		if err != nil {
			return nil, fmt.Errorf("Error reading file %s", fileName)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		wordCount := 0
		for scanner.Scan() {
			// process line, count words
			wordCount += processLine(scanner.Text())
		}

		result = append(result, fmt.Sprintf("%d %s", wordCount, fileName))

	}
	return result, nil
}

func processLine(line string) int {
	return len(strings.Fields(line))
}

func CountLines(fsys fs.FS, s []string) ([]string, error) {
	result := make([]string, 0)
	for i := 0; i < len(s); i++ {
		fileName := s[i]

		file, err := fsys.Open(fileName)

		if err != nil {
			return nil, fmt.Errorf("Error reading file %s", file)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lines := 0
		for scanner.Scan() {
			lines += 1
		}
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("Error reading file %s", file)
		}
		result = append(result, fmt.Sprintf("%d %s", lines, fileName))
	}
	return result, nil
}

func CountBytes(fsys fs.FS, s []string) ([]string, error) {
	if len(s) <= 0 {
		panic("Invalid Arguments")
	}

	result := make([]string, 0)

	for i := 0; i < len(s); i++ {
		file := s[i]

		bytes, err := fs.ReadFile(fsys, file)

		if err != nil {
			return nil, fmt.Errorf("Error reading file %s", file)
		}

		result = append(result, fmt.Sprintf("%d %s", len(bytes), file))
	}
	return result, nil
}

func printResult(s []string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", s[i])
	}
}
