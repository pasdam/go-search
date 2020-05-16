package search

type mockList []int

func (l mockList) Length() int {
	return len(l)
}

func (l mockList) ElementAt(index int) interface{} {
	return l[index]
}
