package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	goredis "github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/djsilvajr/go-skeleton/internal/config"
	userhandler "github.com/djsilvajr/go-skeleton/internal/domain/user/handler"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/repository"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/service"
	"github.com/djsilvajr/go-skeleton/internal/middleware"
)

func Setup(cfg *config.Config, db *gorm.DB, rdb *goredis.Client) *gin.Engine {
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())
	r.Use(gin.Recovery())

	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	userHandler := userhandler.NewUserHandler(userSvc)
	authHandler := userhandler.NewAuthHandler(userSvc, cfg)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// r.GET("/api/documentation/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	api.Use(middleware.RateLimit(60, time.Minute))
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
		}

		protected := api.Group("/")
		protected.Use(middleware.Auth(cfg))
		{
			users := protected.Group("/users")
			{
				users.GET("", userHandler.List)
				users.GET("/:id", userHandler.Show)
				users.POST("", userHandler.Store)
				users.PUT("/:id", userHandler.Update)
				users.DELETE("/:id", middleware.AdminOnly(), userHandler.Destroy)
			}
		}
	}

	return r
}
