module github.com/scylladb/scylla-mgmt-commons/format

go 1.14

require (
	github.com/go-openapi/strfmt v0.19.5
	github.com/pkg/errors v0.9.1
	github.com/scylladb/scylla-mgmt-commons/timeutc v0.0.0-00010101000000-000000000000
)

replace github.com/scylladb/scylla-mgmt-commons/timeutc => ../timeutc
