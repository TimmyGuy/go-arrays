# Golang Array / Slice Functions


## Functions
### IndexOf
IndexOf returns the index of the first occurrence of the specified element in the array/slice.
If the element is not found, it returns -1.
```go
func IndexOf[Type comparable](array []Type, needle interface{}) int

// Example
func main() {
    needle := 3
    array := []int{1, 2, 3, 4, 3, 5}

    idx := arrays.IndexOf(array, needle)
    fmt.Println(idx) // 2
}
```

### LastIndexOf
LastIndexOf returns the index of the last occurrence of the specified element in the array/slice.
If the element is not found, it returns -1.
```go
func LastIndexOf[Type comparable](array []Type, needle interface{}) int

// Example
func main() {
    needle := 3
    array := []int{1, 2, 3, 4, 3, 5}

    idx := arrays.LastIndexOf(array, needle)
    fmt.Println(idx) // 4
}

```

### FindIndex
FindIndex returns the index of the first element in the array/slice that satisfies the given condition.
It takes the original array/slice and a callable function that accepts an item from the array/slice and returns a boolean value indicating whether the condition is met.
The function returns the index of the first matching element, or -1 if no element satisfies the condition.
```go
func FindIndex[Type any](array []Type, callable func(item Type) bool) int

// Example
func main() {
	array := []int{1, 2, 3, 4, 3, 5}

	idx := arrays.FindIndex(array, func(item int) bool {
		return item > 3
	})
	
	fmt.Println(idx) // 3
}
```

### Find
Find returns the first element in the array/slice that satisfies the given condition.
It takes the original array/slice and a callable function that accepts an item from the array/slice and returns a boolean value indicating whether the condition is met.
The function returns the first matching element, or nil if no element satisfies the condition.
```go
func Find[Type any](array []Type, callable func(item Type) bool) interface{}

// Example
func main() {
    array := []int{1, 2, 3, 4, 3, 5}

    item := arrays.Find(array, func(item int) bool {
        return item > 3
    })

    fmt.Println(item) // 4
}
```

### Contains
Contains checks whether the array/slice contains the specified element.
It returns true if the element is found, false otherwise.
```go
func Contains[Type comparable](array []Type, needle interface{}) bool

// Example
func main() {
	array := []int{1, 2, 3, 4, 3, 5}

	contains := arrays.Contains(array, 6)

	fmt.Println(contains) // false
}
```

### Flip
Flip reverses the order of elements in the array/slice and returns the modified array/slice.
```go
func Flip[Type any](array []Type) []Type

// Example
func main() {
	array := []int{1, 2, 3, 4, 3, 5}

	flippedArray := arrays.Flip(array)

	fmt.Println(flippedArray) // [5 3 4 3 2 1]
}
```

### Map
Map applies the given function to each element in the array/slice and returns a new array/slice with the results.
It takes the original array/slice and a callable function that accepts an item from the array/slice and returns a modified or transformed value.
The function returns a new array/slice containing the transformed values.
```go
func Map[Type any](array []Type, callable func(item Type) Type) []Type

// Example
func main() {
	array := []int{1, 2, 3, 4, 3, 5}

	flippedArray := arrays.Map(array, func(item int) int {
		return item + 1
	})

	fmt.Println(flippedArray) // [2 3 4 5 4 6]
}
```

### ForEach
ForEach applies the given function to each element in the array/slice without returning any results.
It takes the original array/slice and a callable function that accepts an item from the array/slice.
The function iterates over each element and applies the given function to it.
```go
func ForEach[Type any](array []Type, callable func(item Type))

// Example
func main() {
	array := []int{1, 2, 3, 4, 3, 5}

	arrays.ForEach(array, func(item int) {
		fmt.Print(item) // 123435
	})
}
```

### Slice
Slice returns a portion of the array/slice.
It takes the original array/slice, a size indicating the number of elements to include, and a start index.
The function returns a new array/slice containing elements from the original array/slice, starting from the given index and including the specified number of elements.
If the start index is beyond the array bounds, an empty slice is returned.
```go
func Slice[Type any](array []Type, size int, start int) []Type

// Example
func main() {
	array := []int{1, 2, 3, 4, 3, 5}

	slicedArray := arrays.Slice(array, 2, 1)

	fmt.Println(slicedArray) // [2 3]
}
```
*Also contains SliceFrom and SliceTo*

### Merge
Merge merges multiple slices into a single slice.
It takes variable arguments representing the slices to be merged and returns the merged slice.
```go
func Merge[Type any](slices ...[]Type) []Type

// Example
func main() {
    array1 := []int{1, 2, 3}
    array2 := []int{4, 3, 5}

    slicedArray := arrays.Merge(array1, array2)

    fmt.Println(slicedArray) // [1 2 3 4 3 5]
}
```

## Array Functions
### IsArray
Check if type is an array

```go
func IsArray(value interface{}) bool

// Example
func main() {
	array := [3]int{1, 2, 3}
	isArray := arrays.IsArray(array)
	fmt.Println(isArray) // True
}
```

## Slice Functions
### IsSlice
Check if type is an slice
```go
func IsSlice(value interface{}) bool

// Example
func main() {
    slice := []int{1, 2, 3}
    isSlice := arrays.IsSlice(array)
    fmt.Println(isSlice) // True
}
```