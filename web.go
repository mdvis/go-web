package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Trace("trace msg")
	log.Debug("debug msg")
	log.Info("info msg")
	log.Warn("warn msg")
	log.Error("error msg")

	log.WithFields(log.Fields{"l_msg": "msg"}).Trace("trace msg")
	log.WithFields(log.Fields{"l_msg": "msg"}).Debug("debug msg")
	log.WithFields(log.Fields{"l_msg": "msg"}).Info("info msg")
	log.WithFields(log.Fields{"l_msg": "msg"}).Warn("warn msg")
	log.WithFields(log.Fields{"l_msg": "msg"}).Error("error msg")
	// log.WithFields(log.Fields{"msg": "msg"}).Fatal("fatal msg")
	// log.WithFields(log.Fields{"msg": "msg"}).Panic("panic msg")

	http.HandleFunc("/", static)
	http.HandleFunc("/api", api)
	http.HandleFunc("/login", login)
	http.HandleFunc("/match", match)
	http.HandleFunc("/escape", escape)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/event_source", sse)
	http.HandleFunc("/register", register)

	err := http.ListenAndServe(":8090", nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
