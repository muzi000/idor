package idor

import "testing"

func TestCheck(t *testing.T) {
	urls := []string{
		"127.0.0.1/user/01",
		"127.0.0.1/user/01/aaa",
	}
	a := make(map[string][]Payload, len(urls))
	for _, u := range urls {
		a[u] = Check(u)
		for _, p := range a[u] {
			if p.Method == "NorNum" {
				t.Logf("payload check: %s", p.Url)
			}
		}
	}
}
