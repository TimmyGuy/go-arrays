package arrays

import "reflect"

// SliceMerge merges multiple slices into a single slice.
// It takes variable arguments representing the slices to be merged and returns the merged slice.
func SliceMerge(slices ...[]interface{}) []interface{} {
	var merged []interface{}
	for _, slice := range slices {
		merged = append(merged, slice...)
	}
	return merged
}

// IsSlice checks whether the given value is a slice.
// It takes a value as input and uses reflection to determine if the value is of kind "slice".
// The function returns true if the value is a slice, and false otherwise.
func IsSlice(value interface{}) bool {
	valueType := reflect.TypeOf(value)
	return valueType.Kind() == reflect.Slice
}
