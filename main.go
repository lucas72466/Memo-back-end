package main

import (
	"Memo/router"
)

func main() {
	r := router.InitRouter()

	r.Run()
}
