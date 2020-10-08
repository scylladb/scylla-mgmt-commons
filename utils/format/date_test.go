// Copyright (C) 2017 ScyllaDB

package format

import (
	"strings"
	"testing"
	"time"

	"github.com/scylladb/scylla-mgmt-commons/utils/timeutc"
)

func TestParseStartDate(t *testing.T) {
	t.Parallel()

	const epsilon = 50 * time.Millisecond

	table := []struct {
		S string
		D time.Duration
		E string
	}{
		{
			S: "now",
			D: nowSafety,
		},
		{
			S: "now-5s",
			E: "start date cannot be in the past"},
		{
			S: "now+5s",
			E: "start date must be at least in",
		},
		{
			S: "now+1h",
			D: time.Hour,
		},
		{
			S: timeutc.Now().Add(-5 * time.Second).Format(time.RFC3339),
			E: "start date cannot be in the past"},
		{
			S: timeutc.Now().Add(5 * time.Second).Format(time.RFC3339),
			E: "start date must be at least in",
		},
		{
			S: timeutc.Now().Add(time.Hour).Format(time.RFC3339),
			D: time.Hour,
		},
		{
			S: "2019-05-02T15:04:05Z07:00",
			E: "extra text",
		},
	}

	for i, test := range table {
		startDate, err := ParseStartDate(test.S)

		msg := ""
		if err != nil {
			msg = err.Error()
		}
		if test.E != "" || msg != "" {
			if !strings.Contains(msg, test.E) {
				t.Error(i, msg)
			}
			if msg != "" && test.E == "" {
				t.Error(i, msg)
			}
			continue
		}

		s := truncateToSecond(time.Time(startDate))
		now := truncateToSecond(timeutc.Now())
		diff := now.Add(test.D).Sub(s)
		if diff < 0 {
			diff *= -1
		}
		if diff > epsilon {
			t.Fatal(i, startDate, test.D, diff, test.S)
		}
	}
}

func truncateToSecond(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
}
