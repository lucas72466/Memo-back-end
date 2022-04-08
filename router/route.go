package router

import (
	"Memo/controller"
	"Memo/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.InjectLogID())

	userRouter := router.Group("/user")
	{
		controller.UserLoginRouteRegister(userRouter)
		controller.UserRegisterRouteRegister(userRouter)
	}

	memoryRouter := router.Group("/memory")
	memoryRouter.Use(middleware.JWTAuth())
	{
		controller.CreateStoryRouteRegister(memoryRouter)
		controller.CreateCommentRouteRegister(memoryRouter)
	}

	fileRouter := router.Group("/file")
	{
		controller.PictureUploadRouteRegister(fileRouter)
	}

	return router
}
