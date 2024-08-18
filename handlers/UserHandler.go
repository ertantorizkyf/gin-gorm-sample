package handlers

import (
	"fmt"
	"log"

	"github.com/ertantorizkyf/gin-gorm-sample/dtos"
	"github.com/ertantorizkyf/gin-gorm-sample/helpers"
	"github.com/ertantorizkyf/gin-gorm-sample/usecases"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHandler(userUsecase usecases.UserUsecase) UserHandler {
	return UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	body := dtos.RegisterReq{}
	c.Bind(&body)

	// validate mail value
	isEmail := helpers.IsEmail(body.Email)
	if !isEmail {
		message := fmt.Sprintf("%s is not a valid email format", body.Email)
		log.Printf("[ERR] %s", message)
		c.JSON(400, gin.H{
			"message": message,
		})

		return
	}

	statusCode, err := h.userUsecase.RegisterUser(body)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"data": "User registered successfully",
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	body := dtos.LoginReq{}
	c.Bind(&body)

	// validate mail value
	isEmail := helpers.IsEmail(body.Email)
	if !isEmail {
		message := fmt.Sprintf("%s is not a valid email format", body.Email)
		log.Printf("[ERR] %s", message)
		c.JSON(400, gin.H{
			"message": message,
		})

		return
	}

	token, statusCode, err := h.userUsecase.Login(body)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully logged in",
		"token":   token,
	})
}

func (h *UserHandler) SampleMiddlewareImpl(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "HELLO WORLD!",
	})
}
