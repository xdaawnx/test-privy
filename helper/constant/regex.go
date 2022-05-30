package constant

import "regexp"

// AlphUnds is a regex only alphanumeric and underscore and dot
var AlphUnds = regexp.MustCompile("^[a-zA-Z0-9_.:-]*$")

// Alph is a regex only alphanumeric
var Alph = regexp.MustCompile("^[a-zA-Z0-9]*$")

// AlphSimple is a regex only alphanumeric and -')(,.
var AlphSimple = regexp.MustCompile("^[a-zA-Z0-9 -')(,:?.]*$")

// Numb is a regex only number
var Numb = regexp.MustCompile("^[0-9]*$")

// NumDate datetime is a regex only NumDate
var NumDate = regexp.MustCompile("^[0-9 :-]*$")

// PhoneNumb is a regex only Phone Number
var PhoneNumb = regexp.MustCompile("^[+]*[0-9]*$")

// NoSpace is a regex no Space
var NoSpace = regexp.MustCompile(`^[a-zA-Z0-9_]*$`)
