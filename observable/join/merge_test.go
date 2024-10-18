package join

import (
	"slices"
	"testing"

	"github.com/octohelm/go-observable/observable"
	observableutil "github.com/octohelm/go-observable/observable/util"
	testingx "github.com/octohelm/x/testing"
)

func TestMerge(t *testing.T) {
	src1 := observable.Of(0, 2, 4)
	src0 := observable.Of(1, 3, 5)

	merged := Merge(src0, src1)

	values, err := observableutil.Collect(merged)
	testingx.Expect(t, err, testingx.BeNil[error]())

	slices.Sort(values)

	testingx.Expect(t, values, testingx.Equal([]int{
		0, 1, 2, 3, 4, 5,
	}))
}
