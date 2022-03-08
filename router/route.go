package router

import (
	"Memo/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/user")
	{
		controller.UserLoginRouteRegister(userRouter)
		controller.UserRegisterRouteRegister(userRouter)
	}

	memoryRouter := router.Group("/memory")
	{
		controller.StoryUploadRouteRegister(memoryRouter)
		controller.CommentUploadRouteRegister(memoryRouter)
	}

	return router
}
