// Package say provides functions to say meaningful words
package say

import "fmt"

// Hello accepts a string and returns a string
// which is a hello greeting.
func Hello(name string) string {
	return fmt.Sprintf("Hello %s.", name)
}
