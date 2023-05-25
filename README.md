# Golang Array / Slice Functions


## Functions
### IndexOf
Returns the first index of the needle in the array
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
Returns the last index of the needle in the array
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
Find the index based on a lambda function
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
Find the item based on a lambda function
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
Check if array contains an item
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
Flips an array
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
Apply the given function to each element
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
Apply the given function, does not have a return value
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
Slice an array
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
Merge 2 or more slices together
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