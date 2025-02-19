// ------
// name: match.go
// author: Deve
// date: 2025-02-18
// ------

package main

import (
	"fmt"
	"net/http"
	"regexp"

	log "github.com/sirupsen/logrus"
)

func match(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("p", r.PostForm)
	fmt.Println("q", r.URL.Query())
	pattern := "a|b"
	re := regexp.MustCompile(pattern)
	str1 := "a"
	str2 := "b"
	str3 := "c"
	log.WithFields(log.Fields{"str1": str1, "str2": str2, "str3": str3}).Info(re.MatchString(str1), re.MatchString(str2), re.MatchString(str3))
}
