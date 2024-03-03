package toolkit

import (
	"testing"
)

func TestTools_RandomString(t *testing.T) {
	var testTools Tools

	s, _ := testTools.RandomString(5)
	if len(s) != 5 {
		t.Error("Wrong length random string returned")
	}
}
