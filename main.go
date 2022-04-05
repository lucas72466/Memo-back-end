package main

import (
	"Memo/public"
	"Memo/router"
)

func main() {
	r := router.InitRouter()

	public.InitValidator()
	public.InitDefaultLogger()

	if err := r.Run(":80"); err != nil {
		panic(err)
	}
}
