package routers

import (
	"log"

	"github.com/irsal-yunus/ms-go-auth-login/Authentication/handlers"
	"github.com/irsal-yunus/ms-go-auth-login/Authentication/models"
	"github.com/irsal-yunus/ms-go-auth-login/Authentication/token"

	"github.com/gin-gonic/gin"
)

type Login struct {
	logger       *log.Logger
	loginHandler *handlers.Login
	flags        *models.Flags
}

func NewRoute(l *log.Logger, f *models.Flags) *Login {
	loginHandler := handlers.NewLogin(l, f)
	token.Init()

	return &Login{
		logger:       l,
		loginHandler: loginHandler,
		flags:        f,
	}
}

func (r *Login) RegisterRoutes() *gin.Engine {
	ginEngine := gin.Default()
	ginEngine.POST("/login", r.loginHandler.Login)
	ginEngine.POST("/verifyToken", r.loginHandler.VerifyToken)
	return ginEngine
}
