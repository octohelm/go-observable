package interval

import (
	"fmt"
	"time"

	"github.com/octohelm/go-observable/observable"
	"github.com/robfig/cron/v3"
)

type CronSpec string

func (spec CronSpec) Schedule() cron.Schedule {
	s, _ := cron.ParseStandard(string(spec))
	return s
}

func (spec *CronSpec) UnmarshalText(text []byte) error {
	s := string(text)

	switch s {
	case "@never":
		return nil
	default:
		_, err := cron.ParseStandard(s)
		if err != nil {
			return fmt.Errorf("invalid cron spec: %s: %w", s, err)
		}
		*spec = CronSpec(s)
	}

	return nil
}

func FromCronSchedule(schedule cron.Schedule) observable.Observable[time.Time] {
	if schedule == nil {
		return observable.Empty[time.Time]()
	}

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
