package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		values mockList
		target int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		found bool
	}{
		{
			name: "Found at index 0 of a 4-element array",
			args: args{
				values: mockList{0, 10, 100, 1000},
				target: 0,
			},
			want:  0,
			found: true,
		},
		{
			name: "Found at index 1 of a 4-element array",
			args: args{
				values: mockList{0, 10, 100, 1000},
				target: 10,
			},
			want:  1,
			found: true,
		},
		{
			name: "Found at index 2 of a 4-element array",
			args: args{
				values: mockList{0, 10, 100, 1000},
				target: 100,
			},
			want:  2,
			found: true,
		},
		{
			name: "Found at index 3 of a 4-element array",
			args: args{
				values: mockList{0, 10, 100, 1000},
				target: 1000,
			},
			want:  3,
			found: true,
		},
		{
			name: "Found at index 0 of a 3-element array",
			args: args{
				values: mockList{0, 10, 100},
				target: 0,
			},
			want:  0,
			found: true,
		},
		{
			name: "Found at index 1 of a 3-element array",
			args: args{
				values: mockList{0, 10, 100},
				target: 10,
			},
			want:  1,
			found: true,
		},
		{
			name: "Found at index 2 of a 3-element array",
			args: args{
				values: mockList{0, 10, 100},
				target: 100,
			},
			want:  2,
			found: true,
		},
		{
			name: "Not found, it should be at index 0 of a 4-element array",
			args: args{
				values: mockList{0, 10, 100, 1000},
				target: -5,
			},
			want:  0,
			found: false,
		},
		{
			name: "Not found, it should be at index 1 of a 4-element array",
			args: args{
				values: mockList{0, 10, 100, 1000},
				target: 5,
			},
			want:  1,
			found: false,
		},
		{
			name: "Not found, it should be at index 2 of a 4-element array",
			args: args{
				values: mockList{0, 10, 100, 1000},
				target: 15,
			},
			want:  2,
			found: false,
		},
		{
			name: "Not found, it should be at index 3 of a 4-element array",
			args: args{
				values: mockList{0, 10, 100, 1000},
				target: 150,
			},
			want:  3,
			found: false,
		},
		{
			name: "Not found, it should be at index 4 of a 4-element array",
			args: args{
				values: mockList{0, 10, 100, 1000},
				target: 1500,
			},
			want:  4,
			found: false,
		},
		{
			name: "Not found, it should be at index 0 of a 3-element array",
			args: args{
				values: mockList{0, 10, 100},
				target: -5,
			},
			want:  0,
			found: false,
		},
		{
			name: "Not found, it should be at index 1 of a 3-element array",
			args: args{
				values: mockList{0, 10, 100},
				target: 5,
			},
			want:  1,
			found: false,
		},
		{
			name: "Not found, it should be at index 2 of a 3-element array",
			args: args{
				values: mockList{0, 10, 100},
				target: 15,
			},
			want:  2,
			found: false,
		},
		{
			name: "Not found, it should be at index 3 of a 3-element array",
			args: args{
				values: mockList{0, 10, 100},
				target: 150,
			},
			want:  3,
			found: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comparator := func(value interface{}) int {
				iv := value.(int)
				return iv - tt.args.target
			}

			got, found := BinarySearch(comparator, tt.args.values)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.found, found)
		})
	}
}
