package length_of_Last_Word

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test5letters(t *testing.T) {
	string1 := "Hello World"
	l := lengthOfLastWord(string1)
	want := 5

	require.Equal(t, want, l)
}

func Test4letters(t *testing.T) {
	string1 := "Hello Word "
	l := lengthOfLastWord(string1)
	want := 4

	require.Equal(t, want, l)
}
