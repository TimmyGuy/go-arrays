package arrays

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	type args[Type comparable] struct {
		array  []Type
		needle interface{}
	}
	type testCase[Type comparable] struct {
		name string
		args args[Type]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "Does contain",
			args: args[int]{
				array:  []int{1, 2, 3, 4},
				needle: 2,
			},
			want: true,
		},
		{
			name: "Does not contain",
			args: args[int]{
				array:  []int{1, 2, 3, 4},
				needle: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.array, tt.args.needle); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFind(t *testing.T) {
	type args[Type any] struct {
		array    []Type
		callable func(item Type) bool
	}
	type testCase[Type any] struct {
		name string
		args args[Type]
		want interface{}
	}
	tests := []testCase[int]{
		{
			name: "Finds the right item",
			args: args[int]{
				array: []int{1, 2, 3, 4},
				callable: func(item int) bool {
					return item > 2
				},
			},
			want: 3,
		},
		{
			name: "Does not find an item",
			args: args[int]{
				array: []int{1, 2, 3, 4},
				callable: func(item int) bool {
					return item > 4
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.array, tt.args.callable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindIndex(t *testing.T) {
	type args[Type any] struct {
		array    []Type
		callable func(item Type) bool
	}
	type testCase[Type any] struct {
		name string
		args args[Type]
		want int
	}
	tests := []testCase[int]{
		{
			name: "Finds the right index",
			args: args[int]{
				array: []int{1, 2, 3, 4},
				callable: func(item int) bool {
					return item > 2
				},
			},
			want: 2,
		},
		{
			name: "Can not find the item so it returns -1",
			args: args[int]{
				array: []int{1, 2, 3, 4},
				callable: func(item int) bool {
					return item > 4
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindIndex(tt.args.array, tt.args.callable); got != tt.want {
				t.Errorf("FindIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlip(t *testing.T) {
	type args[Type any] struct {
		array []Type
	}
	type testCase[Type any] struct {
		name string
		args args[Type]
		want []Type
	}
	tests := []testCase[int]{
		{
			name: "Reverse the array",
			args: args[int]{
				array: []int{1, 2, 3, 4},
			},
			want: []int{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Flip(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	type args[Type any] struct {
		array    []Type
		callable func(item Type)
	}
	type testCase[Type any] struct {
		name string
		args args[Type]
	}
	tests := []testCase[int]{
		{
			name: "Run for each item the function",
			args: args[int]{
				array: []int{1, 2, 3, 4},
				callable: func(item int) {
					t.Log(item)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ForEach(tt.args.array, tt.args.callable)
		})
	}
}

func TestIndexOf(t *testing.T) {
	type args[Type comparable] struct {
		array  []Type
		needle interface{}
	}
	type testCase[Type comparable] struct {
		name string
		args args[Type]
		want int
	}
	tests := []testCase[int]{
		{
			name: "Get index of existing item",
			args: args[int]{
				array:  []int{1, 2, 3, 4},
				needle: 2,
			},
			want: 1,
		},
		{
			name: "Get index of non existing item",
			args: args[int]{
				array:  []int{1, 2, 3, 4},
				needle: 22,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOf(tt.args.array, tt.args.needle); got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsArray(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Item is array",
			args: args{
				value: [2]int{1, 2},
			},
			want: true,
		},
		{
			name: "Item is slice",
			args: args{
				value: []int{1, 2},
			},
			want: false,
		},
		{
			name: "Item is string",
			args: args{
				value: "hello",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsArray(tt.args.value); got != tt.want {
				t.Errorf("IsArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSlice(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Item is array",
			args: args{
				value: [2]int{1, 2},
			},
			want: false,
		},
		{
			name: "Item is slice",
			args: args{
				value: []int{1, 2},
			},
			want: true,
		},
		{
			name: "Item is string",
			args: args{
				value: "hello",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSlice(tt.args.value); got != tt.want {
				t.Errorf("IsSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastIndexOf(t *testing.T) {
	type args[Type comparable] struct {
		array  []Type
		needle interface{}
	}
	type testCase[Type comparable] struct {
		name string
		args args[Type]
		want int
	}
	tests := []testCase[int]{
		{
			name: "Find last index",
			args: args[int]{
				array:  []int{1, 2, 3, 4, 3, 5},
				needle: 3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastIndexOf(tt.args.array, tt.args.needle); got != tt.want {
				t.Errorf("LastIndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[Type any, R any] struct {
		array    []Type
		callable func(item Type) R
	}
	type testCase[Type any, R any] struct {
		name string
		args args[Type, R]
		want []R
	}
	tests := []testCase[int, int]{
		{
			name: "Add 1 to every number",
			args: args[int, int]{
				array: []int{1, 2, 3, 4},
				callable: func(item int) int {
					return item + 1
				},
			},
			want: []int{2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.array, tt.args.callable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args[Type any] struct {
		slices [][]Type
	}
	type testCase[Type any] struct {
		name string
		args args[Type]
		want []Type
	}
	tests := []testCase[int]{
		{
			name: "Add 2 arrays together",
			args: args[int]{
				slices: [][]int{{1, 2, 3}, {4, 5, 6}},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Add 3 arrays together",
			args: args[int]{
				slices: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.slices...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice(t *testing.T) {
	type args[Type any] struct {
		array []Type
		size  int
		start int
	}
	type testCase[Type any] struct {
		name string
		args args[Type]
		want []Type
	}
	tests := []testCase[int]{
		{
			name: "Slice middle",
			args: args[int]{
				array: []int{1, 2, 3, 4, 5, 6},
				size:  2,
				start: 2,
			},
			want: []int{3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice(tt.args.array, tt.args.size, tt.args.start); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceFrom(t *testing.T) {
	type args[Type any] struct {
		array []Type
		start int
	}
	type testCase[Type any] struct {
		name string
		args args[Type]
		want []Type
	}
	tests := []testCase[int]{
		{
			name: "Slice end",
			args: args[int]{
				array: []int{1, 2, 3, 4, 5, 6},
				start: 2,
			},
			want: []int{3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceFrom(tt.args.array, tt.args.start); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceTo(t *testing.T) {
	type args[Type any] struct {
		array []Type
		size  int
	}
	type testCase[Type any] struct {
		name string
		args args[Type]
		want []Type
	}
	tests := []testCase[int]{
		{
			name: "Slice start",
			args: args[int]{
				array: []int{1, 2, 3, 4, 5, 6},
				size:  3,
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceTo(tt.args.array, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
