module github.com/scylladb/scylla-mgmt-commons

go 1.14

require (
	github.com/scylladb/scylla-mgmt-commons/format v0.0.0-20201007133817-3a5a9f249639
	github.com/scylladb/scylla-mgmt-commons/httpx v0.0.0-20201007133817-3a5a9f249639
	github.com/scylladb/scylla-mgmt-commons/managerclient v0.0.0-20201007133817-3a5a9f249639
	github.com/scylladb/scylla-mgmt-commons/middleware v0.0.0-20201007133817-3a5a9f249639
	github.com/scylladb/scylla-mgmt-commons/retry v0.0.0-20201007133817-3a5a9f249639
	github.com/scylladb/scylla-mgmt-commons/scyllaclient v0.0.0-20201007133817-3a5a9f249639
	github.com/scylladb/scylla-mgmt-commons/timeutc v0.0.0-20201007133817-3a5a9f249639
	github.com/scylladb/scylla-mgmt-commons/uuid v0.0.0-20201007133817-3a5a9f249639
)

replace (
	github.com/scylladb/scylla-mgmt-commons/utils/format => ./utils/format
	github.com/scylladb/scylla-mgmt-commons/utils/httpx => ./utils/httpx
	github.com/scylladb/scylla-mgmt-commons/managerclient => ./managerclient
	github.com/scylladb/scylla-mgmt-commons/utils/middleware => ./utils/middleware
	github.com/scylladb/scylla-mgmt-commons/utils/retry => ./utils/retry
	github.com/scylladb/scylla-mgmt-commons/scyllaclient => ./scyllaclient
	github.com/scylladb/scylla-mgmt-commons/utils/timeutc => ./utils/timeutc
	github.com/scylladb/scylla-mgmt-commons/utils/uuid => ./utils/uuid
)
