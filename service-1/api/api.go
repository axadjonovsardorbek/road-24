package api

import (
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
		AllowOrigins:     []string{"http://3.68.216.185:8040", "http://localhost:3600"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/login", h.Login)

	admin := router.Group("/", middleware.JWTMiddleware())
	admin.POST("/admin/user/add", h.Register)
	admin.GET("/admin/user/all", h.GetAllUsers)
	admin.DELETE("/admin/user/delete", h.DeleteProfile)

	protected := router.Group("/", middleware.JWTMiddleware())
	protected.GET("/profile", h.Profile)
	protected.GET("/refresh-token", h.RefreshToken)

	return router
}
