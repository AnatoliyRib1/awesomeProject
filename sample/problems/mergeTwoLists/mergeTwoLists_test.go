package mergeTwoLists

import (
	"github.com/stretchr/testify/require"
	. "sample/utils/linkedlist"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		name string
		got1 []int
		got2 []int
		want []int
	}{
		{
			name: "MergeEquals",
			got1: []int{1, 3, 3, 7, 9},
			got2: []int{1, 3, 3, 7, 9},
			want: []int{1, 1, 3, 3, 3, 3, 7, 7, 9, 9},
		},
		{
			name: "MergeEmpty",
			got1: []int{},
			got2: []int{2, 3},
			want: []int{2, 3},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got1ll := LinkedListFromSlice(test.got1)
			got2ll := LinkedListFromSlice(test.got2)

			var l = mergeTwoLists(got1ll, got2ll)
			lres := l.ToSlice()
			require.Equal(t, test.want, lres)
		})
	}
}
