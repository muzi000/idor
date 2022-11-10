package idor

import (
	"strconv"
	"strings"
)

// Gen 生成url
func Gen(p Payload) []string {
	var urls []string
	switch p.Method {
	case "NorNum":
		urls = append(urls, addExt(p)...)
		urls = append(urls, newNum(p)...)
	case "Params":
		urls = append(urls, newNum(p)...)
		urls = append(urls, doubleParam(p)...)
	default:
	}
	return urls
}

// AddExt 添加扩展
func addExt(p Payload) []string {
	if !strings.HasSuffix(p.Url, "%EXE%") {
		return []string{}
	}
	exts := []string{"json", "txt", "log"}
	var n []string
	for i := range exts {
		n = append(n, p.Origin+"."+exts[i])
	}
	return n
}

// NewNum 产生0，1，极大数，负数
func newNum(p Payload) []string {
	num := []int{0, 1, 9999999999, -1}
	var n []string
	for i := range num {
		n = append(n, strings.Replace(p.Url, "%EXE%", strconv.Itoa(int(num[i])), 1))
	}
	return n
}

func doubleParam(p Payload) []string {
	i1 := strings.LastIndex(p.Url, "=%EXE%")
	i2 := strings.LastIndex(p.Url[:i1], "&")
	if i2 == -1 {
		i2 = strings.Index(p.Url, "?")
	}
	name := p.Origin[i2+1 : i1]
	return []string{p.Origin + "&" + name + "=1"}

}
