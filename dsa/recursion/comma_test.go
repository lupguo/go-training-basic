package recursion

import "testing"

func TestComma(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1",
			args{"1234567"},
			"1,234,567",
		},
		{"t2",
			args{"123"},
			"123",
		},
		{"t3",
			args{"1234"},
			"1,234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Comma(tt.args.s); got != tt.want {
				t.Errorf("Comma() = %v, want %v", got, tt.want)
			}
		})
	}
}
