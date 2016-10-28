package exceptions_test

import (
	"fmt"
	"testing"

	"github.com/eljuanchosf/adventures-in-golang/03-exception-management"
)

func TestExceptions(t *testing.T) {
	result, err := exceptions.ThrowException()
	if err == nil {
		t.Error("ThrowException should have had returned and exception")
	}
	fmt.Sprintf("Returned value: %d", result)
}
