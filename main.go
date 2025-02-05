package main

//23;00
import (
	"log"
	"os"

	"github.com/Endale2/Golang_ecommerce_project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.productData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	router.GET("/addtoCart", app.addtoCart())
	router.GET("/removeItem", app.removeItem())
	router.GET("/cartcheckout", app.cartcheckout())
	router.GET("/instantbuy", app.instantbuy())
	log.Fatal(router.Run(":" + port))
}
