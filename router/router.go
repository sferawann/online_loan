package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/pinjol/controller"
	"github.com/sferawann/pinjol/repository"
)

func NewRouter(userRepository repository.UserRepo, userCon *controller.UserCon, roleCon *controller.RoleCon, proCon *controller.ProductCon, paymedCon *controller.PaymentMethodCon, traCon *controller.TraCon, accstatCon *controller.AcceptStatusCon, PayCon *controller.PaymentCon) *gin.Engine {
	r := gin.Default()

	r.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := r.Group("/api")

	// authenticationRouter := router.Group("/auth")
	// {
	// 	authenticationRouter.POST("/register", authController.Register)
	// 	authenticationRouter.POST("/login", authController.Login)
	// }

	// usersRouter := router.Group("/user")
	// {
	// 	usersRouter.POST("/", userCon.Insert)
	// 	usersRouter.GET("/", userCon.FindAll)
	// 	usersRouter.GET("/:id", userCon.FindById)
	// 	usersRouter.GET("/username/:username", userCon.FindByUsername)
	// 	usersRouter.PUT("/:id", userCon.Update)
	// 	usersRouter.DELETE("/:id", userCon.Delete)
	// }

	// roleRouter := router.Group("/role")
	// {
	// 	roleRouter.POST("/", roleCon.Insert)
	// 	roleRouter.GET("/", roleCon.FindAll)
	// 	roleRouter.GET("/:id", roleCon.FindById)
	// 	roleRouter.GET("/name/:name", roleCon.FindByName)
	// 	roleRouter.PUT("/:id", roleCon.Update)
	// 	roleRouter.DELETE("/:id", roleCon.Delete)
	// }

	// proRouter := router.Group("/product")
	// {
	// 	proRouter.POST("/", proCon.Insert)
	// 	proRouter.GET("/", proCon.FindAll)
	// 	proRouter.GET("/:id", proCon.FindById)
	// 	proRouter.GET("/name/:name", proCon.FindByName)
	// 	proRouter.PUT("/:id", proCon.Update)
	// 	proRouter.DELETE("/:id", proCon.Delete)
	// }

	// paymedRouter := router.Group("/payment_method")
	// {
	// 	paymedRouter.POST("/", paymedCon.Insert)
	// 	paymedRouter.GET("/", paymedCon.FindAll)
	// 	paymedRouter.GET("/:id", paymedCon.FindById)
	// 	paymedRouter.GET("/name/:name", paymedCon.FindByName)
	// 	paymedRouter.PUT("/:id", paymedCon.Update)
	// 	paymedRouter.DELETE("/:id", paymedCon.Delete)
	// }

	// traRouter := router.Group("/transaction")
	// {
	// 	traRouter.POST("/", traCon.Insert)
	// 	traRouter.GET("/", traCon.FindAll)
	// 	traRouter.GET("/:id", traCon.FindById)
	// traRouter.PUT("/:id", traCon.Update)
	// 	traRouter.DELETE("/:id", traCon.Delete)
	// }

	// accstatRouter := router.Group("/accept_status")
	// {
	// 	accstatRouter.POST("/", accstatCon.Insert)
	// 	accstatRouter.GET("/", accstatCon.FindAll)
	// 	accstatRouter.GET("/:id", accstatCon.FindById)
	// 	accstatRouter.PUT("/:id", accstatCon.Update)
	// 	accstatRouter.DELETE("/:id", accstatCon.Delete)
	// }

	payRouter := router.Group("/payment")
	{
		payRouter.POST("/", PayCon.Insert)
		payRouter.GET("/", PayCon.FindAll)
		payRouter.GET("/:id", PayCon.FindById)
		payRouter.PUT("/:id", PayCon.Update)
		payRouter.DELETE("/:id", PayCon.Delete)
	}

	return r
}
