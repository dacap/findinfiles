// Copyright (C) 2017 David Capello
//
// This file is released under the terms of the MIT license.
// Read LICENSE.txt for more information.

package main

import (
	"log"
	"regexp"
)

type FilenameFilter struct {
	dirExcludes, excludes, includes []*regexp.Regexp
}

func (ff *FilenameFilter) excludeDir(pat string) {
	re, err := regexp.Compile(pat)
	if err != nil {
		log.Fatal(err)
	}
	ff.dirExcludes = append(ff.dirExcludes, re);
}

func (ff *FilenameFilter) exclude(pat string) {
	re, err := regexp.Compile(pat)
	if err != nil {
		log.Fatal(err)
	}
	ff.excludes = append(ff.excludes, re);
}

func (ff *FilenameFilter) include(pat string) {
	re, err := regexp.Compile(pat)
	if err != nil {
		log.Fatal(err)
	}
	ff.includes = append(ff.includes, re);
}

func (ff *FilenameFilter) filterDir(dn string) bool {
	for _, re := range ff.dirExcludes {
		if re.MatchString(dn) { return false }
	}
	return true
}

func (ff *FilenameFilter) filterFilename(fn string) bool {
	for _, re := range ff.excludes {
		if re.MatchString(fn) { return false }
	}
	if len(ff.includes) == 0 { return true }
	for _, re := range ff.includes {
		if re.MatchString(fn) { return true }
	}
	return false
}
