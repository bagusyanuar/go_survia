package routes

import (
	AdminController "go-survia/src/controller/admin"
	AuthController "go-survia/src/controller/auth"
	MemberController "go-survia/src/controller/member"
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
	adminCategoryController := AdminController.Category{}
	memberCategoryController := MemberController.Category{}
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

			admin := v1.Group("/admin")
			{
				category := admin.Group("/category")
				{
					category.GET("", adminCategoryController.Index)
					category.POST("", adminCategoryController.Index)
					category.GET("/:id", adminCategoryController.FindByID)
					category.POST("/:id", adminCategoryController.FindByID)
					category.DELETE("/:id", adminCategoryController.FindByID)
				}
			}

			member := v1.Group("/member")
			{
				category := member.Group("/category")
				{
					category.GET("", memberCategoryController.Index)
				}
			}
		}

	}
	return route
}
