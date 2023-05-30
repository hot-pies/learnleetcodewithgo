package datastructure

type Stack []string

func (s *Stack) Push(element string) {
	*s = append(*s, element)
}

func (s *Stack) Pop() string {

	if s.IsEmpty() {
		return ""
	}

	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
