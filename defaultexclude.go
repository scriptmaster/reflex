package main

import "regexp"

var defaultExcludes = []string{
	// Simple ignore all . dirs
	`(^|/)\..*`,
	// Vim
	`~$`,
}

var defaultExcludeMatcher multiMatcher

func init() {
	for _, pattern := range defaultExcludes {
		m := newRegexMatcher(regexp.MustCompile(pattern), true)
		defaultExcludeMatcher = append(defaultExcludeMatcher, m)
	}
}
