package workerpool

import (
	"context"
	"sync"
)

type Result[Out any] struct {
	Index int
	Value Out
	Err   error
}

func Map[In, Out any](ctx context.Context, inputs []In, workers int, fn func(context.Context, In) (Out, error)) []Result[Out] {
	if workers < 1 {
		workers = 1
	}

	type job struct {
		idx int
		in  In
	}

	jobs := make(chan job)
	results := make([]Result[Out], len(inputs))

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				v, err := fn(ctx, j.in)
				results[j.idx] = Result[Out]{Index: j.idx, Value: v, Err: err}
			}
		}()
	}
feed:
	for i, in := range inputs {
		select {
		case jobs <- job{idx: i, in: in}:
		case <-ctx.Done():
			for k := i; k < len(inputs); k++ {
				results[k] = Result[Out]{Index: k, Err: ctx.Err()}
			}
			break feed
		}
	}
	close(jobs)
	wg.Wait()
	return results
}
