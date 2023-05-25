package arrays

import "reflect"

// IsArray checks whether the given value is an array.
// It takes a value as input and uses reflection to determine if the value is of kind "array".
// The function returns true if the value is an array, and false otherwise.
func IsArray(value interface{}) bool {
	valueType := reflect.TypeOf(value)
	return valueType.Kind() == reflect.Array
}
