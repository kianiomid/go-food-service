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
		user.POST("/register", middleware.CORSMiddleware(), userHandler.RegisterUser)
		user.GET("/:id" /*middleware.AuthMiddleware(),*/, middleware.CORSMiddleware(), userHandler.FindUserById)
		user.PUT("/:id", middleware.CORSMiddleware(), userHandler.UpdateUser)
		user.DELETE("/:id", middleware.CORSMiddleware(), userHandler.DeleteUserById)
		user.POST("/login", middleware.CORSMiddleware(), userHandler.Login)
	}

	//FoodHandler
	foodRepository := repository.NewFoodRepository(db)
	foodService := service.NewFoodService(foodRepository, userService)
	foodHandler := handler.NewFoodHandler(foodService)

	food := router.Group("/v1/food")
	{
		food.GET("/", middleware.CORSMiddleware() /*middleware.AuthMiddleware(),*/, foodHandler.GetAllFood)
		food.POST("/", middleware.CORSMiddleware() /*middleware.AuthMiddleware(),*/, foodHandler.SaveFood)
		food.GET("/:id", middleware.CORSMiddleware() /*middleware.AuthMiddleware(),*/, foodHandler.FindFoodById)
		food.PUT("/:id", middleware.CORSMiddleware() /*middleware.AuthMiddleware(),*/, foodHandler.UpdateFood)
		food.DELETE("/:id", middleware.CORSMiddleware() /*middleware.AuthMiddleware(),*/, foodHandler.DeleteById)
	}

	return router
}
