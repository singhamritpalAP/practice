package solutions

import "testing"

func TestSortColors(t *testing.T) {
	type args struct {
		nums     []int
		approach int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Empty slice A1", args{[]int{}, 1}, []int{}},
		{"Empty slice A2", args{[]int{}, 2}, []int{}},
		{"Empty slice A1", args{[]int{0, 1, 0}, 1}, []int{0, 0, 1}},
		{"Empty slice A2", args{[]int{0, 1, 0}, 2}, []int{0, 0, 1}},
		{"Empty slice A1", args{[]int{0}, 1}, []int{0}},
		{"Empty slice A2", args{[]int{0}, 2}, []int{0}},
		{"Empty slice A1", args{[]int{0, 1}, 1}, []int{0, 1}},
		{"Empty slice A2", args{[]int{0, 1}, 2}, []int{0, 1}},
		{"Empty slice A1", args{[]int{2, 2, 0, 1, 0, 1, 2}, 1}, []int{0, 0, 1, 1, 2, 2, 2}},
		{"Empty slice A2", args{[]int{2, 2, 0, 1, 0, 1, 2}, 2}, []int{0, 0, 1, 1, 2, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortColors(tt.args.nums, tt.args.approach)
		})
	}
}
