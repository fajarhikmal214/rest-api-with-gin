package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type authController struct {
}

func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Register",
	})
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login",
	})
}
