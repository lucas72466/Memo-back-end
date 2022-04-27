package router

import (
	"Memo/controller"
	"Memo/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.InjectLogID())
	router.Use(middleware.Cors())

	userRouter := router.Group("/user")
	{
		controller.UserLoginRouteRegister(userRouter)
		controller.UserRegisterRouteRegister(userRouter)
	}

	memoryRouter := router.Group("/memory")

	memoryCreateRouter := memoryRouter.Group("/create")
	memoryCreateRouter.Use(middleware.JWTAuth())
	{
		controller.CreateStoryRouteRegister(memoryCreateRouter)
		controller.CreateCommentRouteRegister(memoryCreateRouter)
	}

	memorySearchRouter := memoryRouter.Group("/search")
	{
		controller.SearchCommentRouteRegister(memorySearchRouter)
		controller.SearchStoryRouteRegister(memorySearchRouter)
	}

	memoryDeleteRouter := memoryRouter.Group("/delete")
	memoryDeleteRouter.Use(middleware.JWTAuth())
	{
		controller.DeleteMemoryRouteRegister(memoryDeleteRouter)
	}

	memoryHugRouter := memoryRouter.Group("/hug")
	{
		controller.HugAddRouteRegister(memoryHugRouter)
	}

	fileRouter := router.Group("/file")
	{
		controller.PictureUploadRouteRegister(fileRouter)
	}

	return router
}
