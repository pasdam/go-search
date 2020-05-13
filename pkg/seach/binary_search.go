package search

// BinarySearch finds the position of a target value within a sorted array. It
// returns the position the matching element, or the position at which it should
// be. The returned boolean flag indicates whether the element was found or not.
// comparator is a function that take an element of the array as input and
// returns a value less than 0 if its input is less than the target, a value
// greater than 0 if it's greater and 0 if it matches the target.
func BinarySearch(comparator func(interface{}) int, values List) (int, bool) {
	left := 0
	right := values.Length() - 1
	for left <= right {
		current := (left + right) / 2
		compare := comparator(values.ElementAt(current))
		if compare < 0 {
			left = current + 1

		} else if compare > 0 {
			right = current - 1

		} else {
			return current, true
		}
	}

	return left, false
}
