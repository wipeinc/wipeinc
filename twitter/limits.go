package twitter

import (
	"net/http"
	"strconv"
	"time"
)

// Limit sructure represent rate limiting header
type Limit struct {
	Remaining int
	Limit     int
	Reset     int64
}

// TimeLeft return time left on current rate limiting window
func (l *Limit) TimeLeft() time.Duration {
	reset := time.Unix(l.Reset, 0)
	return reset.Sub(time.Now())
}

// Delay try to delay rate  limiting
func (l *Limit) Delay() time.Duration {
	return time.Duration(float64(l.TimeLeft()) / (float64(l.Remaining) + 1))
}

// GetLimits extract limits from the headers
func GetLimits(resp *http.Response) (*Limit, error) {
	remainingStr := resp.Header.Get("x-rate-limit-remaining")
	remaining, err := strconv.Atoi(remainingStr)
	if err != nil {
		return nil, err
	}

	limitStr := resp.Header.Get("x-rate-limit-limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, err
	}

	resetStr := resp.Header.Get("x-rate-limit-reset")
	reset, err := strconv.ParseInt(resetStr, 10, 64)
	if err != nil {
		return nil, err
	}

	return &Limit{
		Limit:     limit,
		Remaining: remaining,
		Reset:     reset,
	}, nil
}
