package leetcode

import "strings"

func simplifyPath(path string) string {
	res := []string{}

	tmp := strings.Split(path, "/")
	for _, d := range tmp {
		if d == "" || d == "." {
			continue
		}

		if d == ".." {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else {
			res = append(res, d)
		}
	}
	return "/" + strings.Join(res, "/")
}
