package idor

import (
	"regexp"
	"strings"
)

type Payload struct {
	Method string
	Url    string
	Origin string
}

// 检测类型  /user/01, ?id=1, /user/{guid},/static/111.txt, /static/{time}.txt 替换点改为 %EXE%
func Check(url string) []Payload {
	var p Payload
	var ps []Payload
	var exist bool
	p, exist = ckeckNorNum(url)
	if exist {
		ps = append(ps, p)
	}

	params := parseQuery(url)
	for i := range params {
		p, exist = checkParams(url, params[i])
		if exist {
			ps = append(ps, p)
		}
	}
	return ps
}

func ckeckNorNum(url string) (Payload, bool) {
	r, _ := regexp.Compile(`/\w+/[0-9]+/`)
	p := r.Find([]byte(url + "/"))
	if len(p) == 0 {
		return Payload{}, false
	}
	fir, _, _ := strings.Cut(string(p[1:len(p)-1]), "/")
	newUrl := strings.Replace(url, string(p[:len(p)-1]), "/"+fir+"/%EXE%", 1)
	return Payload{Method: "NorNum", Url: newUrl, Origin: url}, true
}

func checkParams(url, param string) (Payload, bool) {
	name, _, ok := strings.Cut(param, "=")
	if !ok {
		return Payload{}, false
	}
	newp := name + "=%EXE%"
	newurl := strings.Replace(url, param, newp, 1)
	return Payload{Method: "Params", Url: newurl, Origin: url}, true
}

func parseQuery(url string) []string {
	_, queries, _ := strings.Cut(url, "?")
	var params []string
	for queries != "" {
		var key string
		key, queries, _ = strings.Cut(queries, "&")
		params = append(params, key)
	}
	return params
}
