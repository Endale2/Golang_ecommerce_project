package routes

import (
	"github.com/Endale2/Golang_ecommerce_project/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.SignUp())
	incomingRoutes.POST("users/login", controllers.Login())
	incomingRoutes.POST("admin/addProduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("users/product-view", controllers.SearchProduct())
	incomingRoutes.GET("users/search", controllers.SearchProductByQuery())
}
