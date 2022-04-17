package helpers_test

import (
	"goSpider/helpers"
	"testing"
)

const base string = "https://google.com"

func TestProvisionUrlBackSlash(t *testing.T) {
	toProvision := "/example/test"
	want := base + toProvision
	if value, _ := helpers.ProvisionURL(toProvision, base); value != want {
		t.Fatalf("Expected %s got %s", want, value)
	}
}

func TestProvisionUrlDotBackSlash(t *testing.T) {
	toProvision := "./example/test"
	want := base + toProvision[1:]
	if value, _ := helpers.ProvisionURL(toProvision, base); value != want {
		t.Fatalf("Expected %s got %s", want, value)
	}
}

func TestProvisionUrlEmpty(t *testing.T) {
	var toProvision string
	toProvision = ""
	if _, err := helpers.ProvisionURL(toProvision, base); err == nil {
		t.Fatalf("Expected Error Cannot handle empty url")
	}
	toProvision = "#"
	if _, err := helpers.ProvisionURL(toProvision, base); err == nil {
		t.Fatalf("Expected Error Cannot handle # string")
	}
}

func TestProvisionUrlFullLink(t *testing.T) {
	toProvision := "http://example.com/"
	want := toProvision
	if value, _ := helpers.ProvisionURL(toProvision, base); value != want {
		t.Fatalf("Expected %s got %s", want, toProvision)
	}
}
