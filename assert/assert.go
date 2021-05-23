package assert

import (
	"testing"
)

func Equal(t *testing.T, a, b interface{}) {
	t.Helper()
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}

func True(t *testing.T, a bool) {
	t.Helper()
	if !a {
		t.Errorf("Not True %t", a)
	}
}

func False(t *testing.T, a bool) {
	t.Helper()
	if a {
		t.Errorf("Not True %t", a)
	}
}

func Nil(t *testing.T, a interface{}) {
	t.Helper()
	if a != nil {
		t.Error("Not Nil")
	}
}

func NotNil(t *testing.T, a interface{}) {
	t.Helper()
	if a == nil {
		t.Error("Is Nil")
	}
}
