package arrays

import "reflect"

// IsSlice checks whether the given value is a slice.
// It takes a value as input and uses reflection to determine if the value is of kind "slice".
// The function returns true if the value is a slice, and false otherwise.
func IsSlice(value interface{}) bool {
	valueType := reflect.TypeOf(value)
	return valueType.Kind() == reflect.Slice
}
