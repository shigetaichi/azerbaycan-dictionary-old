package main

import (
	"context"
	"go-ddd/infrastructure/email"
	"go-ddd/infrastructure/persistence"
	"go-ddd/interface/handler"
	"go-ddd/usecase"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ken109/gin-jwt"
	"go-ddd/constant"
	"go-ddd/infrastructure/log"
	"go-ddd/interface/middleware"
)

func main() {
	logger := log.Logger()

	err := jwt.SetUp(
		jwt.Option{
			Realm:            constant.DefaultRealm,
			SigningAlgorithm: jwt.HS256,
			SecretKey:        []byte(os.Getenv("GIN_JWT_SECRET")),
			//SecretKey: []byte(config.Env.App.Secret),
		},
	)
	if err != nil {
		panic(err)
	}
	logger.Info("Succeeded in setting up JWT.")

	engine := gin.New()

	engine.Use(middleware.Log(logger, time.RFC3339, false))
	engine.Use(middleware.RecoveryWithLog(logger, true))

	engine.GET("health", func(c *gin.Context) { c.Status(http.StatusOK) })

	// cors
	engine.Use(
		cors.New(
			cors.Config{
				AllowOriginFunc: func(origin string) bool {
					return true
				},
				AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
				AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
				AllowCredentials: true,
				MaxAge:           12 * time.Hour,
			},
		),
	)

	inject(engine)

	logger.Info("Succeeded in setting up routes.")

	// serve
	var port = ":8080"
	if portEnv := os.Getenv("PORT"); portEnv != "" {
		port = ":" + portEnv
	}

	srv := &http.Server{
		Addr:    port,
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	logger.Info("Succeeded in listen and serve.")

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %+v", err)
	}

	logger.Info("Server exiting")
}

func inject(engine *gin.Engine) {
	// dependencies injection
	// ----- infrastructure -----
	emailDriver := email.New()

	// persistence
	userPersistence := persistence.NewUser()
	wordPersistence := persistence.NewWord()
	draftPersistence := persistence.NewDraft()

	// ----- use case -----
	userUseCase := usecase.NewUser(emailDriver, userPersistence)
	wordUseCase := usecase.NewWord(wordPersistence)
	draftUseCase := usecase.NewDraft(draftPersistence, wordPersistence)

	// ----- handler -----
	user := engine.Group("user")
	handler.NewUser(user, userUseCase)
	{
		word := user.Group("word")
		handler.NewWord(word, wordUseCase)
		draft := user.Group("draft")
		handler.NewDraft(draft, draftUseCase)
	}
}
