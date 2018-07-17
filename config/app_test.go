package config

import (
	"net/http"
	"testing"
)

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

func TestRouteGetLinkName(t *testing.T) {
	r := Route{
		Method: http.MethodGet,
		Path:   "/foo/bar/baz",
	}

	expected := "get-foo-bar-baz"
	actual := r.GetLinkName()
	if expected != actual {
		t.Fatalf("expected: %v\n\ngot: %v\n\n", expected, actual)
	}
}
