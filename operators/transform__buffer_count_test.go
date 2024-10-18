package operators

import (
	"testing"

	"github.com/octohelm/go-observable/observable"
	observableutil "github.com/octohelm/go-observable/observable/util"
	testingx "github.com/octohelm/x/testing"
)

func TestBufferCount(t *testing.T) {
	src := observable.From(func(yield func(int, error) bool) {
		for i := range 100 {
			if !yield(i, nil) {
				return
			}
		}
	})

	buffered := observable.Pipe2(
		src,
		BufferCount[int](10),
	)

	valueSeq, err := observableutil.Values(buffered)
	testingx.Expect(t, err, testingx.BeNil[error]())

	i := 0
	for values, _ := range valueSeq {
		testingx.Expect(t, len(values), testingx.Be(10))
		i++
	}

	testingx.Expect(t, i, testingx.Be(10))
}
