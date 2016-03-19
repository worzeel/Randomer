package randomer_test

import (
	"testing"

	"github.com/worzeel/randomer"
)

func setupRandomer(setting randomer.Setter) *randomer.Randomer {
	var s randomer.Setter
	if setting != nil {
		s = setting
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
	s := randomer.NewSetting()
	r := setupRandomer(s)

	gs := r.GetSettings()

	if gs == nil {
		t.Fail()
	}
}

func TestICanGetSettingsIPassedIn(t *testing.T) {
	s := randomer.NewSetting()

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

func TestICanGetADefaultRandomStringOfTenCharacters(t *testing.T) {
	r := setupRandomer(nil)

	randomString := r.GetRandomString()

	if len(randomString) != 10 {
		t.Fail()
	}
}
