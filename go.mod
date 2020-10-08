module github.com/scylladb/scylla-mgmt-commons

go 1.14

require (
	github.com/scylladb/scylla-mgmt-commons/managerclient v0.0.0-20201007133817-3a5a9f249639
	github.com/scylladb/scylla-mgmt-commons/scyllaclient v0.0.0-20201007133817-3a5a9f249639
	github.com/scylladb/scylla-mgmt-commons/utils/format v0.0.0-20201007133817-3a5a9f249639
	github.com/scylladb/scylla-mgmt-commons/utils/httpx v0.0.0-20201008140841-cc57d906f71c
	github.com/scylladb/scylla-mgmt-commons/utils/middleware v0.0.0-20201007133817-3a5a9f249639
	github.com/scylladb/scylla-mgmt-commons/utils/retry v0.0.0-20201008140841-cc57d906f71c
	github.com/scylladb/scylla-mgmt-commons/utils/timeutc v0.0.0-20201008140841-cc57d906f71c
	github.com/scylladb/scylla-mgmt-commons/utils/uuid v0.0.0-20201007133817-3a5a9f249639
)

replace (
	github.com/scylladb/scylla-mgmt-commons/managerclient => ./managerclient
	github.com/scylladb/scylla-mgmt-commons/scyllaclient => ./scyllaclient
	github.com/scylladb/scylla-mgmt-commons/utils/format => ./utils/format
	github.com/scylladb/scylla-mgmt-commons/utils/httpx => ./utils/httpx
	github.com/scylladb/scylla-mgmt-commons/utils/middleware => ./utils/middleware
	github.com/scylladb/scylla-mgmt-commons/utils/retry => ./utils/retry
	github.com/scylladb/scylla-mgmt-commons/utils/timeutc => ./utils/timeutc
	github.com/scylladb/scylla-mgmt-commons/utils/uuid => ./utils/uuid
)
