package routes

import (
	"net/http"

	"github.com/TakasBU/TakasBU/handlers"
	"github.com/labstack/echo/v4"
)

// TODO GRUP SİSTEMİ GETİRELECEK CHAT CHANNEL YAPILACAK ÜLKEDE KİMSE AÇ KALMAYACAK FOLLOWER SİSTEMİ
func Route(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	User(e)
	Product(e)
	AuthRoute(e)
	//FIXME NOT FOUND DİYOR ABİ KAFAYI YİCEM
	e.Static("/static", "static")
}

func User(e *echo.Echo) {
	userRepo := handlers.NewUser()

	e.GET("/api/users/:id", userRepo.GetUserById)
	e.GET("/api/users", userRepo.GetUsers)
	e.PUT("/api/users/:id", userRepo.UpdateUser)
	e.POST("/api/users", userRepo.CreateUser)
	e.DELETE("/api/users/:id", userRepo.DeleteUser)

}

func Product(e *echo.Echo) {
	ProductRepo := handlers.NewProduct()

	e.GET("/api/products/:id", ProductRepo.GetProductById)
	e.GET("/api/products", ProductRepo.GetProducts)
	e.PUT("/api/products/:id", ProductRepo.UpdateProduct)
	e.POST("/api/products", ProductRepo.CreateProduct)
	e.DELETE("/api/products/:id", ProductRepo.DeleteProduct)

}

func AuthRoute(e *echo.Echo) {
	rc := handlers.NewAuthController()
	router := e.Group("/api/auth")

	router.POST("/register", rc.SignUpUser)
	router.POST("/login", rc.SignInUser)
	/*	router.GET("/logout", middleware.DeserializeUser(), rc.authController.LogoutUser)
		router.GET("/verifyemail/:verificationCode", rc.authController.VerifyEmail) */
}
