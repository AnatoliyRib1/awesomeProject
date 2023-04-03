package merge_sorted_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMergeArray(t *testing.T) {
	tests := []struct {
		name string
		got1 []int
		num1 int
		got2 []int
		num2 int
		want []int
	}{
		{
			name: "eigthElements",
			got1: []int{1, 3, 5, 7, 9},
			got2: []int{1, 3, 5, 7, 9},
			num1: 4,
			num2: 4,
			want: []int{1, 1, 3, 3, 5, 5, 7, 7},
		},
		{
			name: "oneElements",
			got1: []int{2, 5},
			got2: []int{},
			num1: 1,
			num2: 0,
			want: []int{2},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := merge(test.got1, test.num1, test.got2, test.num2)
			require.Equal(t, test.want, l)

		})

	}
}
