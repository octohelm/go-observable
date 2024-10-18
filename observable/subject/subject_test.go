package subject

import (
	"fmt"
	"github.com/octohelm/go-observable/observable"
	"slices"
	"sync"
	"testing"

	testingx "github.com/octohelm/x/testing"
)

func TestSubject(t *testing.T) {
	s := NewSubject[int]()

	chRet := make(chan string)

	wg := &sync.WaitGroup{}
	for i := range 3 {
		ob := s.Observe()

		wg.Add(1)
		go func() {
			defer wg.Done()

			runObserve(i+1, ob, chRet)
		}()
	}

	go func() {
		for i := range 10 {
			s.Send(i)
		}

		wg.Wait()
		s.CancelCause(nil)
		close(chRet)
	}()

	results := make([]string, 0)
	for ret := range chRet {
		results = append(results, ret)
	}
	slices.Sort(results)
	testingx.Expect(t, len(results), testingx.Be(2+3+4))

	fmt.Println(results)
}

func runObserve(id int, ob observable.Observer[int], recv chan<- string) {
	defer func() {
		ob.CancelCause(nil)
	}()

	count := 0
	for {
		x, _, ok := ob.Next()
		if !ok {
			return
		}
		recv <- fmt.Sprintf("%d-%d", id, x)
		count++
		if count > id {
			return
		}
	}
}
