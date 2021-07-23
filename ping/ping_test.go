package ping_test

import (
	"testing"

	"github.com/roguexray007/bootcamp/ping"
)

func TestPing(t *testing.T) {
	expected := "pong"
	got := ping.Ping()
	if got != expected {
		t.Error("some issue in package")
	}
}
