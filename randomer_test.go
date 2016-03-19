package randomer_test

import (
	"testing"

	"github.com/worzeel/randomer"
)

func setupRandomer(settings *randomer.Settings) *randomer.Randomer {
	var s *randomer.Settings
	if settings != nil {
		s = settings
	}
	return randomer.NewRandomer(s)
}

func TestCreate(t *testing.T) {
	r := setupRandomer(nil)

	if r == nil {
		t.Fail()
	}
}

func TestICanGetSettings(t *testing.T) {
	s := randomer.NewSettings()
	r := setupRandomer(s)

	gs := r.GetSettings()

	if gs == nil {
		t.Fail()
	}
}

func TestICanGetSettingsIPassedIn(t *testing.T) {
	s := randomer.NewSettings()

	testChars := "ABCD"

	s.SetCharacters(testChars)
	r := setupRandomer(s)

	gs := r.GetSettings()
	if gs == nil {
		t.Fail()
	}

	if gs.GetCharacters() != testChars {
		t.Fail()
	}
}

func TestICanGetSomeDataFromCallingRandomString(t *testing.T) {
	r := setupRandomer(nil)

	randomString := r.GetRandomString()

	if len(randomString) == 0 {
		t.Fail()
	}
}
