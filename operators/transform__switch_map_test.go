package operators

import (
	"testing"
	"time"

	"github.com/octohelm/go-observable/observable"
	observableutil "github.com/octohelm/go-observable/observable/util"
	testingx "github.com/octohelm/x/testing"
)

func TestSwitchMap(t *testing.T) {
	src := observable.Of(0, 1, 2, 3, 4)

	mapped := SwitchMap(src, func(x int) observable.Observable[int] {
		if x%2 == 0 {
			return observable.Of(x * 2)
		}
		return observable.Of(x * x)
	})

	values, err := observableutil.Collect(mapped)
	testingx.Expect(t, err, testingx.BeNil[error]())
	testingx.Expect(t, values, testingx.Equal([]int{
		0 * 2,
		1 * 1,
		2 * 2,
		3 * 3,
		4 * 2,
	}))

	time.Sleep(1 * time.Second)
}
