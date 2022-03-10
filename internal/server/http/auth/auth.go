package auth

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.WithFields(logrus.Fields{
	"package": "auth",
	"layer":   "server",
})

const (
	TokenPrefix = "Bearer"
)
