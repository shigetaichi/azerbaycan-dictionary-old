package main

import (
	"github.com/gin-gonic/gin"
	"go-ddd/infrastructure/email"
	"go-ddd/infrastructure/persistence"
	"go-ddd/interface/handler"
	"go-ddd/usecase"
)

func inject(engine *gin.Engine) {
	// dependencies injection
	// ----- infrastructure -----
	emailDriver := email.New()

	// persistence
	userPersistence := persistence.NewUser()
	wordPersistence := persistence.NewWord()

	// ----- use case -----
	userUseCase := usecase.NewUser(emailDriver, userPersistence)
	wordUseCase := usecase.NewWord(wordPersistence)

	// ----- handler -----
	user := engine.Group("user")
	handler.NewUser(user, userUseCase)
	{
		word := user.Group("word")
		handler.NewWord(word, wordUseCase)
	}
}
