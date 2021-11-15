package sch

import (
	"context"
	"time"
)

// PingAt takes a time hh-mm-ss. It sends an empty struct to the channel ch
// when system time = given time. Repeats until ctx is closed.
func PingAt(ctx context.Context, hh, mm, ss int, ch chan<- struct{}) {
	d := diff(hh, mm, ss)

	<-time.After(d)
	ch <- struct{}{}

	// and repeat daily
	d = 24 * time.Hour
	go repeat(ctx, d, ch)
}

// PingAfter sends an empty struct to ch after duration d until ctx is closed.
func PingAfter(ctx context.Context, d time.Duration, ch chan<- struct{}) {
	go repeat(ctx, d, ch)
}

// diff returns the number of seconds till hh-mm-ss
func diff(hh, mm, ss int) time.Duration {
	t := time.Now()
	trgtTime := time.Date(t.Year(), t.Month(), t.Day(), hh, mm, ss, 0, t.Location())
	d := trgtTime.Sub(t)
	if d < 0 {
		trgtTime = trgtTime.Add(24 * time.Hour)
		d = trgtTime.Sub(t)
	}
	return d
}

// repeat waits for duration d to send to channel ch, repetedly
func repeat(ctx context.Context, d time.Duration, ch chan<- struct{}) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(d):
			ch <- struct{}{}
		}
	}
}
