// Copyright (C) 2017 David Capello
//
// This file is released under the terms of the MIT license.
// Read LICENSE.txt for more information.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

var ignoreCase bool
var oneLevel bool
var verbose bool
var filenameFilter = new(FilenameFilter)

func isBinary(line string) bool {
	return strings.IndexByte(line, byte(0)) >= 0
}

func matchPatterns(line string) int {
	firstIndex := -1
	matches := 0
	for _, arg := range flag.Args() {
		i := strings.Index(line, arg)
		if i >= 0 {
			if (firstIndex < 0 ||
			    firstIndex > i) {
				firstIndex = i
			}
			matches++
		}
	}
	if matches == flag.NArg() {
		return firstIndex
	}
	return -1
}

func searchInFileContent(fn string) error {
	file, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		if isBinary(line) {
			if verbose { fmt.Println(fn + " is binary") }
			break
		}
		i := matchPatterns(line)
		if i >= 0 {
			fmt.Println(
				fn + ":" +
				strconv.Itoa(lineNum) + ":" +
				strconv.Itoa(i+1) + ":", line)
		}
		lineNum++
	}

	return scanner.Err()
}

func main() {
	for _, pat := range excludeDirs {
		filenameFilter.excludeDir(pat)
	}
	for _, pat := range excludeFiles {
		filenameFilter.exclude(pat)
	}

	var filterFlag FilterFlag
	filterFlag.filenameFilter = filenameFilter

	flag.BoolVar(&ignoreCase, "i", false, "ignore case")
	flag.BoolVar(&oneLevel, "1", false, "just one level")
	flag.BoolVar(&verbose, "v", false, "verbose output")
	flag.Var(&filterFlag, ".", "filter by extension (e.g. -. txt)")
	flag.Parse()

	filepath.Walk(".", func(fn string, info os.FileInfo, err error) error {
		_, fnOnly := path.Split(fn)

		if info.IsDir() {
			if (oneLevel && fn != ".") ||
			   !filenameFilter.filterDir(fnOnly) {
				return filepath.SkipDir
			}
			return nil
		}

		if !filenameFilter.filterFilename(fnOnly) {
			return nil
		}

		fullFn := "./" + fn
		if verbose {
			fmt.Println("Searching in " + fullFn)
		}

		return searchInFileContent(fullFn)
	})
}
