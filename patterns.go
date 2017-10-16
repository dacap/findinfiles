// Copyright (C) 2017 David Capello
//
// This file is released under the terms of the MIT license.
// Read LICENSE.txt for more information.

package main

import (
	"log"
	"regexp"
)

type Patterns struct {
	patterns []*regexp.Regexp
}

func (p *Patterns) append(pat string) {
	re, err := regexp.Compile(pat)
	if err != nil {
		log.Fatal(err)
	}
	p.patterns = append(p.patterns, re)
}

func (p *Patterns) match(line string) int {
	firstIndex := -1
	matches := 0
	for _, re := range p.patterns {
		match := re.FindStringIndex(line)
		if match != nil {
			if (firstIndex < 0 ||
			    firstIndex > match[0]) {
				firstIndex = match[0]
			}
			matches++
		} else {
			return -1
		}
	}
	if matches == len(p.patterns) {
		return firstIndex
	}
	return -1
}
