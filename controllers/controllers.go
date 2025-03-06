package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HashPassword(password *string) string {
	return ""
}

func Verfy_password(user_password *string, given_password *string) (string, bool) {
	return " ", false
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

        count, err := database.UserCollection.CountDocuments(ctx, bson.M{"email":user.Email})
        if err != nil && count > 0 {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
       
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

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
