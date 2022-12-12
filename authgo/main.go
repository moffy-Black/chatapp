package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/moffy-Black/authgo/pkg/config"
	"github.com/moffy-Black/authgo/pkg/user"
)

func main() {
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", config.DbSource)
	if err != nil {
		log.Panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err.Error())
	}

	router := gin.Default()
	router.Use(errorMiddleware())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.FrontAppSource}
	router.Use(cors.New(corsConfig))
	router.GET("/", index(config.Version, config.PodName))
	router.POST("/signin", signIn(db))
	router.POST("/signup", signUp(db))
	router.Run(":8000")
}

func errorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		err := ctx.Errors.ByType(gin.ErrorTypePublic).Last()
		if err != nil {
			log.Print(err.Err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"sign":   0,
				"result": err.Error(),
			})
		}
	}
}

func index(version string, podName string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version": version,
			"podname": podName,
			"message": "moffyblack/user-app container",
		})
	}
}

func signIn(db *sql.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var (
			cu     user.ChatUser
			result gin.H
			status int
		)

		err := ctx.Bind(&cu)
		if err != nil {
			ctx.Error(err).SetType(gin.ErrorTypePublic)
			return
		}

		chatuser, err := cu.Get(db)

		if err != nil {
			ctx.Error(err).SetType(gin.ErrorTypePublic)
			return
		}

		result = gin.H{
			"sign":   1,
			"result": chatuser,
		}
		status = http.StatusOK

		ctx.JSON(status, result)
	}
}

func signUp(db *sql.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var (
			cu     user.ChatUser
			result gin.H
			status int
		)

		err := ctx.Bind(&cu)
		if err != nil {
			ctx.Error(err).SetType(gin.ErrorTypePublic)
			return
		}

		if cu.Name == "" || cu.Password == "" {
			err := errors.New("名前かパスワードがカラの値です")
			ctx.Error(err).SetType(gin.ErrorTypePublic)
			return
		}

		cu.CreatedBy = time.Now()
		chatuser, err := cu.Add(db)

		if err != nil {
			ctx.Error(err).SetType(gin.ErrorTypePublic)
			return
		}

		result = gin.H{
			"sign":   1,
			"result": chatuser,
		}
		status = http.StatusOK
		ctx.JSON(status, result)
	}
}
