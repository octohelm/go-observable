package interval

import (
	"fmt"
	"time"

	"github.com/octohelm/go-observable/observable"
	"github.com/robfig/cron/v3"
)

func FromCronSchedule(schedule cron.Schedule) observable.Observable[time.Time] {
	next := func() time.Duration {
		now := time.Now()
		return schedule.Next(now).Sub(now)
	}

	return observable.From(func(yield func(time.Time, error) bool) {
		timer := time.NewTimer(next())
		defer timer.Stop()

		for {
			select {
			case v := <-timer.C:
				if !yield(v, nil) {
					return
				}
			}

			timer.Reset(next())
		}
	})
}

func FromCron(rule string) (observable.Observable[time.Time], error) {
	schedule, err := cron.ParseStandard(rule)
	if err != nil {
		return nil, fmt.Errorf("parse cron failed: %s: %w", rule, err)
	}

	return FromCronSchedule(schedule), nil
}
