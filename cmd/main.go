package main

import (
	"log"
	"os"

	"github.com/MrDavudov/todo/internal/server"
	"github.com/MrDavudov/todo/pkg/handler"
	"github.com/MrDavudov/todo/pkg/repository"
	"github.com/MrDavudov/todo/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error Initializing configs: %s",err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err)
	}

	db, err :=  repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err)
	}

	repos := repository.NewRepository(db)
	servises := servise.NewService(repos)
	handlers := handler.NewHandler(servises)

	srv := new(server.Server)
	if err := srv.Start(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("errors occured while running http server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}