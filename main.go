package main

import "Memo/router"

func main() {
	r := router.InitRouter()

	r.Run()
	//fmt.Println(public.GenerateUserToken("shifeize"))
	//
	//r := gin.Default()
	//r.Use(middleware.JWTAuth())
	//r.POST("/test", func(context *gin.Context) {
	//	info, err := public.GeCtUserTokenInfoFromContext(context)
	//	if err != nil {
	//		panic(err)
	//	}
	//	log.Println(public.JsonString(info))
	//})
	//r.Run()

	//tokenString, err := public.GenerateUserToken("shifeize")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(tokenString)
	//_, claim, err := public.ParseUserToken(tokenString)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(public.JsonString(claim))
}
