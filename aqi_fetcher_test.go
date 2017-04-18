package main

import (
	"testing"
)

func TestGenFeedURLs(t *testing.T) {
	r := genFeedURLs()
	expectedURL := "http://feed.aqicn.org/feed/tianjin/en/feed.v1.json"
	for _, url := range r {
		if url == expectedURL {
			return
		}
	}
	t.Errorf("Can't found expected value %s in <genFeedURLs> return values %s", expectedURL, r)
}
