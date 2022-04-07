package main

import (
	"Memo/conf"
	"Memo/dao"
	"Memo/public"
	"Memo/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	conf.InitConfig()
	dao.InitDAOInst()
	public.InitPublicComponents()

	gin.SetMode(viper.GetString("mode"))
	r := router.InitRouter()

	if err := r.Run(":" + viper.GetString("port")); err != nil {
		panic(err)
	}
}
