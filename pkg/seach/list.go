package search

// List contains the methods required to search for an element in a list
type List interface {

	// Length returns the number of elements in the list
	Length() int

	// ElementAt returns the element at the specified index
	ElementAt(index int) interface{}
}
