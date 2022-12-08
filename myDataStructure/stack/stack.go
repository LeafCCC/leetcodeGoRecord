package stack

import (
	"errors"
	"strconv"
)

const LIMIT = 4

type Stack struct {
	index int
	data  [LIMIT]int
}

func (s *Stack) Push(n int) {
	if s.index >= LIMIT {
		return
	}
	s.data[s.index] = n
	s.index++
}

func (s *Stack) Pop() (int, error) {
	if s.index == 0 {
		return 0, errors.New("Empty")
	}
	s.index--
	return s.data[s.index], nil
}

func (s Stack) String() string {
	res := ""
	for i := 0; i < s.index; i++ {
		res += "[" + strconv.Itoa(i) + ":->" + strconv.Itoa(s.data[i]) + "] "
	}
	return res
}
