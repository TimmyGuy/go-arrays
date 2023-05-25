package arrays

// IndexOf returns the index of the first occurrence of the specified element in the array/slice.
// If the element is not found, it returns -1.
func IndexOf[Type comparable](array []Type, needle interface{}) int {
	for idx, item := range array {
		if item == needle {
			return idx
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of the specified element in the array/slice.
// If the element is not found, it returns -1.
func LastIndexOf[Type comparable](array []Type, needle interface{}) int {
	for idx := len(array) - 1; idx >= 0; idx-- {
		if array[idx] == needle {
			return idx
		}
	}
	return -1
}

// FindIndex returns the index of the first element in the array/slice that satisfies the given condition.
// It takes the original array/slice and a callable function that accepts an item from the array/slice and returns a boolean value indicating whether the condition is met.
// The function returns the index of the first matching element, or -1 if no element satisfies the condition.
func FindIndex[Type any](array []Type, callable func(item Type) bool) int {
	for idx, item := range array {
		if callable(item) {
			return idx
		}
	}
	return -1
}

// Find returns the first element in the array/slice that satisfies the given condition.
// It takes the original array/slice and a callable function that accepts an item from the array/slice and returns a boolean value indicating whether the condition is met.
// The function returns the first matching element, or nil if no element satisfies the condition.
func Find[Type any](array []Type, callable func(item Type) bool) interface{} {
	for _, item := range array {
		if callable(item) {
			return item
		}
	}
	return nil
}

// Contains checks whether the array/slice contains the specified element.
// It returns true if the element is found, false otherwise.
func Contains[Type comparable](array []Type, needle interface{}) bool {
	return IndexOf(array, needle) != -1
}

// Flip reverses the order of elements in the array/slice and returns the modified array/slice.
func Flip[Type any](array []Type) []Type {
	length := len(array)
	for i := 0; i < length/2; i++ {
		array[i], array[length-1-i] = array[length-1-i], array[i]
	}
	return array
}

// Map applies the given function to each element in the array/slice and returns a new array/slice with the results.
// It takes the original array/slice and a callable function that accepts an item from the array/slice and returns a modified or transformed value.
// The function returns a new array/slice containing the transformed values.
func Map[Type any](array []Type, callable func(item Type) Type) []Type {
	var res []Type
	for _, item := range array {
		res = append(res, callable(item))
	}
	return res
}

// ForEach applies the given function to each element in the array/slice without returning any results.
// It takes the original array/slice and a callable function that accepts an item from the array/slice.
// The function iterates over each element and applies the given function to it.
func ForEach[Type any](array []Type, callable func(item Type)) {
	for _, item := range array {
		callable(item)
	}
}

// Slice returns a portion of the array/slice.
// It takes the original array/slice, a size indicating the number of elements to include, and a start index.
// The function returns a new array/slice containing elements from the original array/slice, starting from the given index and including the specified number of elements.
// If the start index is beyond the array bounds, an empty slice is returned.
func Slice[Type any](array []Type, size int, start int) []Type {
	if start >= len(array) {
		return []Type{}
	}

	end := start + size
	if end > len(array) {
		end = len(array)
	}

	return array[start:end]
}

// SliceTo returns an array/slice from index 0 to the specified size.
// It is a convenience function that calls the Slice function with a start index of 0.
func SliceTo[Type any](array []Type, size int) []Type {
	return Slice(array, size, 0)
}

// SliceFrom returns an array/slice starting from the specified index until the end.
// It is a convenience function that calls the Slice function with a size of 0, indicating all elements from the start index to the end of the array/slice.
func SliceFrom[Type any](array []Type, start int) []Type {
	return Slice(array, 0, start)
}

// Merge merges multiple slices into a single slice.
// It takes variable arguments representing the slices to be merged and returns the merged slice.
func Merge[Type any](slices ...[]Type) []Type {
	var merged []Type
	for _, slice := range slices {
		merged = append(merged, slice...)
	}
	return merged
}
