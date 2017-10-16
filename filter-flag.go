// Copyright (C) 2017 David Capello
//
// This file is released under the terms of the MIT license.
// Read LICENSE.txt for more information.

package main

type FilterFlag struct {
	filenameFilter *FilenameFilter
}

func (ff *FilterFlag) String() string {
	return ""
}

func (ff *FilterFlag) Set(value string) error {
	ff.filenameFilter.include("(?i).*\\." + value)
	return nil
}
