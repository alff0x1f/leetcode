package String_Compression

import (
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	tests := []struct {
		byteArr          []byte
		compressedLength int
	}{
		{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, 6},
		{[]byte{'a'}, 1},
		{[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, 4},
		{[]byte{'a', 'b', 'c'}, 3},
	}
	for _, test := range tests {
		l := compress(test.byteArr)
		if l != test.compressedLength {
			strStr := fmt.Sprint(test.byteArr)
			t.Errorf("Compressed %s must be %d length, not %d", strStr, test.compressedLength, l)
		}
	}
}
