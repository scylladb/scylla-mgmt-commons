// Copyright (C) 2017 ScyllaDB

package format

import (
	"testing"
)

func TestByteCounting(t *testing.T) {
	table := []struct {
		Bytes     int64
		Formatted string
	}{
		{
			Bytes:     1000,
			Formatted: "1000B",
		},
		{
			Bytes:     1024,
			Formatted: "1KiB",
		},
		{
			Bytes:     2 * 1024 * 1024,
			Formatted: "2MiB",
		},
		{
			Bytes:     2 * 1024 * 1024 * 1024,
			Formatted: "2.00GiB",
		},
		{
			Bytes:     123*1024*1024*1024 + 100*1024*1024 + 10*1024*1024,
			Formatted: "123.11GiB",
		},
		{
			Bytes:     2*1024*1024*1024*1024 + 10*1024*1024*1024,
			Formatted: "2.010TiB",
		},
		{
			Bytes:     2*1024*1024*1024*1024*1024 + 10*1024*1024*1024*1024,
			Formatted: "2.010PiB",
		},
		{
			Bytes:     2*1024*1024*1024*1024*1024*1024 + 10*1024*1024*1024*1024*1024,
			Formatted: "2.010EiB",
		},
	}

	for i := range table {
		test := table[i]
		t.Run(test.Formatted, func(t *testing.T) {
			t.Parallel()

			if s := StringByteCount(test.Bytes); s != test.Formatted {
				t.Errorf("StringByteCount() expected %s got %s", test.Formatted, s)
			}
			b, err := ParseByteCount(test.Formatted)
			if err != nil {
				t.Errorf("ParseByteCount() err %s", err)
			}
			if StringByteCount(b) != test.Formatted {
				t.Errorf("ParseByteCount() expected %s got %s", StringByteCount(test.Bytes), StringByteCount(b))
			}
		})
	}
}
