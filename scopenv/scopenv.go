package scopenv

import (
	"fmt"
	"os"
	"strings"
)

// Scopenv - REDEFINE CREATE : Maybe also return error
func Scopenv(m []string, v ...string) map[string]string {

	r, _ := reverse(v)
	s := make(map[string]string)

	scope := find(s, m, r)

	return scope
}

func find(s map[string]string, m, rev []string) map[string]string {
	for i := range m {
		tmp := strings.Join(rev, "_")
		f := upper(tmp) + "_" + upper(m[i])
		if len(rev) == 0 {
			panic(fmt.Sprintf("Error: you are missing a env variable for one of the values in m."))
		}
		f = strings.Replace(f, "_", "", 1)

		check := os.Getenv(f)
		if check == "" {
			rev = removeFirst(rev)

			find(s, m, rev)
		} else {
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
		r = append(r, l) // Suggestion: do `last := len(s)-1` before the loop

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
