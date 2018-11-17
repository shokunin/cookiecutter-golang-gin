package main

import (
<<<<<<< HEAD
	"github.com/{{cookiecutter.github_handle}}/{{cookiecutter.app_name}}/handlers"
=======
	"{{cookiecutter.app_name}}/handlers/healthcheck"
>>>>>>> 8b5d5dbe68327d37f3cfddbe57e078ff65fa90e8
	"github.com/gin-gonic/gin"
	"github.com/shokunin/contrib/ginrus"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	router := gin.New()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true, "{{cookiecutter.app_name}}"))

	// Start routes
	router.GET("/health", healthcheck.HealthCheck)

	// RUN rabit run
	router.Run() // listen and serve on 0.0.0.0:8080
}
