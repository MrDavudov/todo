package main

import (
	"os"

	"github.com/MrDavudov/todo/internal/server"
	"github.com/MrDavudov/todo/pkg/handler"
	"github.com/MrDavudov/todo/pkg/repository"
	servise "github.com/MrDavudov/todo/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error Initializing configs: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	}, viper.GetString("db_url"))
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err)
	}
	logrus.Info("Run db postgres")

	repos := repository.NewRepository(db)
	servises := servise.NewService(repos)
	handlers := handler.NewHandler(servises)

	srv := new(server.Server)
	if err := srv.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("errors occured while running http server: %s", err)
	}
	logrus.Info("Start server")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
