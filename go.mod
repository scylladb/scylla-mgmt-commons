module github.com/scylladb/scylla-mgmt-commons

go 1.14

require (
	github.com/scylladb/scylla-mgmt-commons/format v0.0.0-20201007123433-27727073dfc6
	github.com/scylladb/scylla-mgmt-commons/managerclient v0.0.0-20201007123433-27727073dfc6
	github.com/scylladb/scylla-mgmt-commons/timeutc v0.0.0-20201007123433-27727073dfc6
	github.com/scylladb/scylla-mgmt-commons/uuid v0.0.0-20201007123433-27727073dfc6
)

replace (
	github.com/scylladb/scylla-mgmt-commons/format => ./format
	github.com/scylladb/scylla-mgmt-commons/managerclient => ./managerclient
	github.com/scylladb/scylla-mgmt-commons/timeutc => ./timeutc
	github.com/scylladb/scylla-mgmt-commons/uuid => ./uuid
)
