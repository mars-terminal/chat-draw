package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/mars-terminal/chat-draw/config"
	http2 "github.com/mars-terminal/chat-draw/internal/server/http"
	authHandler "github.com/mars-terminal/chat-draw/internal/server/http/auth"
	docsHandler "github.com/mars-terminal/chat-draw/internal/server/http/docs"
	messageHandler "github.com/mars-terminal/chat-draw/internal/server/http/message"
	userHandler "github.com/mars-terminal/chat-draw/internal/server/http/user"
	authService "github.com/mars-terminal/chat-draw/internal/service/auth"
	messageService "github.com/mars-terminal/chat-draw/internal/service/message"
	userService "github.com/mars-terminal/chat-draw/internal/service/user"
	"github.com/mars-terminal/chat-draw/internal/storage/postgres"
	messageStorage "github.com/mars-terminal/chat-draw/internal/storage/postgres/message"
	userStorage "github.com/mars-terminal/chat-draw/internal/storage/postgres/user"
	"github.com/mars-terminal/chat-draw/internal/storage/redis"
	authStorage "github.com/mars-terminal/chat-draw/internal/storage/redis/auth"
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

// @securityDefinitions.apikey ApiKeyAuth
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
	handlerMessages := messageHandler.NewHandler(serviceMessage)

	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	handlerDocs.InitRoutes(r)
	handlerAuth.InitRoutes(r)
	{
		authGroup := r.Group("/", handlerAuth.Middleware())

		handlerUsers.InitRoutes(authGroup)
		handlerMessages.InitRoutes(authGroup)
	}

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
