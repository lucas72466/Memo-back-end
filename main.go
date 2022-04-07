package main

import (
	"Memo/conf"
	"Memo/dao"
	"Memo/public"
	"Memo/router"
)

func main() {
	conf.InitConfig()

	dao.InitDAOInst()

	r := router.InitRouter()

	public.InitValidator()
	public.InitDefaultLogger()

	if err := r.Run(":80"); err != nil {
		panic(err)
	}
}
