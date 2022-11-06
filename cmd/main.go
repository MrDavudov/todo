package main

import (
	"log"

	"github.com/MrDavudov/todo"
	"github.com/MrDavudov/todo/pkg/handler"
	"github.com/MrDavudov/todo/pkg/repository"
	servise "github.com/MrDavudov/todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	servises := servise.NewService(repos)
	handlers := handler.NewHandler(servises)

	srv := new(todo.Server)
	if err := srv.Start("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("errors occured while running http server: %s", err)
	}
}