package operators

import (
	"testing"

	"github.com/octohelm/go-observable/observable"
	observableutil "github.com/octohelm/go-observable/observable/util"
	testingx "github.com/octohelm/x/testing"
)

func TestMap(t *testing.T) {
	src := observable.From(func(yield func(int, error) bool) {
		for i := range 5 {
			if !yield(i, nil) {
				return
			}
		}
	})

	mapped := observable.Pipe2(
		src,
		Map(func(x int) bool {
			return x%2 == 0
		}),
	)

	values, err := observableutil.Collect(mapped)
	testingx.Expect(t, err, testingx.BeNil[error]())
	testingx.Expect(t, values, testingx.Equal([]bool{
		true,
		false,
		true,
		false,
		true,
	}))
}
