package concurrency

import (
	"reflect"
	"testing"
)

const errorWebsiteURL = "waat://invalidurl"

func mockWebsiteChecker(url string) bool {
	return url != errorWebsiteURL
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://facebook.com",
		errorWebsiteURL,
	}

	want := map[string]bool{
		"https://google.com":   true,
		"https://facebook.com": true,
		errorWebsiteURL:        false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
