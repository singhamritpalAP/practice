package solutions

import "testing"

func TestIsRotateSorted(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Empty slice", args{[]int{}}, true},                                      // An empty slice is considered sorted
		{"Single element", args{[]int{1}}, true},                                  // A single element is considered sorted
		{"Already sorted", args{[]int{1, 2, 3, 4, 5}}, true},                      // Already sorted array
		{"Rotated sorted", args{[]int{3, 4, 5, 1, 2}}, true},                      // Rotated sorted array
		{"Not rotated sorted", args{[]int{2, 3, 4, 5, 1}}, true},                  // Not a rotated sorted array
		{"Two elements sorted", args{[]int{1, 2}}, true},                          // Two elements in order
		{"Two elements not sorted", args{[]int{2, 1}}, true},                      // Two elements out of order
		{"Multiple rotations", args{[]int{4, 5, 6, 7, 0, 1, 2}}, true},            // Multiple rotations
		{"Multiple rotations not sorted", args{[]int{4, 5, 6, 7, 8, 0, 1}}, true}, // A rotated sorted array

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRotateSorted(tt.args.nums); got != tt.want {
				t.Errorf("IsRotateSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
