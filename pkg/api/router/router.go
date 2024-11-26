package router

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.Engine {
	// Set CORS config
	// r.Use(cors.New(cors.Config{
	// 	AllowCredentials: false,
	// 	AllowOrigins: []string{"http://localhost:3333"},
	// 	AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTION", "HEAD", "PATCH", "COMMON"},
	// 	AllowHeaders: []string{"Content-Type", "Content-Length", "Authorization", "accept", "origin", "Referer", "User-Agent"},
	// }))

	// r.Use(middleware.CORSMiddleware())

	// api := handler.NewUserHandler()
	return nil
}