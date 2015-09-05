package models

import ()

type Activitylog struct {
	ID    string `binding:"required" json:"id" validate:"min=3,max=40"` // id
	Enter bool   `json:"enter"`                                         // enter
	Date  string `json:"date"`                                          // date
	Age   int    `json:"age" validate:"min=18"`
}

/*
max
	For numeric numbers, max will simply make sure that the value is
	equal to the parameter given. For strings, it checks that
	the string length is exactly that number of characters. For slices,
	arrays, and maps, validates the number of items. (Usage: len=10)

max
	For numeric numbers, max will simply make sure that the value is
	lesser or equal to the parameter given. For strings, it checks that
	the string length is at most that number of characters. For slices,
	arrays, and maps, validates the number of items. (Usage: max=10)

min
	For numeric numbers, min will simply make sure that the value is
	greater or equal to the parameter given. For strings, it checks that
	the string length is at least that number of characters. For slices,
	arrays, and maps, validates the number of items. (Usage: min=10)

nonzero
	This validates that the value is not zero. The appropriate zero value
	is given by the Go spec (e.g. for int it's 0, for string it's "", for
	pointers is nil, etc.) Usage: nonzero

regexp
	Only valid for string types, it will validate that the value matches
	the regular expression provided as parameter. (Usage: regexp=^a.*b$)
*/
