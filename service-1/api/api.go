package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "auth/api/docs"
	"auth/api/handler"
	"auth/api/middleware"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewApi(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/login", h.Login) //for admins -> front and mobile(for gai login)
	router.POST("/login/driver", h.LoginDriver) //for driver -> mobile

	admin := router.Group("/", middleware.JWTMiddleware())
	admin.POST("/admin/user/add", h.Register) //for admins -> front
	admin.GET("/admin/user/all", h.GetAllUsers) //for admins -> front
	admin.DELETE("/admin/user/delete", h.DeleteProfile) //for admins -> front


	admin.POST("/admin/car/add", h.CarCreate) //for admins -> front
	admin.GET("/admin/car/{id}", h.CarGetById) //for admins and driver -> front and mobile
	admin.GET("/admin/car/all", h.CarGetAll) //for admins -> front (maybe mobile too for gai)
	admin.PUT("/car/upload/image", h.CarUpdate) //for driver -> mobile
	admin.DELETE("/admin/car/{id}", h.CarDelete) //for admin -> front

	admin.POST("/admin/car/type/add", h.CarTypeCreate) //for admins -> front
	admin.GET("/admin/car/type/{id}", h.CarTypeGetById) //for admins and driver -> front and mobile
	admin.GET("/admin/car/type/all", h.CarTypeGetAll) //for admins -> front and mobile
	admin.DELETE("/admin/car/type/{id}", h.CarTypeDelete) //for admin -> front

	admin.POST("/admin/car/service/add", h.CarServiceCreate) //for admins -> front
	admin.GET("/admin/car/service/{id}", h.CarServiceGetById) //for admins and driver -> front and mobile
	admin.GET("/admin/car/service/all", h.CarServiceGetAll) //for admins and drivers with own car_number -> front and mobile
	admin.DELETE("/admin/car/service/{id}", h.CarServiceDelete) //for admin -> front

	admin.POST("/admin/driver/licence/add", h.LicenceCreate) //for admins -> front
	admin.GET("/admin/driver/licence/{id}", h.LicenceGetById) //for admins -> front
	admin.GET("/admin/driver/licence/all", h.LicenceGetAll) //for admins -> front
	admin.DELETE("/admin/driver/licence/{id}", h.LicenceDelete) //for admin -> front

	admin.POST("/admin/fine/add", h.FineCreate) //for gai -> mobile
	admin.GET("/admin/fine/{id}", h.FineGetById) //for admins and driver -> front and mobile
	admin.GET("/admin/fine/all", h.FineGetAll) //for admins and driver -> front and mobile
	admin.PUT("/admin/fine/{id}", h.FineUpdate) //for driver-> mobile for change image_url

	admin.POST("/admin/service/add", h.ServiceCreate) //for admins -> front
	admin.GET("/admin/service/{id}", h.ServiceGetById) //for admins -> front
	admin.GET("/admin/service/all", h.ServiceGetAll) //for admins -> front
	admin.DELETE("/admin/service/{id}", h.ServiceDelete) //for admin -> front

	admin.POST("/admin/service/type/add", h.ServiceTypeCreate) //for admins -> front
	admin.GET("/admin/service/type/{id}", h.ServiceTypeGetById) //for admins -> front
	admin.GET("/admin/service/type/all", h.ServiceTypeGetAll) //for admins -> front
	admin.DELETE("/admin/service/type/{id}", h.ServiceTypeDelete) //for admin -> front

	protected := router.Group("/", middleware.JWTMiddleware())
	protected.GET("/profile", h.Profile) //for admins -> front (maybe mobile too)
	protected.GET("/refresh-token", h.RefreshToken) //front and mobile
	protected.GET("/refresh-token/driver", h.RefreshTokenForDriver) //for mobile

	return router
}
