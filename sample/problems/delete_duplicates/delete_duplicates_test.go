package delete_duplicates

import (
	"github.com/stretchr/testify/require"
	. "sample/utils/linkedlist"
	"testing"
)

func Test_Delete_duplicates(t *testing.T) {
	tests := []struct {
		name string
		got  []int
		want []int
	}{
		{
			name: "Delete_3",
			got:  []int{1, 3, 3, 7, 9},
			want: []int{1, 3, 7, 9},
		},
		{
			name: "Delete_4_6",
			got:  []int{2, 3, 4, 4, 4, 5, 6, 6, 6, 6, 7, 8},
			want: []int{2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotll := LinkedListFromSlice(test.got)

			var l = deleteDuplicates(gotll)
			lres := l.ToSlice()
			require.Equal(t, test.want, lres)

		})

	}
}
