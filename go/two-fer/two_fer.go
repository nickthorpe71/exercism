// Package twofer implements a function that takes an optional name and returns a string containing either the name or a replacement word.
package twofer

// ShareWith - takes an optional name string and returns a sentance with the provided name or a replacement word
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
