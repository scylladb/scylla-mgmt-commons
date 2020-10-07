// Copyright (C) 2017 ScyllaDB

// +build modules

package scylla_mgmt_commons

// This file exists only to trick modules into keeping single vendor dir for
// entire repo.

import (
	_ "github.com/scylladb/scylla-mgmt-commons/format"
	_ "github.com/scylladb/scylla-mgmt-commons/httpx"
	_ "github.com/scylladb/scylla-mgmt-commons/managerclient"
	_ "github.com/scylladb/scylla-mgmt-commons/middleware"
	_ "github.com/scylladb/scylla-mgmt-commons/retry"
	_ "github.com/scylladb/scylla-mgmt-commons/scyllaclient"
	_ "github.com/scylladb/scylla-mgmt-commons/timeutc"
	_ "github.com/scylladb/scylla-mgmt-commons/uuid"
)
