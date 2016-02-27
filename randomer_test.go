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

func TestSettings(t *testing.T) {
	randomer := setupRandomer()

	randomer.Settings("")
}
