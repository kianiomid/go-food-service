package route

import (
	"food-service/internal/domain/handler"
	"food-service/internal/domain/repository"
	"food-service/internal/domain/service"
	"food-service/internal/presentation/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	//UserHandler
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	user := router.Group("/v1/user")
	{
		user.GET("/", middleware.CORSMiddleware(), userHandler.GetAllUser)
		user.GET("/:id", middleware.CORSMiddleware(), middleware.AuthMiddleware(), userHandler.FindUserById)
		user.POST("", middleware.CORSMiddleware(), middleware.AuthMiddleware(), userHandler.RegisterUser)
		user.PUT("/:id", middleware.CORSMiddleware(), middleware.AuthMiddleware(), userHandler.UpdateUser)
		user.DELETE("/:id", middleware.CORSMiddleware(), middleware.AuthMiddleware(), userHandler.DeleteUserById)
		user.POST("/login", middleware.CORSMiddleware(), middleware.AuthMiddleware(), userHandler.Login)
	}

	//FoodHandler
	foodRepository := repository.NewFoodRepository(db)
	foodService := service.NewFoodService(foodRepository, userRepository)
	foodHandler := handler.NewFoodHandler(foodService)

	food := router.Group("/v1/food")
	{
		food.GET("/", middleware.CORSMiddleware(), middleware.AuthMiddleware(), foodHandler.GetAllFood)
		food.GET("/:id", middleware.CORSMiddleware(), middleware.AuthMiddleware(), foodHandler.FindFoodById)
		food.POST("", middleware.CORSMiddleware(), middleware.AuthMiddleware(), foodHandler.SaveFood)
		food.PUT("/:id", middleware.CORSMiddleware(), middleware.AuthMiddleware(), foodHandler.UpdateFood)
		food.DELETE("/:id", middleware.CORSMiddleware(), middleware.AuthMiddleware(), foodHandler.DeleteById)
	}

	return router
}
