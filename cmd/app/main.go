package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	http2 "repositorie/internal/http"
	"repositorie/internal/http/handler"
	"repositorie/internal/storage"
	"repositorie/internal/storage/postgres"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

var log = logrus.WithField("package", "main")

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs : %s", err.Error())
	}

	var options = struct {
		DBHost     string
		DBPort     string
		DBUser     string
		DBPassword string
		DBName     string
		DBSSLMode  string

		HttpPORT string
	}{
		DBHost:     viper.GetString("db.host"),
		DBPort:     viper.GetString("db.port"),
		DBUser:     viper.GetString("db.user"),
		DBPassword: os.Getenv("USE_DB_PASSWORD"),
		DBName:     viper.GetString("db.dbname"),
		DBSSLMode:  viper.GetString("db.sslmode"),
		HttpPORT:   viper.GetString("http.port"),
	}

	db, err := postgres.NewStore(postgres.Config{
		Host:     options.DBHost,
		Port:     options.DBPort,
		User:     options.DBUser,
		Password: options.DBPassword,
		DBName:   options.DBName,
		SSLMode:  options.DBSSLMode,
	})

	if err != nil {
		log.WithError(err).Fatal("failed to initialize db")
	}
	log.Info("db connected")

	repos := storage.NewStorage(db)
	handlers := handler.NewHandler(repos)

	srv := new(http2.Server)
	if err := srv.Run(options.HttpPORT, handlers.InitRoutes()); err != nil {
		log.WithError(err).Fatal("error occured while running http http")
	}

}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
