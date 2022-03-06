package main

import (
	"Memo/router"
	"log"
)

func main() {
	r := router.InitRouter()
	log.Println("start running...")
	r.Run()
}
