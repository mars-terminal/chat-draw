package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"repositorie/config"
	http2 "repositorie/internal/server/http"

	"repositorie/internal/storage/redis"
	"time"

	authHandler "repositorie/internal/server/http/auth"
	docsHandler "repositorie/internal/server/http/docs"
	userHandler "repositorie/internal/server/http/user"
	authService "repositorie/internal/service/auth"
	messageService "repositorie/internal/service/message"
	userService "repositorie/internal/service/user"
	messageStorage "repositorie/internal/storage/postgres/message"
	userStorage "repositorie/internal/storage/postgres/user"
	authStorage "repositorie/internal/storage/redis/auth"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"repositorie/internal/storage/postgres"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

var log = logrus.WithField("package", "main")

// @title           ChatDraw API
// @version         1.0
// @description     ChatDraw API specs.

// @contact.name   API Support
// @contact.url    https://google.com

// @host      localhost:8000
// @BasePath  /

// @securityDefinitions.apikey APIToken @in header @name Authorization
// @in header
// @name Authorization
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

		MessageTable string
		UsersTable   string

		AuthSalt     string
		AuthSignKey  string
		AuthTokenTTL time.Duration

		HttpPORT string
	}{
		DBPostgresHost:     config.GetStringOrDefault(viper.GetViper(), "db.postgres.host", "localhost"),
		DBPostgresPort:     viper.GetString("db.postgres.port"),
		DBPostgresUser:     viper.GetString("db.postgres.user"),
		DBPostgresPassword: viper.GetString("db.postgres.password"),
		DBPostgresName:     viper.GetString("db.postgres.dbname"),
		DBPostgresSSLMode:  viper.GetString("db.postgres.sslmode"),
		HttpPORT:           viper.GetString("http.port"),

		RDBHost:       viper.GetString("db.redis.host"),
		RDBPort:       viper.GetString("db.redis.port"),
		RDBAuthPrefix: config.GetStringOrDefault(viper.GetViper(), "db.redis.prefix.auth", "auth"),

		MessageTable: config.GetStringOrDefault(viper.GetViper(), "db.postgres.tables.message_table", "messages"),
		UsersTable:   config.GetStringOrDefault(viper.GetViper(), "db.postgres.tables.user_table", "users"),

		AuthSalt:     viper.GetString("auth.salt"),
		AuthSignKey:  viper.GetString("auth.sign_key"),
		AuthTokenTTL: viper.GetDuration("auth.token_ttl"),
	}

	log.Info(options)

	db, err := postgres.NewStore(ctx, postgres.Config{
		Host:     options.DBPostgresHost,
		Port:     options.DBPostgresPort,
		User:     options.DBPostgresUser,
		Password: options.DBPostgresPassword,
		DBName:   options.DBPostgresName,
		SSLMode:  options.DBPostgresSSLMode,
	})

	if err != nil {
		log.WithError(err).Fatal("failed to initialize postgres store")
	}

	rdb, err := redis.NewRedisStorage(ctx, redis.Config{
		Host: options.RDBHost,
		Port: options.RDBPort,
	})

	if err != nil {
		log.WithError(err).Fatal("failed to initialize redis store")
	}
	log.Info("db connected")

	storageMessage := messageStorage.NewStore(db, options.MessageTable)
	storageUser := userStorage.NewStore(db, options.UsersTable)
	storageAuth := authStorage.NewStore(rdb, options.RDBAuthPrefix)

	serviceMessage := messageService.NewService(storageMessage)
	serviceUser := userService.NewService(storageUser)
	serviceAuth := authService.NewService(options.AuthSalt, options.AuthSignKey, options.AuthTokenTTL, storageAuth, serviceUser, serviceMessage)

	handlerAuth := authHandler.NewHandler(serviceAuth)
	handlerDocs := docsHandler.NewHandler(options.HttpPORT)
	handlerUsers := userHandler.NewHandler(serviceUser)

	r := gin.New()

	handlerDocs.InitRoutes(r)
	handlerAuth.InitRoutes(r)
	handlerUsers.InitRoutes(r.Group("/", handlerAuth.Middleware()))

	srv := new(http2.Server)
	if err := srv.Run(options.HttpPORT, r); err != nil {
		log.WithError(err).Fatal("error occured while running server server")
	}
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
