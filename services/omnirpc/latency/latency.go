package latency

import (
	"context"
	"errors"
	"fmt"
	"github.com/bradhe/stopwatch"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
	"net/url"
	"sync"
	"time"
)

// Result is the result of a latency check on a url.
type Result struct {
	// URL is the url of the latency being tested
	URL string
	// Latency is the latency time in seconds
	Latency time.Duration
	// HasError is wether or not the result has an error
	HasError bool
	// Error is the error recevied when trying to establish latency
	Error error
}

// GetRPCLatency gets latency from a list of rpcs.
func GetRPCLatency(parentCtx context.Context, timeout time.Duration, rpcList []string) (latSlice []Result) {
	var mux sync.Mutex

	timeCtx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()

	g, ctx := errgroup.WithContext(timeCtx)
	for _, rpcURL := range rpcList {
		// capture func literal
		rpcURL := rpcURL
		g.Go(func() error {
			latency := getLatency(ctx, rpcURL)

			mux.Lock()
			latSlice = append(latSlice, latency)
			mux.Unlock()

			return nil
		})
	}

	// we don't error at all above
	_ = g.Wait()
	return latSlice
}

func getLatency(ctx context.Context, rpcURL string) (l Result) {
	l = Result{URL: rpcURL, HasError: true}

	parsedURL, err := url.Parse(rpcURL)
	if err != nil {
		l.Error = fmt.Errorf("url invalid: %w", err)
		return l
	}

	// maybe we should allow this?
	if slices.Contains([]string{"ws", "wss"}, parsedURL.Scheme) {
		l.Error = errors.New("websockets not supported")
		return l
	}

	timer := stopwatch.Start()

	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		l.Error = fmt.Errorf("could not connect to %s: %w", rpcURL, err)
		return l
	}

	_, err = client.ChainID(ctx)
	if err != nil {
		l.Error = fmt.Errorf("could not get chainid: %w", err)
		return l
	}

	l.Latency = timer.Stop().Milliseconds()
	l.HasError = false
	return l
}