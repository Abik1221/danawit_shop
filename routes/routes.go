package routes

import "github.com/gin-gonic/gin"

func User_Routes(incoming_routes gin.Engine){
	incoming_routes.POST("users/signup", controllers.SignUP())
	incoming_routes.POST("users/login", controllers.Login())
    incoming_routes.POST("admins/addproduct", controllers.AddProduct())
    incoming_routes.GET("users/product_view", controllers.ProductView())
    incoming_routes.GET("users/search", controllers.search())
}