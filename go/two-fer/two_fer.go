// Package twofer contains the ShareWith method
package twofer

import "fmt"

// ShareWith returns a string of the format "One for X, one for me.", where X is the input.
func ShareWith(someone string) string {
	if len(someone) == 0 {
		someone = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", someone)
}
