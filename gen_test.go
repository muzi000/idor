package idor

import "testing"

func TestGen(t *testing.T) {
	urls := []string{
		"127.0.0.1/user/01",
		"127.0.0.1/user/01/aaa",
		"127.0.0.1/user/?id=1",
		"127.0.0.1/user/?id=1&name=root&pri=0",
	}
	a := make(map[string][]Payload, len(urls))
	for _, u := range urls {
		a[u] = Check(u)
		var exp []string
		for _, p := range a[u] {
			exp = append(exp, Gen(p)...)
		}
		for i := range exp {
			t.Logf("exp: %s", exp[i])
		}
	}
}
