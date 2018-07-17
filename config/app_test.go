package config

import "testing"

func TestResourceGetLinkName(t *testing.T) {
	r := Resource{
		Name: "FOOBAR",
	}

	expected := "foobar"
	actual := r.GetLinkName()
	if expected != actual {
		t.Fatalf("expected: %v\n\ngot: %v\n\n", expected, actual)
	}
}
