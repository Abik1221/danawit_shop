package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/addtocart", app.Addtocart())
	router.GET("/removefromcart", app.Removefromcart())
	router.GET("/cartcheckout", app.Cartcheckout())
	router.GET("/instantbuy", app.Instantbuy())

	log.Fatal(router.Run(":" + port))
}
