package reverse_Linked_List

import (
	"github.com/stretchr/testify/require"
	. "sample/utils/linkedlist"
	"testing"
)

func Test_reverseLL(t *testing.T) {
	tests := []struct {
		name string
		got  []int
		want []int
	}{
		{
			name: "Zero",
			got:  []int{},
			want: []int{},
		},
		{
			name: "FiveElements",
			got:  []int{1, 15, 17, 98, 90},
			want: []int{90, 98, 17, 15, 1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotll := LinkedListFromSlice(test.got)

			var l = reverseList(gotll)
			lres := l.ToSlice()
			require.Equal(t, test.want, lres)
		})
	}
}
