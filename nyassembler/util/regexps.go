package util

import "regexp"

var CommentsRegex = regexp.MustCompile(`;.*`)
var LabelRegex = regexp.MustCompile(`(^[A-z-_]*)(?::)`)
