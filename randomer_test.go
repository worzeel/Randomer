package randomer_test

import (
	"testing"

	"github.com/worzeel/randomer"
)

func setupRandomer() *randomer.Randomer {
	return new(randomer.Randomer)
}

func TestCreate(t *testing.T) {
	randomer := setupRandomer()

	if randomer == nil {
		t.Fail()
	}
}

func TestICanCallSettings(t *testing.T) {
	randomer := setupRandomer()

	randomer.Settings("")
}

func TestICanCallGetRandomString(t *testing.T) {
	randomer := setupRandomer()

	randomer.GetRandomString()
}

func TestIGetSomeDataFromCallingRandomString(t *testing.T) {
	randomer := setupRandomer()

	randomString := randomer.GetRandomString()

	if len(randomString) == 0 {
		t.Fail()
	}
}
