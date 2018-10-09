package main

import (
	"net/http"
	"time"

	"github.com/yznu-cn/yznu-tool/yznu-go/config"
	"github.com/yznu-cn/yznu-tool/yznu-go/router"

	log "github.com/sirupsen/logrus"
)

func main() {
	r := router.Router()
	s := &http.Server{
		Addr:           ":" + config.ServerConf.Port,
		Handler:        r,
		ReadTimeout:    120 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.SetLevel(log.DebugLevel)
	s.ListenAndServe()
}
