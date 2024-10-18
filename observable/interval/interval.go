package interval

import (
	"github.com/octohelm/go-observable/observable"
	"time"
)

func Interval(d time.Duration) observable.Observable[time.Time] {
	return observable.From(func(yield func(time.Time, error) bool) {
		ticker := time.NewTicker(d)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if !yield(time.Now(), nil) {
					return
				}
			}
		}
	})
}
