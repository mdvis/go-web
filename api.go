// ------
// name: api.go
// author: Deve
// date: 2025-02-19
// ------
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func db(w http.ResponseWriter, r *http.Request) {
	md()
	rds()
	del("8")
	query("11")
	alter("cc", "1")
	insert("aa", "2024-12-21")

}

func escape(w http.ResponseWriter, r *http.Request) {
	str := "<script>console.log(666);</script>"

	fmt.Fprint(w, str, template.HTMLEscapeString(str))
	template.HTMLEscape(w, []byte(str))
}

func api(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	type Info struct {
		Other string `json:"other"`
	}

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age,string"`
		Sex  int    `json:"-"`
		Info Info   `json:"info" json:"j"`
	}

	u := User{"gouzi", 24, 0, Info{"??"}}
	a, err := json.Marshal(u)

	if err != nil {
		log.Fatal(err)
	}

	var e User

	err = json.Unmarshal(a, &e)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(a))
}
