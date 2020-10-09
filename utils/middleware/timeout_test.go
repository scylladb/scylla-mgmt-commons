// Copyright (C) 2017 ScyllaDB

package middleware_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/scylladb/scylla-mgmt-commons/utils/middleware"
	"github.com/scylladb/scylla-mgmt-commons/utils/middleware/scyllaclienttest"
)

func TestCustomTimeout(t *testing.T) {
	td := []struct {
		Name            string
		Context         context.Context
		Timeout         time.Duration
		ExpectedTimeout time.Duration
	}{
		{
			Name:            "timeout in context",
			Context:         middleware.CustomTimeout(context.Background(), 5*time.Millisecond),
			Timeout:         0,
			ExpectedTimeout: 5 * time.Millisecond,
		},
		{
			Name:            "timeout in parameter",
			Context:         context.Background(),
			Timeout:         5 * time.Millisecond,
			ExpectedTimeout: 5 * time.Millisecond,
		},
		{
			Name:            "both timeouts, choose min",
			Context:         middleware.CustomTimeout(context.Background(), 5*time.Millisecond),
			Timeout:         10 * time.Millisecond,
			ExpectedTimeout: 5 * time.Millisecond,
		},
		{
			Name:            "both timeouts, choose min",
			Context:         middleware.CustomTimeout(context.Background(), 10*time.Millisecond),
			Timeout:         5 * time.Millisecond,
			ExpectedTimeout: 5 * time.Millisecond,
		},
	}

	for i := range td {
		test := td[i]
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			done := make(chan struct{})

			hang := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				<-done
			})

			host, port, closeServer := scyllaclienttest.MakeServer(t, hang)
			defer closeServer()
			client := scyllaclienttest.MakeClient(t, host, port, func(config *scyllaclienttest.Config) {
				config.Timeout = test.Timeout
			})
			_, err := client.NodeInfo(test.Context, host)
			close(done)

			var golden = fmt.Sprintf("after %s: timeout", test.ExpectedTimeout)
			if err == nil || err.Error() != golden {
				t.Fatalf("Ping() error = %v, expected %s", err, golden)
			}
		})
	}

}
