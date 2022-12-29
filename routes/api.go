package routes

import (
	"go-survia/src/controller"
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
	adminCampaignController := AdminController.Campaign{}
	adminCityController := AdminController.City{}
	adminJobController := AdminController.Job{}
	adminSecController := AdminController.Sec{}
	adminSecQuestionController := AdminController.SecQuestion{}

	memberCategoryController := MemberController.Category{}

	categoryController := controller.Category{}
	bankController := controller.Bank{}
	provinceController := controller.Province{}
	cityController := controller.City{}

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
					category.GET("", categoryController.Index)
					category.POST("", categoryController.Store)
					category.GET("/:id", categoryController.Show)
					category.PATCH("/:id", categoryController.Update)
					category.DELETE("/:id", categoryController.Destroy)
				}

				bank := admin.Group("/bank")
				{
					bank.GET("", bankController.Index)
					bank.POST("", bankController.Store)
					bank.GET("/:id", bankController.Show)
					bank.PATCH("/:id", bankController.Update)
					bank.DELETE("/:id", bankController.Destroy)
				}

				province := admin.Group("/province")
				{
					province.GET("", provinceController.Index)
					province.POST("", provinceController.Store)
					province.GET("/:id", provinceController.Show)
					province.PATCH("/:id", provinceController.Update)
					province.DELETE("/:id", provinceController.Destroy)
				}

				city := admin.Group("/city")
				{
					city.GET("", cityController.Index)
					city.POST("", cityController.Store)
					city.GET("/:id", cityController.Show)
					city.PATCH("/:id", adminCityController.FindByID)
					city.DELETE("/:id", adminCityController.FindByID)
				}

				campaign := admin.Group("/campaign")
				{
					campaign.GET("", adminCampaignController.Index)
					campaign.POST("", adminCampaignController.Index)
					campaign.GET("/:id", adminCampaignController.FindByID)
					campaign.PATCH("/:id", adminCampaignController.FindByID)
					campaign.DELETE("/:id", adminCampaignController.FindByID)
				}

				job := admin.Group("/job")
				{
					job.GET("", adminJobController.Index)
					job.POST("", adminJobController.Index)
					job.GET("/:id", adminJobController.FindByID)
					job.PATCH("/:id", adminJobController.FindByID)
					job.DELETE("/:id", adminJobController.FindByID)
				}
				sec := admin.Group("/sec")
				{
					sec.GET("", adminSecController.Index)
					sec.POST("", adminSecController.Index)
					sec.GET("/:id", adminSecController.FindByID)
					sec.PATCH("/:id", adminSecController.FindByID)
					sec.DELETE("/:id", adminSecController.FindByID)
				}
				secQuestion := admin.Group("/sec-question")
				{
					secQuestion.GET("", adminSecQuestionController.Index)
					secQuestion.POST("", adminSecQuestionController.Index)
					secQuestion.GET("/:id", adminSecQuestionController.FindByID)
					secQuestion.PATCH("/:id", adminSecController.FindByID)
					secQuestion.DELETE("/:id", adminSecQuestionController.FindByID)
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
