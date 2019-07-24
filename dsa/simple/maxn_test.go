package simple

import "testing"

func TestMax(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"t1",
			args{
				[]int{3, 2, -1, 5, 7, 9, 3},
			},
			9,
		},
		{
			"t2",
			args{
				[]int{100, 0733, 0x32, 27, 'a', 'z', 'y'},
			},
			0733,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.list); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
