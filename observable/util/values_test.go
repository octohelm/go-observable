package util

import (
	"testing"

	"github.com/octohelm/go-observable/observable"
	testingx "github.com/octohelm/x/testing"
)

func TestValues(t *testing.T) {
	o := observable.From(func(yield func(int, error) bool) {
		for i := range 10 {
			if !yield(i, nil) {
				return
			}
		}
	})

	t.Run("collect first value", func(t *testing.T) {
		first, err := FirstValue(o)
		testingx.Expect(t, err, testingx.BeNil[error]())
		testingx.Expect(t, first, testingx.Be(0))
	})

	t.Run("collect all values", func(t *testing.T) {
		values, err := Collect(o)
		testingx.Expect(t, err, testingx.BeNil[error]())
		testingx.Expect(t, len(values), testingx.Be(10))
	})
}
