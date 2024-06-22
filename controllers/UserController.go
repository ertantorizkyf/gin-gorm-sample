package controllers

import (
	"fmt"
	"log"

	"github.com/ertantorizkyf/gin-gorm-sample/helpers"
	"github.com/ertantorizkyf/gin-gorm-sample/initializers"
	"github.com/ertantorizkyf/gin-gorm-sample/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var body struct {
		FullName string
		Email    string
		Password string
	}

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

	// check if email already exist in db
	var existingUser models.User
	initializers.DB.Where("email = ?", body.Email).First(&existingUser)
	if existingUser.Email != "" {
		message := fmt.Sprintf("User with %s email already exists", body.Email)
		log.Printf("[ERR] %s", message)
		c.JSON(400, gin.H{
			"message": message,
		})

		return
	}

	// store to db
	body.Password, _ = helpers.HashPassword(body.Password)
	user := models.User{FullName: body.FullName, Email: body.Email, Password: body.Password}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		message := "Error while registering user"
		log.Printf("[ERR] %s", message)
		c.JSON(500, gin.H{
			"message": message,
		})

		return
	}

	c.JSON(200, gin.H{
		"data": "Successfully register user",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

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

	// check if email already exist in db
	var user models.User
	initializers.DB.Where("email = ?", body.Email).First(&user)
	if user.Email == "" {
		message := "Incorrect email or password"
		log.Printf("[ERR] %s", message)
		c.JSON(400, gin.H{
			"message": message,
		})

		return
	}

	// check password
	isMatch := helpers.CheckPasswordHash(body.Password, user.Password)

	if !isMatch {
		message := "Incorrect email or password"
		log.Printf("[ERR] %s", message)
		c.JSON(400, gin.H{
			"message": message,
		})

		return
	}

	// generate jwt
	token, err := helpers.GenerateToken(user)
	if err != nil {
		message := "Something went wrong while generating auth token"
		log.Printf("[ERR] %s", message)
		c.JSON(400, gin.H{
			"message": message,
		})

		return
	}

	c.JSON(200, gin.H{
		"data":  "Successfully logged in",
		"token": token,
	})
}
