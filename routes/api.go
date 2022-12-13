package routes

import (
	AuthController "go-survia/src/controller/auth"
	"go-survia/src/controller/user"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
	}))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("assets"))))
	route.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"app_name": "survia development",
			"version":  "1.0.0",
		})
	})

	authController := AuthController.AuthAdmin{}
	api := route.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/users", user.Index)

			//authentication
			auth := v1.Group("/auth")
			{
				auth_admin := auth.Group("/admin")
				{
					auth_admin.POST("/sign-in", authController.SignIn)
				}
			}
		}
		
	}
	return route
}
