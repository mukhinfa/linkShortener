package main

import (
	"fmt"

	"github.com/muhinfa/linkShortener/configs"
	"github.com/muhinfa/linkShortener/internal/auth"
	"github.com/muhinfa/linkShortener/pkg/db"

	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
