package scopenv

import (
	"fmt"
	"os"
	"strings"
)

// Scopenv - Main function of this package
func Scopenv(m []string, v ...string) map[string]string {

	r, _ := reverse(v)
	s := make(map[string]string)

	scope := find(s, m, r)

	return scope
}

// find - recurively finds and sets up env variables
func find(s map[string]string, m, rev []string) map[string]string {
	for i := range m {
		// Combine prefixes with variables and uppercase them
		tmp := strings.Join(rev, "_")
		f := upper(tmp) + "_" + upper(m[i])

		// If there are no more prefixes, remove the leading _
		if len(rev) == 0 {
			f = strings.Replace(f, "_", "", 1)
		}
		// Get your env vars
		check := os.Getenv(f)
		if check == "" {
			// Check for the length of the slice before removing
			if len(rev) == 0 {
				panic(fmt.Sprintf("Error: you are missing a env variable for one of the values in m."))
			}
			rev = removeFirst(rev)

			// Recursion !!!
			find(s, m, rev)
		} else {
			// Put the env vars in the map
			s[m[i]] = check
		}
	}

	return s
}

// reverse - reverses the order of a slice of strings
func reverse(v []string) ([]string, int) {
	var r []string
	for i := range v {
		l := v[len(v)-1-i]
		r = append(r, l)

	}
	leng := len(r)
	return r, leng
}

// upper - Helper function for capitalizing arguments
func upper(arg string) string {
	return strings.ToUpper(arg)
}

// removeFirst - removes the first item in a slice
func removeFirst(s []string) []string {
	s = append(s[:0], s[1:]...)
	return s
}
