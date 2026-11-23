package main

import "testing"

func TestRoutes(t *testing.T) {
	cases := []struct {
		path, message string
		status        int
		body          string
	}{
		{"/", "release-1", 200, "release-1\n"},
		{"/", "release-2", 200, "release-2\n"},
		{"/error", "ignored", 500, "controlled error\n"},
		{"/missing", "ignored", 404, "not found\n"},
	}
	for _, c := range cases {
		status, body := response(c.path, c.message)
		if status != c.status || body != c.body {
			t.Fatalf("%s: got %d %q", c.path, status, body)
		}
	}
}
