package main

import (
	"context"
	"os"
	"repositorie/config"
	"repositorie/internal/storage/redis"

	authService "repositorie/internal/service/auth"
	messageService "repositorie/internal/service/message"
	userService "repositorie/internal/service/user"
	messageStorage "repositorie/internal/storage/postgres/message"
	userStorage "repositorie/internal/storage/postgres/user"
	authStorage "repositorie/internal/storage/redis/auth"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	http2 "repositorie/internal/http"
	"repositorie/internal/http/handler"
	"repositorie/internal/storage/postgres"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

var log = logrus.WithField("package", "main")

func main() {
	ctx := context.Background()

	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs : %s", err.Error())
	}

	var options = struct {
		DBPostgresHost     string
		DBPostgresPort     string
		DBPostgresUser     string
		DBPostgresPassword string
		DBPostgresName     string
		DBPostgresSSLMode  string

		RDBHost       string
		RDBPort       string
		RDBAuthPrefix string
		MessageTable  string
		UsersTable    string

		HttpPORT string
	}{
		DBPostgresHost:     config.GetStringOrDefault(viper.GetViper(), "db.postgres.host", "localhost"),
		DBPostgresPort:     viper.GetString("db.postgres.port"),
		DBPostgresUser:     viper.GetString("db.postgres.user"),
		DBPostgresPassword: os.Getenv("USE_DB_PASSWORD"),
		DBPostgresName:     viper.GetString("db.postgres.dbname"),
		DBPostgresSSLMode:  viper.GetString("db.postgres.sslmode"),
		HttpPORT:           viper.GetString("http.port"),

		RDBHost:       viper.GetString("db.redis.host"),
		RDBPort:       viper.GetString("db.redis.port"),
		RDBAuthPrefix: config.GetStringOrDefault(viper.GetViper(), "db.redis.prefix.auth", "auth"),

		MessageTable: config.GetStringOrDefault(viper.GetViper(), "db.postgres.tables.message_table", "messages"),
		UsersTable:   config.GetStringOrDefault(viper.GetViper(), "db.postgres.tables.user_table", "users"),
	}

	db, err := postgres.NewStore(ctx, postgres.Config{
		Host:     options.DBPostgresHost,
		Port:     options.DBPostgresPort,
		User:     options.DBPostgresUser,
		Password: options.DBPostgresPassword,
		DBName:   options.DBPostgresName,
		SSLMode:  options.DBPostgresSSLMode,
	})

	rdb, err := redis.NewRedisStorage(ctx, redis.Config{
		Host: options.RDBHost,
		Port: options.RDBPort,
	})

	if err != nil {
		log.WithError(err).Fatal("failed to initialize db")
	}
	log.Info("db connected")

	storageMessage := messageStorage.NewMessageStore(db, options.MessageTable)
	storageUser := userStorage.NewUserStore(db, options.UsersTable)
	storageAuth := authStorage.NewAuthStorage(rdb, options.RDBAuthPrefix)

	serviceMessage := messageService.NewMessageService(storageMessage)
	serviceUser := userService.NewService(storageUser)
	serviceAuth := authService.NewService(storageAuth, serviceUser, serviceMessage)

	handlers := handler.NewHandler(serviceAuth)

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
