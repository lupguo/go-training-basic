package sub

import (
	"reflect"
	"testing"
)

func TestSlice_Remove1(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name      string
		s         Slice
		args      args
		wantSlice *Slice
		wantErr   bool
	}{
		{
			name:      "remove head",
			s:         Slice{0, 1, 2, 3, 4, 5},
			args:      args{0},
			wantSlice: &Slice{1, 2, 3, 4, 5},
			wantErr:   false,
		},
		{
			name:      "remove middle",
			s:         Slice{0, 1, 2, 3, 4, 5},
			args:      args{2},
			wantSlice: &Slice{0, 1, 3, 4, 5},
			wantErr:   false,
		},
		{
			name:      "remove rear",
			s:         Slice{0, 1, 2, 3, 4, 5},
			args:      args{5},
			wantSlice: &Slice{0,1, 2, 3, 4},
			wantErr:   false,
		},
		{
			name:      "remove out range elem",
			s:         Slice{0, 1, 2, 3, 4, 5},
			args:      args{6},
			wantSlice: &Slice{0, 1, 2, 3, 4, 5},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSlice, err := tt.s.Remove(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSlice, tt.wantSlice) {
				t.Errorf("Remove() gotSlice = %v, want %v", gotSlice, tt.wantSlice)
			}
		})
	}
}
