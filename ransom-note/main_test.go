package ransomnote

import (
	"reflect"
	"testing"
)

func TestRansomNode(t *testing.T) {
	in := [][2]string{{"a", "b"}, {"aa", "ab"}, {"aa", "aab"}}
	out := []bool{false, false, true}
	for i := range in {
		ransomNode, magazine := in[i][0], in[i][1]
		res := canConstruct(ransomNode, magazine)
		want := out[i]
		if reflect.DeepEqual(res, want) == false {
			t.Errorf("i = %v, result = %v, want %v", i, res, want)
		}
	}

}
