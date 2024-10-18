package join

import (
	"testing"

	"github.com/octohelm/go-observable/observable"
	observableutil "github.com/octohelm/go-observable/observable/util"
	testingx "github.com/octohelm/x/testing"
)

func TestCombineLatest2(t *testing.T) {
	src1 := observable.Of(0, 2)
	src0 := observable.Of(1, 3)

	combined := CombineLatest2(src0, src1, func(a int, b int) int {
		return a + b
	})

	values, err := observableutil.Collect(combined)
	testingx.Expect(t, err, testingx.BeNil[error]())
	testingx.Expect(t, values[len(values)-1], testingx.Be(2+3))
}
