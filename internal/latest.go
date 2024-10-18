package internal

import (
	"sync/atomic"
)

type Latest[T any] struct {
	v atomic.Pointer[T]
}

func (p *Latest[T]) Watch(ob Observer[T], changes chan<- struct{}) error {
	for {
		v, err, ok := ob.Next()
		if !ok {
			break
		}
		if err != nil {
			return err
		}
		p.v.Store(&v)
		changes <- struct{}{}
	}

	return nil
}

func (p *Latest[T]) Value() (v *T) {
	return p.v.Load()
}
