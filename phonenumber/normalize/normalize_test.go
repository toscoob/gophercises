package normalize

import (
	"fmt"
	"testing"
)

func TestNormalize(t *testing.T){
	var testData = []struct {
		in string
		want string
	}{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"123-456-7890", "1234567890"},
		{"1234567892", "1234567892"},
		{"(123)456-7892", "1234567892"},
		{"asfgghfjgj1", ""},
	}

	for i, td := range testData {
		testname := fmt.Sprintf("%d: %s", i, td.in)
		t.Run(testname, func(t *testing.T) {
			ans, _ := Normalize(td.in)
			if ans != td.want {
				t.Errorf("got %s, want %s", ans, td.want)
			}
		})
	}
}
