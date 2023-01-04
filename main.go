package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mod/controllers"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	//config.AllowOrigins = []string{"http://localhost:3001"}
	router.Use(cors.New(config))

	//dealers
	dealer := router.Group("/dealer")
	dealer.POST("/add", controllers.CreateDealer)
	dealer.GET("/getall", controllers.GetAllDealer)
	dealer.GET("/getallreuqests", controllers.GetAllDealerRequests)
	dealer.POST("/login", controllers.DealerLogin)
	dealer.GET("/getbyid/:id", controllers.GetByID)
	dealer.POST("/form/new", controllers.CreateDealerForm)
	dealer.GET("/form/getall", controllers.GetAllDealerForms)
	dealer.GET("/form/getbyid/:id", controllers.GetFormByID)

	//carts
	cart := router.Group("/cart")
	cart.GET("/getcarts/:id")
	cart.DELETE("/deletecart/:id", controllers.DeleteCartByID)
	cart.POST("/addtocart/:id", controllers.CreateCart)

	//order
	order := router.Group("/order")
	order.POST("/createorder", controllers.CreteOrder)
	order.GET("/getall", controllers.GetAllOrders)
	order.POST("/update/:id", controllers.UpdateOrder)

	//admin
	admin := router.Group("/admin")
	admin.POST("/new", controllers.CreateAdmin)
	admin.POST("/login", controllers.Login)
	admin.DELETE("/delete/:id", controllers.DeleteAdminByID)

	//admin
	contact := router.Group("/contact")
	contact.POST("/new", controllers.CreateContact)
	contact.POST("/getrequests", controllers.GetAllDealerRequests)
	contact.GET("/getamsered", controllers.GetContactForms)
	contact.POST("/update/:id", controllers.UpdateContact)
	contact.GET("/delete/:id", controllers.GetContactFormByID)

	//products
	product := router.Group("/product")
	product.POST("/add", controllers.CreateProduct)
	product.GET("/getbyid/:id", controllers.GetProductByID)
	product.DELETE("/delete/:id", controllers.DeleteProductByID)
	product.GET("/getall", controllers.GetAllProducts)
	product.POST("/uploadimage/:id", controllers.UploadProductImageImage)

	router.Run("localhost:8080")
}
