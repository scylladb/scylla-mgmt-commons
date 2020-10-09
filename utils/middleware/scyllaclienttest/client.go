// Copyright (C) 2017 ScyllaDB

package scyllaclienttest

import (
	"context"
	"net/http"
	"sync"
	"testing"
	"time"

	api "github.com/go-openapi/runtime/client"
	apiMiddleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/hailocab/go-hostpool"
	"github.com/scylladb/go-log"
	"github.com/scylladb/scylla-mgmt-commons/scyllaclient"
	agentOperations "github.com/scylladb/scylla-mgmt-commons/scyllaclient/agent/client/operations"
	"github.com/scylladb/scylla-mgmt-commons/scyllaclient/agent/models"
	agentClient "github.com/scylladb/scylla-mgmt-commons/scyllaclient/scylla/client"
	scyllaClient "github.com/scylladb/scylla-mgmt-commons/scyllaclient/scylla/client"
	scyllaOperations "github.com/scylladb/scylla-mgmt-commons/scyllaclient/scylla/client/operations"
	scyllaV2Client "github.com/scylladb/scylla-mgmt-commons/scyllaclient/scylla_v2/client"
	"github.com/scylladb/scylla-mgmt-commons/utils/middleware"
)

// TestHost should be used if a function in test requires host parameter.
const TestHost = "127.0.0.1"

// ClientOption allows to modify configuration in MakeClient.
type ClientOption func(*Config)

type Config struct {
	BackoffConfig middleware.BackoffConfig
	Transport     http.RoundTripper
	Port          string
	Hosts         []string
	Timeout       time.Duration
}

type Client struct {
	Config Config

	scyllaOps *scyllaOperations.Client
	agentOps  *agentOperations.Client
}

// MakeClient creates a Client for testing. Typically host and port are set
// based on MakeServer result.
func MakeClient(t *testing.T, host, port string, opts ...ClientOption) *Client {
	t.Helper()

	setOpenAPIGlobals()

	logger := log.NewDevelopment()

	config := Config{}
	config.Transport = http.DefaultTransport
	config.Hosts = []string{host}
	config.Port = port

	for i := range opts {
		opts[i](&config)
	}

	pool := hostpool.NewEpsilonGreedy(config.Hosts, 5*time.Minute, &hostpool.LinearEpsilonValueCalculator{})

	if config.Transport == nil {
		config.Transport = http.DefaultTransport
	}
	transport := config.Transport
	transport = middleware.Timeout(transport, config.Timeout)
	transport = middleware.RequestLogger(transport, logger)
	transport = middleware.HostPool(transport, pool, config.Port)
	transport = middleware.FixScyllaContentType(transport)

	c := &http.Client{Transport: config.Transport}
	scyllaRuntime := api.NewWithClient(
		scyllaClient.DefaultHost, scyllaClient.DefaultBasePath, []string{"http"}, c,
	)
	agentRuntime := api.NewWithClient(
		agentClient.DefaultHost, agentClient.DefaultBasePath, []string{"http"}, c,
	)

	scyllaOps := scyllaOperations.New(middleware.Retryable(scyllaRuntime, config.BackoffConfig, config.BackoffConfig, len(config.Hosts), scyllaclient.StatusCodeOf, logger), strfmt.Default)
	agentOps := agentOperations.New(middleware.Retryable(agentRuntime, config.BackoffConfig, config.BackoffConfig, len(config.Hosts), scyllaclient.StatusCodeOf, logger), strfmt.Default)

	return &Client{
		Config:    config,
		scyllaOps: scyllaOps,
		agentOps:  agentOps,
	}
}

// NodeInfo returns basic information about `host` node.
func (c *Client) NodeInfo(ctx context.Context, host string) (*models.NodeInfo, error) {
	p := agentOperations.NodeInfoParams{
		Context: middleware.ForceHost(ctx, host),
	}
	resp, err := c.agentOps.NodeInfo(&p)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// ClusterName returns cluster name.
func (c *Client) ClusterName(ctx context.Context) (string, error) {
	resp, err := c.scyllaOps.StorageServiceClusterNameGet(&scyllaOperations.StorageServiceClusterNameGetParams{
		Context: ctx,
	})
	if err != nil {
		return "", err
	}

	return resp.Payload, nil
}

// ConfigClient provides means to interact with Scylla config API on a given
// host if it's directly accessible.
type ConfigClient struct {
	addr   string
	client *scyllaV2Client.Scylla2
}

func NewConfigClient(addr string) *ConfigClient {
	setOpenAPIGlobals()

	t := http.DefaultTransport
	t = middleware.FixScyllaContentType(t)
	c := &http.Client{
		Timeout:   30 * time.Second,
		Transport: t,
	}

	scyllaV2Runtime := api.NewWithClient(
		addr, scyllaV2Client.DefaultBasePath, scyllaV2Client.DefaultSchemes, c,
	)

	return &ConfigClient{
		addr:   addr,
		client: scyllaV2Client.New(scyllaV2Runtime, strfmt.Default),
	}
}

var setOpenAPIGlobalsOnce sync.Once

func setOpenAPIGlobals() {
	setOpenAPIGlobalsOnce.Do(func() {
		// Timeout is defined in http client that we provide in api.NewWithClient.
		// If Context is provided to operation, which is always the case here,
		// this value has no meaning since OpenAPI runtime ignores it.
		api.DefaultTimeout = 0
		// Disable debug output to stderr, it could have been enabled by setting
		// SWAGGER_DEBUG or DEBUG env variables.
		apiMiddleware.Debug = false
	})
}
