package controllers

import (
	"fmt"
	"net/http"
	"time"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string) (string, error) {
	hash, err := bycrypt.GenerateFromPassword([]byte(*password), bycrypt.DefealtCost)
	if err != nil {
		return "", err
	}
	return hash, err
}

func Verfy_password(user_password *string, given_password *string) (string, bool) {
	err := bcrypt.CompareHashAndPassword([]byte(*user_password), []byte(*given_password))
	if err != nil {
		msg := fmt.Sprintf("Error occured while comparing password %s", err.Error())
		return msg, false
	}
	return "", true
}

func SignUP() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Backaground(), 100*time.Second)
		defer cancel()

		var user model.Use
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var validate = valdator.New()

		err = validtor.Struct(user)
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		count, err := database.UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil && count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		count, err = database.UserCollection.CountDocuments(ctx, bson.M{"email": user.Phone})
		if err != nil && count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		password, err := HashPassword(&user.Password)

		user.Created_at, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		user.Updated_at, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		token, refreshToken, err := generate.TokenGenerator(user.Email, user.First_name, user.Last_name, user.User_type)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		user.Token = token
		user.RefreshToken = refreshToken

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var findUser model.User
		var user model.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errror": err.Error()})
			return
		}

		var validationErr = validator.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		err := database.UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&findUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Email"})
			return
		}

		msg, isvalid := Verfy_password(&findUser.Password, &user.Password)
		if msg != "" && isvalid == false {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		findUser.RefreshToken, findUser.Token, err = generate.TokenGenerator(findUser.Email, findUser.First_name, findUser.Last_name, findUser.User_type)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	}
}

func AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ProductView() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Search() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
