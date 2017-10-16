// Copyright (C) 2017 David Capello
//
// This file is released under the terms of the MIT license.
// Read LICENSE.txt for more information.

package main

var excludeDirs = []string {
  "(?i)\\.git",
  "(?i)\\.hg",
  "(?i)\\.svn",
}

var excludeFiles = []string {
  "#.*",
  "(?i).*\\.a",
  "(?i).*\\.bak",
  "(?i).*\\.bkp",
  "(?i).*\\.elc",
  "(?i).*\\.exe",
  "(?i).*\\.ilk",
  "(?i).*\\.lib",
  "(?i).*\\.o",
  "(?i).*\\.obj",
  "(?i).*\\.pch",
  "(?i).*\\.pdb",
  "(?i).*\\.pyc",
  ".*~",
  "BROWSE",
  "TAGS",
}
