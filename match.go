// ------
// name: match.go
// author: Deve
// date: 2025-02-18
// ------

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

func match(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("p", r.PostForm)
	fmt.Println("q", r.URL.Query())

	if r.Method == "GET" {
		lg := log.New()
		fl, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		mw := io.MultiWriter(os.Stdout, fl)
		defer fl.Close()

		if err != nil {
			log.Fatal("log")
		}

		// lg.Out = fl
		lg.SetOutput(mw)
		lg.SetLevel(log.DebugLevel)

		lg.Trace("trace msg")
		lg.Debug("debug msg")
		lg.Info("info msg")
		lg.Warn("warn msg")
		lg.Error("error msg")

		pattern := "a|b"
		re := regexp.MustCompile(pattern)
		str1 := "a"
		str2 := "b"
		str3 := "c"
		log.WithFields(log.Fields{"str1": str1, "str2": str2, "str3": str3}).Info(re.MatchString(str1), re.MatchString(str2), re.MatchString(str3))
	}

	if r.Method == "POST" {
		src := "<HTML></HTML>"
		re, err := regexp.Compile("<[\\S\\s]+?>")
		if err != nil {
			log.Fatal("")
		}
		src = re.ReplaceAllStringFunc(src, strings.ToLower)

		one := re.Find([]byte(src))

		log.Info(string(one))

		fmt.Fprintln(w, src)
		src = re.ReplaceAllString(src, "body")
		fmt.Fprintln(w, src)
	}
}
