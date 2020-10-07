// Copyright (C) 2017 ScyllaDB

package middleware

import (
	"context"
	"net"
	"net/http"

	"github.com/hailocab/go-hostpool"
	"github.com/pkg/errors"
	"github.com/scylladb/scylla-mgmt-commons/httpx"
)

var errPoolServerError = errors.New("server error")

// HostPool sets request host from a pool.
func HostPool(next http.RoundTripper, pool hostpool.HostPool, port string) http.RoundTripper {
	return httpx.RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
		ctx := req.Context()

		var (
			h   string
			hpr hostpool.HostPoolResponse
		)

		// Get host from context
		h, ok := ctx.Value(ctxHost).(string)

		// Get host from pool
		if !ok {
			hpr = pool.Get()
			h = hpr.Host()
		}

		// Clone request
		r := httpx.CloneRequest(req)

		// Set host and port
		hp := net.JoinHostPort(h, port)
		r.Host = hp
		r.URL.Host = hp

		// RoundTrip shall not modify requests, here we modify it to fix error
		// messages see https://github.com/scylladb/mermaid/pkg/issues/266.
		// This is legit because we own the whole process. The modified request
		// is not being sent.
		req.Host = h
		req.URL.Host = h

		resp, err := next.RoundTrip(r)

		// Mark response
		if hpr != nil {
			switch {
			case err != nil:
				hpr.Mark(err)
			case resp.StatusCode == 401 || resp.StatusCode == 403 || resp.StatusCode >= 500:
				hpr.Mark(errPoolServerError)
			default:
				hpr.Mark(nil)
			}
		}

		return resp, err
	})
}

// ForceHost makes hostPool middleware use the given host instead of selecting
// one.
func ForceHost(ctx context.Context, host string) context.Context {
	return context.WithValue(ctx, ctxHost, host)
}
