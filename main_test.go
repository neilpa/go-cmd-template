package main

import (
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		in string
		out string
		exit int
	} {
		{ "-v", version, 0 },
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			var w strings.Builder
			exit := realMain(strings.Split(tt.in, " "), &w)
			if exit != tt.exit {
				t.Fatalf("exit: got %d, want %d", exit, tt.exit)
			}
			got := strings.TrimSpace(w.String())
			if tt.out != got {
				t.Errorf("got\n%s\nwant\n%s", got, tt.out)
			}
		})
	}
}

