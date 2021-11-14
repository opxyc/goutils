package sch

import (
	"context"
	"fmt"
	"time"
)

// PingAt takes a time hh-mm-ss. It sends an empty struct to the channel ch
// when system time = given time. Repeats until ctx is closed.
func PingAt(ctx context.Context, hh, mm, ss int, ch chan<- struct{}) {
	// calculate repeat interval
	t := time.Now()
	trgtTime := time.Date(t.Year(), t.Month(), t.Day(), hh, mm, 0, 0, t.Location())
	d := trgtTime.Sub(t)
	if d < 0 {
		trgtTime = trgtTime.Add(24 * time.Hour)
		d = trgtTime.Sub(t)
	}

	// and repeat
	go repeat(ctx, d, ch)
}

func repeat(ctx context.Context, d time.Duration, ch chan<- struct{}) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx closed")
			return
		case <-time.After(d):
			ch <- struct{}{}
		}
	}
}
