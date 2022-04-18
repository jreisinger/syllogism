package syllogism

import (
	"reflect"
	"testing"
)

func TestSyllogism(t *testing.T) {
	testcases := []struct {
		major      Proposition
		minor      Proposition
		conclusion Proposition
	}{
		{
			Proposition{[]string{"rational animal"}, []string{"mortal"}},
			Proposition{[]string{"socrates"}, []string{"rational animal"}},
			Proposition{[]string{"socrates"}, []string{"mortal"}},
		},
	}
	for _, tc := range testcases {
		conclusion, err := Conclude(tc.major, tc.minor)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(conclusion, tc.conclusion) {
			t.Fatalf("\nGot: %+v\nExp: %+v", conclusion, tc.conclusion)
		}
		t.Logf("\nmajor: %v\nminor: %v\nconcl: %v", tc.major, tc.minor, conclusion)
	}
}
