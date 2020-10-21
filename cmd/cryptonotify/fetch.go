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

// coinMap is an internal mutex protected type used for results accumulation
type coinMap struct {
	mu sync.Mutex
	cm map[int]coinlore.Coin
}

func (cm *coinMap) add(c coinlore.Coin) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.cm[c.ID] = c
}

func fetch(ctx context.Context, ids []int) (map[int]coinlore.Coin, error) {
	// Callers count
	n := int(math.Min(float64(len(ids)), maxSimultaneousCalls))

	wg := sync.WaitGroup{}
	wg.Add(n)

	jobs := make(chan int, n)
	res := coinMap{
		sync.Mutex{},
		make(map[int]coinlore.Coin),
	}

	cl := coinlore.NewClient(reqTimeout)

	for i := 0; i < n; i++ {
		go caller(ctx, cl, jobs, &wg, &res)
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
		return res.cm, nil
	}
}

func caller(ctx context.Context, cl coinlore.Client, calls <-chan int, wg *sync.WaitGroup, cm *coinMap) {
	defer wg.Done()

	for id := range calls {
		coin, err := cl.GetCoin(ctx, id)
		if err != nil {
			return
		}

		cm.add(coin)
	}
}
