package leetcode

import "sort"

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	//注意字典序中 '/' 是小于英文字母的

	res := []string{folder[0]}

	master := folder[0]

	for i := 1; i < len(folder); i++ {
		//判断如果不是子文件夹 则选择一个新的主目录
		child := folder[i]
		if len(child) <= len(master) || child[:len(master)] != master || child[len(master)] != '/' {
			master = child
			res = append(res, child)
		}
	}

	return res
}
