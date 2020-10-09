// Copyright (C) 2017 ScyllaDB

package scyllaclienttest

import (
	"net"
	"net/http"
	"testing"
)

func NewFakeScyllaServer(t *testing.T, file string) (client *Client, closeServer func()) {
	return NewFakeScyllaServerMatching(t, FileMatcher(file))
}

func NewFakeScyllaServerRequestChecker(t *testing.T, file string, check func(t *testing.T, r *http.Request)) (client *Client, closeServer func()) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Error("ParseForm() error", err)
		}
		check(t, r)
		// Emulate ScyllaDB bug
		r.Header.Set("Content-Type", "text/plain")
		SendFile(t, w, FileMatcher(file)(r))
	})
	return NewFakeScyllaServerWithHandler(t, h)
}

func NewFakeScyllaServerMatching(t *testing.T, m Matcher) (client *Client, closeServer func()) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Error("ParseForm() error", err)
		}
		// Emulate ScyllaDB bug
		r.Header.Set("Content-Type", "text/plain")
		SendFile(t, w, m(r))
	})
	return NewFakeScyllaServerWithHandler(t, h)
}

func NewFakeScyllaServerWithHandler(t *testing.T, h http.Handler) (client *Client, closeServer func()) {
	host, port, closeServer := MakeServer(t, h)
	client = MakeClient(t, host, port)
	return
}

func NewFakeScyllaV2Server(t *testing.T, file string) (client *ConfigClient, closeServer func()) {
	return NewFakeScyllaV2ServerMatching(t, FileMatcher(file))
}

func NewFakeScyllaV2ServerMatching(t *testing.T, m Matcher) (client *ConfigClient, closeServer func()) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		SendFile(t, w, m(r))
	})
	return NewFakeScyllaV2ServerWithHandler(t, h)
}

func NewFakeScyllaV2ServerWithHandler(t *testing.T, h http.Handler) (client *ConfigClient, closeServer func()) {
	host, port, closeServer := MakeServer(t, h)
	client = NewConfigClient(net.JoinHostPort(host, port))
	return
}
