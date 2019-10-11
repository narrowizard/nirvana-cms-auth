package utils

import "testing"

func TestMatchURL(t *testing.T) {
	var pathname = "/material/{\\d+}"
	var url = "/material/15"
	var unmatchedURL = "/material/aaa15"
	var unmatchedURL2 = "/material/15/put"
	if !MatchURL(pathname, url) {
		t.Errorf("expected url %s match pathname %s, but got unmatched", url, pathname)
	}
	if MatchURL(pathname, unmatchedURL) {
		t.Errorf("expected url %s unmatch pathname %s, but got matched", unmatchedURL, pathname)
	}
	if MatchURL(pathname, unmatchedURL2) {
		t.Errorf("expected url %s unmatch pathname %s, but got matched", unmatchedURL2, pathname)
	}
}
