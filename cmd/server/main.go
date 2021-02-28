package main

import (
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	svr := server{
		router: NewRouter(),
		logger: logrus.New(),
	}

	svr.routes()

	log.Fatal(http.ListenAndServe(":8080", svr.router))
}
