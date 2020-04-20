package concurrency

import (
	"reflect"
	"testing"
	"time"
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

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}
