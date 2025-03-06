package controllers

import "github.com/gin-gonic/gin"

func HashPassword(password *string) string {
     return ""
}

func Verfy_password(user_password *string, given_password *string) (string, bool) {
   return " ", false
}

func SignUP() gin.HandlerFunc {
    return func(c *gin.Context){
        
    }
}

func Login() gin.HandlerFunc {
   return func(c *gin.Context){

   }
}

func AddProduct() gin.HandlerFunc {
   return func(c *gin.Context){

   }
}

func ProductView() gin.HandlerFunc {
    return func(c *gin.Context){

    }
}

func Search() gin.HandlerFunc {
    return func(c *gin.Context){

    }
}
