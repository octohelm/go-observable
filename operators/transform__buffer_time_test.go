package operators

import (
	"fmt"
	"testing"
	"time"

	"github.com/octohelm/go-observable/observable"
	observableinterval "github.com/octohelm/go-observable/observable/interval"
	observableutil "github.com/octohelm/go-observable/observable/util"
	testingx "github.com/octohelm/x/testing"
)

func TestBufferTime(t *testing.T) {
	src := observableinterval.Interval(1 * time.Microsecond)

	buffered := observable.Pipe3(
		src,
		Count[time.Time](),
		BufferTime[int](50*time.Microsecond),
	)

	valueSeq, err := observableutil.Values(buffered)
	testingx.Expect(t, err, testingx.BeNil[error]())

	i := 0
	for values, _ := range valueSeq {
		fmt.Printf("buffered %d\n", len(values))

		i++
		if i == 10 {
			break
		}
	}
}
