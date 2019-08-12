package utils

import (
	"reflect"
	"testing"
)

func TestMessage(t *testing.T) {
	result := Message(false, "Something went wrong!")

	expectation := map[string]interface{}{"status": false, "message": "Something went wrong!"}

	if !reflect.DeepEqual(result, expectation) {
		t.Error("Return of Message does not meet expectation")
	}
}
