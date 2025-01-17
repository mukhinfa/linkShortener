package main

import (
	"fmt"

	"github.com/muhinfa/linkShortener/configs"
	"github.com/muhinfa/linkShortener/internal/auth"
	"github.com/muhinfa/linkShortener/internal/link"
	"github.com/muhinfa/linkShortener/pkg/db"
	"github.com/muhinfa/linkShortener/pkg/middleware"

	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	linkRepository := link.NewLinkRepository(db)

	//Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.HandlerDeps{
		Repository: linkRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logging(router),
	}

	fmt.Println("Server is listening on port 8081")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
