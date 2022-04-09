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

	memoryCreateRouter := router.Group("/memory/crate")
	memoryCreateRouter.Use(middleware.JWTAuth())
	{
		controller.CreateStoryRouteRegister(memoryCreateRouter)
		controller.CreateCommentRouteRegister(memoryCreateRouter)
	}

	memorySearchRouter := router.Group("/memory/search")
	{
		controller.SearchCommentRouteRegister(memorySearchRouter)
	}

	fileRouter := router.Group("/file")
	{
		controller.PictureUploadRouteRegister(fileRouter)
	}

	return router
}
