package greetings_test

import (
	"fmt"
	"testing"

	"github.com/roguexray007/bootcamp/greetings"
)

func testHello(t *testing.T) {
	expected := "Hi, kunal . Welcome!"
	got := greetings.Hello("kunal")
	if got != expected {
		fmt.Println("Hello-error: ", got)
	}
}

func testHelloFromMaster(t *testing.T) {
	expected := "kunal's soul is mine "
	got := greetings.Hello("kunal")
	if got != expected {
		fmt.Println("Hello-From-Master-error: ", got)
	}
}
