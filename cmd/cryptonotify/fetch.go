package main

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/Sewiti/crypto-notify/pkg/coinlore"
)

const (
	maxSimultaneousCalls = 4
	reqTimeout           = 30 * time.Second
)

// monitor is an internal mutex protected type used for results accumulation.
type monitor struct {
	mu   sync.Mutex            // Mutex that protects other fields.
	cm   map[int]coinlore.Coin // Coins map where results are stored in a convenient manner.
	errs []error               // Errors slice to catch any errors that occured in goroutines.
}

// add appends a coin to a map.
func (m *monitor) add(c coinlore.Coin) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cm[c.ID] = c
}

// err appends an error to catch it later. Goroutine should return after this.
func (m *monitor) err(err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.errs = append(m.errs, err)
}

// fetch calls API requests simultaneously the given coin ids.
func fetch(ctx context.Context, ids []int) (map[int]coinlore.Coin, error) {
	n := int(math.Min(float64(len(ids)), maxSimultaneousCalls)) // callers count

	wg := sync.WaitGroup{}
	wg.Add(n)

	jobs := make(chan int, n)
	mon := monitor{
		mu: sync.Mutex{},
		cm: make(map[int]coinlore.Coin),
	}

	cl := coinlore.NewClient(reqTimeout)

	for i := 0; i < n; i++ {
		go caller(ctx, cl, jobs, &wg, &mon)
	}

	for _, v := range ids {
		jobs <- v
	}
	close(jobs)

	wg.Wait()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	default:
		for _, v := range mon.errs {
			return nil, v
		}

		return mon.cm, nil
	}
}

// caller is a goroutine function that calls API requests on demand.
func caller(ctx context.Context, cl coinlore.Client, calls <-chan int, wg *sync.WaitGroup, mon *monitor) {
	defer wg.Done()

	for id := range calls {
		coin, err := cl.GetCoin(ctx, id)
		if err != nil {
			mon.err(err)
			return
		}

		mon.add(coin)
	}
}
