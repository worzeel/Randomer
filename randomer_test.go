package randomer_test

import (
	"testing"

	"github.com/worzeel/randomer"
)

func setupRandomer(settings *randomer.Settings) *randomer.Randomer {
	var s *randomer.Settings
	if settings == nil {
		s = new(randomer.Settings)
	} else {
		s = settings
	}
	return randomer.NewRandomer(s)
}

func TestCreate(t *testing.T) {
	randomer := setupRandomer(nil)

	if randomer == nil {
		t.Fail()
	}
}

func TestICanGetSettings(t *testing.T) {
	randomer := setupRandomer(nil)

	s := randomer.GetSettings()

	if s == nil {
		t.Fail()
	}
}

func TestICanGetSettingsIPassedIn(t *testing.T) {
	s := new(randomer.Settings)

	testChars := "ABCD"

	s.SetCharacters(testChars)
	randomer := setupRandomer(s)

	if randomer.GetSettings().GetCharacters() != testChars {
		t.Fail()
	}
}

func TestICanCallGetRandomString(t *testing.T) {
	randomer := setupRandomer(nil)

	randomer.GetRandomString()
}

func TestIGetSomeDataFromCallingRandomString(t *testing.T) {
	randomer := setupRandomer(nil)

	randomString := randomer.GetRandomString()

	if len(randomString) == 0 {
		t.Fail()
	}
}
