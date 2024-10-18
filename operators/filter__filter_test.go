package operators

import (
	"testing"

	"github.com/octohelm/go-observable/observable"
	observableutil "github.com/octohelm/go-observable/observable/util"
	testingx "github.com/octohelm/x/testing"
)

func TestFilter(t *testing.T) {
	src := observable.From(func(yield func(int, error) bool) {
		for i := range 10 {
			if !yield(i, nil) {
				return
			}
		}
	})

	filtered := observable.Pipe(
		src,
		Filter(func(x int) bool {
			return x%2 == 0
		}),
	)

	values, err := observableutil.Collect(filtered)
	testingx.Expect(t, err, testingx.BeNil[error]())
	testingx.Expect(t, values, testingx.Equal([]int{
		0,
		2,
		4,
		6,
		8,
	}))
}
