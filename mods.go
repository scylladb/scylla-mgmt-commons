// Copyright (C) 2017 ScyllaDB

// +build modules

package scylla_mgmt_commons

// This file exists only to trick modules into keeping single vendor dir for
// entire repo.

import (
	_ "github.com/scylladb/scylla-mgmt-commons/managerclient"
	_ "github.com/scylladb/scylla-mgmt-commons/scyllaclient"
	_ "github.com/scylladb/scylla-mgmt-commons/utils/format"
	_ "github.com/scylladb/scylla-mgmt-commons/utils/httpx"
	_ "github.com/scylladb/scylla-mgmt-commons/utils/middleware"
	_ "github.com/scylladb/scylla-mgmt-commons/utils/retry"
	_ "github.com/scylladb/scylla-mgmt-commons/utils/timeutc"
	_ "github.com/scylladb/scylla-mgmt-commons/utils/uuid"
)
