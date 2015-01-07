package main

import "fmt"

func main() {
	l := Scopenv([]string{"host", "token"}, "core")
	fmt.Println(l)
}

// // Scopenv - REDEFINE CREATE : Maybe also return error
// func Scopenv(m []string, v ...string) map[string]string {
//
// 	r := reverse(v)
// 	s := make(map[string]string)
//
// 	scope := find(s, m, r)
//
// 	return scope
// }
//
// func find(s map[string]string, m, rev []string) map[string]string {
//
// 	for i := range m {
// 		tmp := strings.Join(rev, "_")
// 		f := upper(tmp) + "_" + upper(m[i])
// 		check := os.Getenv(f)
// 		if check == "" {
// 			rev = removeFirst(rev)
// 			find(s, m, rev)
// 		} else {
// 			s[m[i]] = check
// 		}
// 	}
//
// 	return s
// }
//
// // reverse - reverses the order of a slice of strings
// func reverse(v []string) []string {
// 	var r []string
// 	for i := range v {
// 		l := v[len(v)-1-i]
// 		r = append(r, l) // Suggestion: do `last := len(s)-1` before the loop
// 	}
// 	return r
// }
//
// // upper - Helper function for capitalizing arguments
// func upper(arg string) string {
// 	return strings.ToUpper(arg)
// }
//
// // removeFirst - removes the first item in a slice
// func removeFirst(s []string) []string {
// 	s = append(s[:0], s[1:]...)
// 	return s
// }
