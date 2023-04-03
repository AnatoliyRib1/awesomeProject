package rotateArray

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test3Elements(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	l := rotate(slice, 3)

	require.Equal(t, []int{3, 4, 5, 1, 2}, l)
}

func Test5Elements(t *testing.T) {
	slice := []int{1, 2, 3}
	l := rotate(slice, 5)

	require.Equal(t, []int{2, 3, 1}, l)
}

func TestAnyElements(t *testing.T) {
	tests := []struct {
		name string
		got  []int
		want []int
		num  int
	}{
		{
			name: "FourElements",
			got:  []int{1, 3, 5, 7, 9},
			want: []int{3, 5, 7, 9, 1},
			num:  4,
		},
		{
			name: "SixElements",
			got:  []int{5, -7, 156, -1},
			want: []int{156, -1, 5, -7},
			num:  6,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := rotate(test.got, test.num)
			require.Equal(t, test.want, l)

		})

	}
}
