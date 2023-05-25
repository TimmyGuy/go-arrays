package arrays

import "reflect"

// ArrayMerge merges multiple arrays/slices into a single array/slice.
// It takes variable arguments representing the arrays to be merged and returns the merged array/slice.
func ArrayMerge(arrays ...[]interface{}) []interface{} {
	merged := make([]interface{}, 0)
	for _, arr := range arrays {
		merged = append(merged, arr...)
	}
	return merged
}

// IsArray checks whether the given value is an array.
// It takes a value as input and uses reflection to determine if the value is of kind "array".
// The function returns true if the value is an array, and false otherwise.
func IsArray(value interface{}) bool {
	valueType := reflect.TypeOf(value)
	return valueType.Kind() == reflect.Array
}
